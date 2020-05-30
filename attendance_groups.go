package godingtalk

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

func (d *DingtalkClient) OapiAttendanceGetusergroupRequest(userid string) (AttendanceGetusergroupResp, error) {
	reqData := struct {
		UserID string `json:"userid"`
	}{
		UserID: userid,
	}
	var respData AttendanceGetusergroupResp
	return respData, rpc(d, "topapi/attendance/getusergroup", d.params, reqData, &respData)
}

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
