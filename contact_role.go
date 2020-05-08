package godingtalk

import "encoding/json"

type RoleListRequest struct {
	Size   int `json:"size"`
	Offset int `json:"offset"`
}

func (r RoleListRequest) ToBytes() ([]byte, error) {
	return json.Marshal(r)
}

func (d *DingtalkClient) OapiRoleListRequest(reqData RoleListRequest) (RoleListResp, error) {
	var respData RoleListResp
	err := d.httpRequestWithStd("topapi/role/list", d.params, reqData, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}
