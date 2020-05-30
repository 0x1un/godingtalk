package godingtalk

import (
	"strconv"
)

// OapiAuthScopesRequest 获取通讯录权限范围
func (d *DingtalkClient) OapiAuthScopesRequest() (AuthScopesResp, error) {
	var respData AuthScopesResp
	return respData, rpc(d, "auth/scopes", d.params, nil, &respData)
}

// OapiUserCreateRequest 创建一个用户 Method: POST
func (d *DingtalkClient) OapiUserCreateRequest(userInfo UserCreateReq) (UserCreateResp, error) {
	var respData UserCreateResp
	return respData, rpc(d, "user/create", d.params, userInfo, &respData)
}

// OapiUserUpdateRequest 更新用户的信息 Method: POST
func (d *DingtalkClient) OapiUserUpdateRequest(userInfo UserUpdateReq) (UserUpdateResp, error) {
	var respData UserUpdateResp
	return respData, rpc(d, "user/create", d.params, userInfo, &respData)
}

// OapiUserDeleteRequest 删除一个用户 Method: GET
func (d *DingtalkClient) OapiUserDeleteRequest(userid string) (UserDeleteResp, error) {
	params := d.params
	params.Set("userid", userid)
	var respData UserDeleteResp
	return respData, rpc(d, "user/delete", params, nil, &respData)
}

// OapiUserGetRequest 获取一个用户的详情 Method: GET
func (d *DingtalkClient) OapiUserGetRequest(userid string) (UserGetResp, error) {
	params := d.params
	params.Set("userid", userid)
	var respData UserGetResp
	return respData, rpc(d, "user/get", params, nil, &respData)
}

// OapiUserGetDeptMemberRequest 获取部门所有用户的userid Method: GET
func (d *DingtalkClient) OapiUserGetDeptMemberRequest(depID string) (UserGetDeptMemberResp, error) {
	params := d.params
	params.Set("deptId", depID)
	var respData UserGetDeptMemberResp
	return respData, rpc(d, "user/getDeptMember", params, nil, &respData)
}

/*
OapiUserSimplelistRequest 获取部门用户 Method: GET
depID: 部门id号
offset: 支持分页查询，与size参数同时设置时才生效，此参数代表偏移量
size: 支持分页查询，与offset参数同时设置时才生效，此参数代表分页大小，最大100
order:
	entry_asc：代表按照进入部门的时间升序，
	entry_desc：代表按照进入部门的时间降序，
	modify_asc：代表按照部门信息修改时间升序，
	modify_desc：代表按照部门信息修改时间降序，
	custom：代表用户定义(未定义时按照拼音)排序
*/
func (d *DingtalkClient) OapiUserSimplelistRequest(depID string, offset, size int64, order string) (UserSimplelistResp, error) {
	params := d.params
	params.Set("department_id", depID)
	params.Set("order", order)
	if size != 0 {
		params.Set("offset", strconv.FormatInt(offset, 10))
		params.Set("size", strconv.FormatInt(size, 10))
	}
	var respData UserSimplelistResp
	return respData, rpc(d, "user/simplelist", params, nil, &respData)
}

/*
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
*/
func (d *DingtalkClient) OapiUserListbypageRequest(depID string, offset, size int64, order string) (UserListbypageResp, error) {
	params := d.params
	params.Set("department_id", depID)
	params.Set("offset", strconv.FormatInt(offset, 10))
	params.Set("size", strconv.FormatInt(size, 10))
	params.Set("order", order)
	var respData UserListbypageResp
	return respData, rpc(d, "user/simplelist", params, nil, &respData)

}

// OapiUserGetAdminRequest 获取所有管理员
func (d *DingtalkClient) OapiUserGetAdminRequest() (UserGetAdminResp, error) {
	var respData UserGetAdminResp
	return respData, rpc(d, "user/get_admin", d.params, nil, &respData)
}

// OapiUserGetAdminScopeRequest 根据员工userid获取其所管理的部门 Method: GET
func (d *DingtalkClient) OapiUserGetAdminScopeRequest(userid string) (UserGetAdminScopeResp, error) {
	params := d.params
	params.Set("userid", userid)
	var respData UserGetAdminScopeResp
	return respData, rpc(d, "topapi/user/get_admin_scope", params, nil, &respData)
}

// OapiUserGetUseridByUnionidRequest 根据unionid获取userid Method: GET
// unionid: 员工的unionid
func (d *DingtalkClient) OapiUserGetUseridByUnionidRequest(unionid string) (UserGetUseridByUnionidResp, error) {
	params := d.params
	params.Set("unionid", unionid)
	var respData UserGetUseridByUnionidResp
	return respData, rpc(d, "user/getUseridByUnionid", params, nil, &respData)
}

// OapiUserGetByMobileRequest 根据手机号码获取用户userid Method: GET
// mobile: 手机号码
func (d *DingtalkClient) OapiUserGetByMobileRequest(mobile string) (UserGetByMobileResp, error) {
	params := d.params
	params.Set("mobile", mobile)
	var respData UserGetByMobileResp
	return respData, rpc(d, "user/get_by_mobile", params, nil, &respData)
}

// OapiUserGetOrgUserCountRequest 获取企业员工人数 Method: GET
// onlyActive: 0：包含未激活钉钉的人员数量 1：不包含未激活钉钉的人员数量
func (d *DingtalkClient) OapiUserGetOrgUserCountRequest(onlyActive string) (UserGetOrgUserCountResp, error) {
	params := d.params
	params.Set("onlyActive", onlyActive)
	var respData UserGetOrgUserCountResp
	return respData, rpc(d, "user/get_org_user_count", params, nil, &respData)
}

// OapiInactiveUserGetRequest 获取未登录的员工列表
// querData: 20200503 日期格式
// offset: 分页数据偏移量，从0开始
// size: 每页大小，最大100
func (d *DingtalkClient) OapiInactiveUserGetRequest(queryDate, offset, size string) (InactiveUserGetResp, error) {
	reqData := &RequestData{}
	var respData InactiveUserGetResp
	reqData.Set("query_date", queryDate)
	reqData.Set("offset", offset)
	reqData.Set("size", size)
	return respData, rpc(d, "topapi/inactive/user/get", d.params, reqData, &respData)
}
