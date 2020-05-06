package godingtalk

// OapiDepartmentCreateRequest 创建一个部门 Method: POST
func (d *DingtalkClient) OapiDepartmentCreateRequest(reqData DepartmentCreateReq) (DepartmentCreateResp, error) {
	var respData DepartmentCreateResp
	err := d.httpRequestWithStd("department/create", d.params, reqData, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}

// OapiDepartmentDeleteRequest 删除一个部门 Method: GET
func (d *DingtalkClient) OapiDepartmentDeleteRequest(id string) (DepartmentDeleteResp, error) {
	params := d.params
	params.Set("id", id)
	var respData DepartmentDeleteResp
	err := d.httpRequestWithStd("department/delete", params, nil, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}

// OapiDepartmentListIdsRequest 获取子部门ID列表 Method: GET
func (d *DingtalkClient) OapiDepartmentListIdsRequest(id string) (DepartmentListIdsResp, error) {
	params := d.params
	params.Set("id", id)
	var respData DepartmentListIdsResp
	err := d.httpRequestWithStd("department/list_ids", params, nil, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}
