package godingtalk

func (d *DingtalkClient) OapiAttendanceShiftListRequest(opUserID string, cursor int64) (AttendanceShiftListResp, error) {
	reqData := AttendanceShiftListReq{
		OpUserID: opUserID,
		Cursor:   cursor,
	}
	var respData AttendanceShiftListResp
	return respData, rpc(d, "topapi/attendance/shift/list", d.params, reqData, &respData)
}

func (d *DingtalkClient) OapiAttendanceShiftQueryRequest(opUserID string, shiftID int64) (AttendanceShiftQueryResp, error) {
	reqData := AttendanceShiftQueryReq{
		OpUserID: opUserID,
		ShiftID:  shiftID,
	}
	var respData AttendanceShiftQueryResp
	return respData, rpc(d, "topapi/attendance/shift/query", d.params, reqData, &respData)
}

func (d *DingtalkClient) OapiAttendanceShiftSearchRequest(opUserID, shiftName string) (AttendanceShiftSearchResp, error) {
	reqData := AttendanceShiftSearchReq{
		OpUserID:  opUserID,
		ShiftName: shiftName,
	}
	var respData AttendanceShiftSearchResp
	return respData, rpc(d, "topapi/attendance/shift/search", d.params, reqData, &respData)
}
