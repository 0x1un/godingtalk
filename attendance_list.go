package godingtalk

func (d *DingtalkClient) OapiAttendanceListRecordRequest(userIDs []string, checkDateFrom string, checkDateTo string, isI18n bool) (AttendanceListRecordResp, error) {
	reqData := AttendanceListRecordReq{
		UserIds:       userIDs,
		CheckDateFrom: checkDateFrom,
		CheckDateTo:   checkDateTo,
		IsI18N:        isI18n,
	}
	var respData AttendanceListRecordResp
	return respData, rpc(d, "attendance/listRecord", d.params, reqData, &respData)
}

func (d *DingtalkClient) OapiAttendanceListRequest(userIDList []string, workDateFrom, workDateTo string, offset, limit int64) (AttendanceListResp, error) {
	reqData := AttendanceListReq{
		WorkDateFrom: workDateFrom,
		WorkDateTo:   workDateTo,
		UserIDList:   userIDList,
		Offset:       offset,
		Limit:        limit,
	}
	var respData AttendanceListResp
	return respData, rpc(d, "attendance/list", d.params, reqData, &respData)
}
