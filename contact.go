package godingtalk

import "net/url"

type AuthScopesResp struct {
	Base
	AuthUserField []string `json:"auth_user_field"`
	AuthOrgScopes struct {
		AuthedDept []int    `json:"authed_dept"`
		AuthedUser []string `json:"authed_user"`
	} `json:"auth_org_scopes"`
}

// OapiAuthScopesRequest 获取通讯录权限范围
func (d *DingtalkClient) OapiAuthScopesRequest() (*AuthScopesResp, error) {
	asr := &AuthScopesResp{}
	params := url.Values{}
	params.Set("access_token", d.AccessToken.Token)
	err := d.httpRequestWithStd("auth/scopes", params, nil, asr)
	if err != nil {
		return nil, err
	}
	return asr, nil
}
