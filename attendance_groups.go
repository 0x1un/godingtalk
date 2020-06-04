package godingtalk

import "strings"

// OapiAttendanceGetsimplegroupsRequest 批量获取企业考勤组详情
// offset 偏移位置，从0、1、2...依次递增，默认值：0
// size 分页大小，最大10，默认值：10
func (d *DingtalkClient) OapiAttendanceGetsimplegroupsRequest(offset int64, size int64) (AttendanceGetsimplegroupsResp, error) {
	reqData := AttendanceGetsimplegroupsReq{
		Offset: offset,
		Size:   size,
	}
	var respData AttendanceGetsimplegroupsResp
	return respData, rpc(d, "topapi/attendance/getsimplegroups", d.params, reqData, &respData)
}

// OapiAttendanceGroupMinimalismListRequest 获取考勤组摘要
// 分页获取企业内所有考勤组摘要信息
func (d *DingtalkClient) OapiAttendanceGroupMinimalismListRequest(opUserID string, cursor int64) (AttendanceGroupMinimalismListResp, error) {
	reqData := struct {
		OpUserID string `json:"op_user_id"`
		Cursor   int64  `json:"cursor"`
	}{
		OpUserID: opUserID,
		Cursor:   cursor,
	}
	var respData AttendanceGroupMinimalismListResp
	return respData, rpc(d, "topapi/attendance/group/minimalism/list", d.params, reqData, &respData)
}

// OapiAttendanceGroupSearchRequest 搜索考勤组摘要
func (d *DingtalkClient) OapiAttendanceGroupSearchRequest(opUserID, groupName string) (AttendanceGroupSearchResp, error) {
	reqData := struct {
		OpUserID  string `json:"op_user_id"`
		GroupName string `json:"group_name"`
	}{
		OpUserID:  opUserID,
		GroupName: groupName,
	}
	var respData AttendanceGroupSearchResp
	return respData, rpc(d, "topapi/attendance/group/search", d.params, reqData, &respData)

}

// OapiAttendanceGroupQueryRequest 获取考勤组详情
func (d *DingtalkClient) OapiAttendanceGroupQueryRequest(opUserID string, groupID int) (AttendanceGroupQueryResp, error) {
	reqData := struct {
		OpUserID string `json:"op_user_id"`
		GroupID  int    `json:"group_id"`
	}{
		OpUserID: opUserID,
		GroupID:  groupID,
	}
	var respData AttendanceGroupQueryResp
	return respData, rpc(d, "topapi/attendance/group/query", d.params, reqData, &respData)
}

// OapiAttendanceGetusergroupRequest 获取用户考勤组
func (d *DingtalkClient) OapiAttendanceGetusergroupRequest(userid string) (AttendanceGetusergroupResp, error) {
	reqData := struct {
		UserID string `json:"userid"`
	}{
		UserID: userid,
	}
	var respData AttendanceGetusergroupResp
	return respData, rpc(d, "topapi/attendance/getusergroup", d.params, reqData, &respData)
}
// OapiAttendanceGroupMemberusersListRequest 获取考勤组员工id
func (d *DingtalkClient) OapiAttendanceGroupMemberusersListRequest(opUserID string, cursor, groupID int64) (AttendanceGroupMemberuserListResp, error) {
	reqData := struct {
		OpUserID string `json:"op_user_id"`
		Cursor   int64  `json:"cursor"`
		GroupID  int64  `json:"group_id"`
	}{
		opUserID,
		cursor,
		groupID,
	}
	var respData AttendanceGroupMemberuserListResp
	return respData, rpc(d, "topapi/attendance/group/memberusers/list", d.params, reqData, &respData)
}

// OapiAttendanceGroupMemberListbyidsRequest 考勤组成员校验
// 校验某个部门或者员工是否属于某个考勤组，返回值为属于这个考勤组的部门id或者员工id
func (d *DingtalkClient) OapiAttendanceGroupMemberListbyidsRequest(opUserid string,
	memberIds []string, memberType, groupID int64) (AttendanceGroupMemberListbyidsResp, error){
	reqData := struct {
		OpUserID string `json:"op_user_id"`
		MemberIDs string `json:"member_ids"`
		MemberType int64 `json:"member_type"`
		GroupID int64 `json:"group_id"`
	}{
		opUserid,
		strings.Join(memberIds, ","),
		memberType,
		groupID,
	}
	var respData AttendanceGroupMemberListbyidsResp
	return respData, rpc(d, "topapi/attendance/group/member/listbyids", d.params, reqData, &respData)
}

// OapiAttendanceGroupMemberUpdateRequest 考勤组成员更新
func (d *DingtalkClient) OapiAttendanceGroupMemberUpdateRequest(reqData AttendanceGroupMemberUpdateReq)(AttendanceGroupMemberUpdateResp, error) {
	var respData AttendanceGroupMemberUpdateResp
	return respData, rpc(d, "topapi/attendance/group/member/update", d.params, reqData, &respData)
}

// OapiAttendanceGroupMemberListRequest 获取考勤组成员
// 注意:如果某个员工A属于人力资源部门，那么通过这个接口是无法查询到A的，该接口仅能获取到人力资源部和欢欢两个成员信息
func (d *DingtalkClient) OapiAttendanceGroupMemberListRequest(opUserID string, cursor, groupID int64)(AttendanceGroupMemberListResp, error) {
	reqData := struct {
		OpUserID string `json:"op_user_id"`
		Cursor int64 `json:"cursor"`
		GroupID int64 `json:"group_id"`
	}{
		opUserID,
		cursor,
		groupID,
	}
	var respData AttendanceGroupMemberListResp
	return respData, rpc(d, "topapi/attendance/group/member/list", d.params, reqData, &respData)
}