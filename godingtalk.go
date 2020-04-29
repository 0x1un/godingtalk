package godingtalk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// 基础设置
const (
	ExpiresIn = 7200
	OapiURL   = "oapi.dingtalk.com"
	TokenFile = ".token"
)

// 错误集合
var (
	ErrEmptyToken     = errors.New("can't refresh a empty token")
	ErrTokenAvaliable = errors.New("token is avaliable")
)

type Base struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type DingtalkClient struct {
	Client      *http.Client
	BaseURL     string
	AppKey      string
	AppSecret   string
	AccessToken AccessTokenResp
}

type AccessTokenContent struct {
	Token       string `json:"access_token"` // 这里用来存储之后要用的凭证码
	ExpiresTime int64  `json:"expires_in"`   // 凭证码过期的时间(= 凭证创建时间 + 7200s)
}

type AccessTokenResp struct {
	Base
	AccessTokenContent
}

type RequestData map[string]interface{}

// 为Base实现Unmarshallable接口，ErrorCode不为空，肯定为可Unmarshal的
func (b Base) checkErr() error {
	if b.ErrCode != 0 {
		return fmt.Errorf("%d:%s", b.ErrCode, b.ErrMsg)
	}
	return nil
}

func NewDingtalkClient(appkey, appsecret string) *DingtalkClient {
	return &DingtalkClient{
		Client: &http.Client{
			Timeout: time.Duration(time.Second * 60),
		},
		BaseURL:   OapiURL,
		AppKey:    appkey,
		AppSecret: appsecret,
	}
}

// setToken 获取到token后存入Client中并持久化存储到文本
func (d *DingtalkClient) setToken() error {
	// tr => tokenResponse
	params := url.Values{}
	access := AccessTokenResp{}
	err := d.httpRequestWithStd("gettoken", params, nil, &access)
	if err != nil {
		return err
	}
	if access.ErrCode != 0 {
		return errors.New(access.ErrMsg)
	}
	d.AccessToken.Token = access.Token
	d.AccessToken.ExpiresTime = time.Now().Unix() + ExpiresIn

	err = storeToken(d.AccessToken)
	if err != nil {
		return err
	}
	return nil
}

// RefreshToken 刷新token,根据现有Client中的token判断
// 暴露为外部可调用函数，方便手动刷新现有token的需求
func (d *DingtalkClient) RefreshToken() error {
	if d.AccessToken.Token == "" {
		return ErrEmptyToken
	}
	// token过期则重新获取token
	if checkExpireTime(d.AccessToken.ExpiresTime) {
		return d.setToken()
	}
	return ErrTokenAvaliable
}

// checkExpireTime 判断token是否过期
func checkExpireTime(etime int64) bool {
	if time.Now().Unix() < etime {
		return false
	}
	return true
}

func storeToken(tokenResp Unmarshalable) error {
	if err := tokenResp.checkErr(); err != nil {
		return err
	}
	data, err := json.Marshal(tokenResp)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(TokenFile, data, 0666)
}

func readToken() (Unmarshalable, error) {
	bytes, err := ioutil.ReadFile(TokenFile)
	if err != nil {
		return nil, err
	}
	tokenResp := &AccessTokenResp{}
	err = json.Unmarshal(bytes, tokenResp)
	return tokenResp, err
}
