package godingtalk

import (
	"encoding/json"
)

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
		return respData, err
	}
	return respData, nil
}

//  OapiUserGetDeptMemberRequest 获取部门所有用户的userid Method: GET
func (d *DingtalkClient) OapiUserGetDeptMemberRequest(depID string) (UserGetDeptMemberResp, error) {
	params := d.params
	params.Set("deptId", depID)
	var respData UserGetDeptMemberResp
	err := d.httpRequestWithStd("user/getDeptMember", params, nil, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}

// OapiUserSimplelistRequest 获取部门用户 Method: GET
func (d *DingtalkClient) OapiUserSimplelistRequest(depID, offset, size string) (UserSimplelistResp, error) {
	params := d.params
	params.Set("department_id", depID)
	params.Set("offset", offset)
	params.Set("size", size)
	var respData UserSimplelistResp
	err := d.httpRequestWithStd("user/simplelist", params, nil, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}

/****************
OapiUserListbypageRequest 获取部门用户的详情 Method: GET
depID: 部门id，1表示根部门
offset: 支持分页查询，与size参数同时设置时才生效，此参数代表偏移量，偏移量从0开始
size: 支持分页查询，与offset参数同时设置时才生效，此参数代表分页大小，最大100
order:
	entry_asc：代表按照进入部门的时间升序，
	entry_desc：代表按照进入部门的时间降序，
	modify_asc：代表按照部门信息修改时间升序，
	modify_desc：代表按照部门信息修改时间降序，
	custom：代表用户定义(未定义时按照拼音)排序
****************/
func (d *DingtalkClient) OapiUserListbypageRequest(depID, offset, size, order string) (UserListbypageResp, error) {
	params := d.params
	params.Set("department_id", depID)
	params.Set("offset", offset)
	params.Set("size", size)
	params.Set("order", order)
	var respData UserListbypageResp
	err := d.httpRequestWithStd("user/simplelist", params, nil, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil

}

func (d *DingtalkClient) OapiUserGetAdminRequest() (UserGetAdminResp, error) {
	var respData UserGetAdminResp
	err := d.httpRequestWithStd("user/get_admin", d.params, nil, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}

// OapiUserGetAdminScopeRequest 根据员工userid获取其所管理的部门 Method: GET
func (d *DingtalkClient) OapiUserGetAdminScopeRequest(userid string) (UserGetAdminScopeResp, error) {
	params := d.params
	params.Set("userid", userid)
	var respData UserGetAdminScopeResp
	err := d.httpRequestWithStd("topapi/user/get_admin_scope", params, nil, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}

// OapiUserGetUseridByUnionidRequest 根据unionid获取userid Method: GET
func (d *DingtalkClient) OapiUserGetUseridByUnionidRequest(unionid string) (UserGetUseridByUnionidResp, error) {
	params := d.params
	params.Set("unionid", unionid)
	var respData UserGetUseridByUnionidResp
	err := d.httpRequestWithStd("user/getUseridByUnionid", params, nil, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}

// OapiUserGetByMobileRequest 根据手机号码获取用户userid Method: GET
func (d *DingtalkClient) OapiUserGetByMobileRequest(mobile string) (UserGetByMobileResp, error) {
	params := d.params
	params.Set("mobile", mobile)
	var respData UserGetByMobileResp
	err := d.httpRequestWithStd("user/get_by_mobile", params, nil, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}

func (d *DingtalkClient) OapiUserGetOrgUserCountRequest(onlyActive string) (UserGetOrgUserCountResp, error) {
	params := d.params
	params.Set("onlyActive", onlyActive)
	var respData UserGetOrgUserCountResp
	err := d.httpRequestWithStd("user/get_org_user_count", params, nil, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}

// OapiInactiveUserGetRequest 获取未登录的员工列表
// querData: 20200503 日期格式
// offset: 分页数据偏移量，从0开始
// size: 每页大小，最大100
func (d *DingtalkClient) OapiInactiveUserGetRequest(queryData, offset, size string) (InactiveUserGetResp, error) {
	reqData := &RequestData{}
	var respData InactiveUserGetResp
	reqData.Set("query_data", queryData)
	reqData.Set("offset", offset)
	reqData.Set("size", size)
	err := d.httpRequestWithStd("topapi/inactive/user/get", d.params, reqData, &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}
