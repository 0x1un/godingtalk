package godingtalk



func (d *DingtalkClient) OapiDepartmentCreateRequest(reqData DepartmentCreateReq) (DepartmentCreateResp, error) {
	var respData DepartmentCreateResp
	err := d.httpRequestWithStd("department/create", d.params, reqData, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}