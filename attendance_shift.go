package godingtalk

// OapiAttendanceShiftListRequest 批量查询班次摘要信息
func (d *DingtalkClient) OapiAttendanceShiftListRequest(opUserID string, cursor int64) (AttendanceShiftListResp, error) {
	reqData := AttendanceShiftListReq{
		OpUserID: opUserID,
		Cursor:   cursor,
	}
	var respData AttendanceShiftListResp
	return respData, rpc(d, "topapi/attendance/shift/list", d.params, reqData, &respData)
}

// OapiAttendanceShiftQueryRequest 查询班次详情
func (d *DingtalkClient) OapiAttendanceShiftQueryRequest(opUserID string, shiftID int64) (AttendanceShiftQueryResp, error) {
	reqData := AttendanceShiftQueryReq{
		OpUserID: opUserID,
		ShiftID:  shiftID,
	}
	var respData AttendanceShiftQueryResp
	return respData, rpc(d, "topapi/attendance/shift/query", d.params, reqData, &respData)
}

// OapiAttendanceShiftSearchRequest 按名称搜索班次
// op_user_id 操作人userId
// shift_name 班次名称
func (d *DingtalkClient) OapiAttendanceShiftSearchRequest(opUserID, shiftName string) (AttendanceShiftSearchResp, error) {
	reqData := AttendanceShiftSearchReq{
		OpUserID:  opUserID,
		ShiftName: shiftName,
	}
	var respData AttendanceShiftSearchResp
	return respData, rpc(d, "topapi/attendance/shift/search", d.params, reqData, &respData)
}
