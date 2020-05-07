package godingtalk

import "strconv"

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

// OapiDepartmentListRequest 获取部门列表 Method: GET
// id 父部门id（如果不传，默认部门为根部门，根部门ID为1）
// fetchChild 是否递归部门的全部子部门
func (d *DingtalkClient) OapiDepartmentListRequest(id string, fetchChild bool) (DepartmentListResp, error) {
	params := d.params
	params.Set("id", id)
	params.Set("fetch_child", strconv.FormatBool(fetchChild))
	var respData DepartmentListResp
	err := d.httpRequestWithStd("department/list", params, nil, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}

func (d *DingtalkClient) OapiDepartmentGetRequest(id string) (DepartmentGetResp, error) {
	params := d.params
	params.Set("id", id)
	var respData DepartmentGetResp
	err := d.httpRequestWithStd("department/get", params, nil, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}

func (d *DingtalkClient) OapiDepartmentListParentDeptsByDeptRequest(id string) (DepartmentListParentDeptsByDeptResp, error) {
	params := d.params
	params.Set("id", id)
	var respData DepartmentListParentDeptsByDeptResp
	err := d.httpRequestWithStd("department/list_parent_depts_by_dept", params, nil, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}

func (d *DingtalkClient) OapiDepartmentListParentDeptsRequest(userid string) (DepartmentListParentDeptsResp, error) {
	params := d.params
	params.Set("userId", userid)
	var respData DepartmentListParentDeptsResp
	err := d.httpRequestWithStd("department/list_parent_depts", params, nil, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}
