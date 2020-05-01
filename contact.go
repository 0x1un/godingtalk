package godingtalk

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type AuthScopesResp struct {
	Base
	AuthUserField []string `json:"auth_user_field"`
	AuthOrgScopes struct {
		AuthedDept []int    `json:"authed_dept"`
		AuthedUser []string `json:"authed_user"`
	} `json:"auth_org_scopes"`
}

type UserCreateResp struct {
	Base
	Userid string `json:"userid"`
}

type UserUpdateResp struct {
	Base
}

type UserDeleteResp struct {
	Base
}

type UserCreateReq struct {
	Mobile string `json:"mobile"` // 手机号码，企业内必须唯一，不可重复。如果是国际号码，请使用+xx-xxxxxx的格式
	UserUpdateReq
}

type UserUpdateReq struct {
	Userid       string `json:"userid"`       // 员工在当前企业内的唯一标识，也称staffId。可由企业在创建时指定，并代表一定含义比如工号，创建后不可修改，企业内必须唯一。长度为1~64个字符，如果不传，服务器将自动生成一个userid
	Name         string `json:"name"`         // 必要, 用户名称
	OrderInDepts string `json:"orderInDepts"` // 在对应的部门中的排序，Map结构的json字符串，key是部门的Id, value是人员在这个部门的排序值
	Department   []int  `json:"department"`   // 数组类型，数组里面值为整型，成员所属部门id列表
	Position     string `json:"position"`     // 职位信息。长度为0~64个字符
	Tel          string `json:"tel"`          // 必要, 分机号，长度为0~50个字符，企业内必须唯一，不可重复
	WorkPlace    string `json:"workPlace"`    // 办公地点，长度为0~50个字符
	Remark       string `json:"remark"`       // 备注，长度为0~1000个字符
	Email        string `json:"email"`        // 邮箱。长度为0~64个字符。企业内必须唯一，不可重复
	OrgEmail     string `json:"orgEmail"`     // 员工的企业邮箱，员工的企业邮箱已开通，才能增加此字段， 否则会报错
	Jobnumber    string `json:"jobnumber"`    // 员工工号。对应显示到OA后台和客户端个人资料的工号栏目。长度为0~64个字符
	IsHide       bool   `json:"isHide"`       // 是否号码隐藏，true表示隐藏，false表示不隐藏。隐藏手机号后，手机号在个人资料页隐藏，但仍可对其发DING、发起钉钉免费商务电话。
	IsSenior     bool   `json:"isSenior"`     // 是否高管模式，true表示是，false表示不是
	Extattr      struct {
		Interest string `json:"爱好"`
		Age      string `json:"年龄"`
	} `json:"extattr"` // 扩展属性，可以设置多种属性（手机上最多显示10个扩展属性.具体显示哪些属性，请到OA管理后台->设置->通讯录信息设置和OA管理后台->设置->手机端显示信息设置）。该字段的值支持链接类型填写，同时链接支持变量通配符自动替换，目前支持通配符有：userid，corpid。示例： [工位地址](http://www.dingtalk.com?userid=#userid#&corpid=#corpid#)
}
type UserGetResp struct {
	Base
	Unionid         string `json:"unionid"`
	Remark          string `json:"remark"`
	Userid          string `json:"userid"`
	IsLeaderInDepts string `json:"isLeaderInDepts"`
	IsBoss          bool   `json:"isBoss"`
	HiredDate       int64  `json:"hiredDate"`
	IsSenior        bool   `json:"isSenior"`
	Tel             string `json:"tel"`
	Department      []int  `json:"department"`
	WorkPlace       string `json:"workPlace"`
	Email           string `json:"email"`
	OrderInDepts    string `json:"orderInDepts"`
	Mobile          string `json:"mobile"`
	Active          bool   `json:"active"`
	Avatar          string `json:"avatar"`
	IsAdmin         bool   `json:"isAdmin"`
	IsHide          bool   `json:"isHide"`
	Jobnumber       string `json:"jobnumber"`
	Name            string `json:"name"`
	Extattr         struct {
	} `json:"extattr"`
	StateCode string `json:"stateCode"`
	Position  string `json:"position"`
	Roles     []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		GroupName string `json:"groupName"`
	} `json:"roles"`
}

func (u UserUpdateReq) ToBytes() ([]byte, error) {
	return json.Marshal(u)

}

// OapiAuthScopesRequest 获取通讯录权限范围
func (d *DingtalkClient) OapiAuthScopesRequest() (AuthScopesResp, error) {
	var asr AuthScopesResp
	err := d.httpRequestWithStd("auth/scopes", d.params, nil, &asr)
	if err != nil {
		return asr, err
	}
	return asr, nil
}

// OapiUserCreateRequest 创建一个用户 Method: POST
func (d *DingtalkClient) OapiUserCreateRequest(userInfo UserCreateReq) (UserCreateResp, error) {
	var respData UserCreateResp
	err := d.httpRequestWithStd("user/create", d.params, userInfo, respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}

// OapiUserUpdateRequest 更新用户的信息 Method: POST
func (d *DingtalkClient) OapiUserUpdateRequest(userInfo UserUpdateReq) (UserUpdateResp, error) {
	var respData UserUpdateResp
	err := d.httpRequestWithStd("user/create", d.params, userInfo, respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}

// OapiUserDeleteRequest 删除一个用户 Method: GET
func (d *DingtalkClient) OapiUserDeleteRequest(userid string) (UserDeleteResp, error) {
	params := d.params
	params.Set("userid", userid)
	var respData UserDeleteResp
	err := d.httpRequestWithStd("user/delete", params, nil, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}

// OapiUserGetRequest 获取一个用户的详情 Method: GET
func (d *DingtalkClient) OapiUserGetRequest(userid string) (UserGetResp, error) {
	params := d.params
	params.Set("userid", userid)
	var respData UserGetResp
	err := d.httpRequestWithStd("user/get", params, nil, &respData)
	if err != nil {
		fmt.Println("xxx")
		return respData, err
	}
	return respData, nil
}

func getTagName() {
	s := &UserCreateReq{}
	box := make(RequestData)
	t := reflect.TypeOf(s).Elem()
	parseStructAssig2Map(t, box)
	fmt.Println(box)

}
func parseStructAssig2Map(t reflect.Type, box RequestData) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Type.Kind() == reflect.Struct {
			parseStructAssig2Map(field.Type, box)
			continue
		}
		key := field.Tag.Get("json")
		fmt.Println(key)
		box.Set(key, field.Name)
	}
}
