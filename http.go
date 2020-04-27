package godingtalk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	ErrEmptyParams = errors.New("params is require")
)

type Unmarshalable interface {
	checkErr() error
}

func (rd RequestData) Set(key string, value interface{}) {
	switch value := value.(type) {
	case []interface{}:
		rd[key] = value
		return
	}
	rd[key] = value
}

func (rd RequestData) ToBytes() ([]byte, error) {
	byte, err := json.Marshal(rd)
	if err != nil {
		return nil, err
	}
	return byte, nil
}

// httpRequestWithStd 使用原生http实现
// 以reqData为判断, nil == Get, !nil == Post
func (d *DingtalkClient) httpRequestWithStd(path string, params url.Values, reqData interface{}, respData Unmarshalable) error {
	if params == nil {
		return ErrEmptyParams
	}
	client := d.Client
	if params.Get("appkey") == "" {
		params.Set("appkey", d.AppKey)
		params.Set("appsecret", d.AppSecret)
	}
	uri := &url.URL{
		Scheme:   "https",
		Host:     d.BaseURL,
		Path:     path,
		RawQuery: params.Encode(),
	}
	data, err := request(client, uri.String(), reqData)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, respData); err != nil {
		return err
	}
	return nil
}

// httpRequestWithFastHttp 使用fasthttp实现
// func httpRequestWithFastHttp() {}

func request(client *http.Client, uri string, reqData interface{}) ([]byte, error) {
	var request *http.Request
	var err error
	if reqData != nil {
		// POST
		rd, err := reqData.(RequestData).ToBytes()
		if err != nil {
			return nil, err
		}
		request, err = http.NewRequest(http.MethodPost, uri, bytes.NewReader(rd))
		if err != nil {
			return nil, err
		}
		request.Header.Set("Content-Type", "application/json")
	} else {
		// GET
		request, err = http.NewRequest(http.MethodGet, uri, nil)
		if err != nil {
			return nil, err
		}
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}