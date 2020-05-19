package godingtalk

type RoleListRequest struct {
	Size   int `json:"size"`
	Offset int `json:"offset"`
}

func (d *DingtalkClient) OapiRoleListRequest(size, offset int) (RoleListResp, error) {
	reqData := RoleListRequest{
		Size:   size,
		Offset: offset,
	}
	var respData RoleListResp
	return respData, rpc(d, "topapi/role/list", d.params, reqData, &respData)
}

func (d *DingtalkClient) OapiRoleSimplelistRequest(roleID string) (RoleSimplelistResp, error) {
	params := d.params
	var respData RoleSimplelistResp
	return respData, rpc(d, "topapi/role/simplelist", params, nil, &respData)
}
