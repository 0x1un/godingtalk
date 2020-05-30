package godingtalk

// OapiAttendanceListRecordRequest 获取考勤打卡详情
// userIds 企业内的员工id列表，最多不能超过50个
// checkDateFrom 查询考勤打卡记录的起始工作日。格式为“yyyy-MM-dd hh:mm:ss”。
// checkDateTo 查询考勤打卡记录的结束工作日。格式为“yyyy-MM-dd hh:mm:ss”。注意，起始与结束工作日最多相隔7天
// isI18n 取值为true和false，表示是否为海外企业使用，默认为false。其中，true：海外平台使用，false：国内平台使用
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

// OapiAttendanceListRequest 获取考勤打卡结果
// workDateFrom 查询考勤打卡记录的起始工作日。格式为“yyyy-MM-dd HH:mm:ss”，HH:mm:ss可以使用00:00:00，将返回此日期从0点到24点的结果
// workDateTo 查询考勤打卡记录的结束工作日。格式为“yyyy-MM-dd HH:mm:ss”，HH:mm:ss可以使用00:00:00，将返回此日期从0点到24点的结果。注意，起始与结束工作日最多相隔7天
// userIdList 员工在企业内的UserID列表，企业用来唯一标识用户的字段。最多不能超过50个
// offset 表示获取考勤数据的起始点，第一次传0，如果还有多余数据，下次获取传的offset值为之前的offset+limit，0、1、2...依次递增
// limit 表示获取考勤数据的条数，最大不能超过50条
// isI18n 取值为true和false，表示是否为海外企业使用，默认为false。其中，true：海外平台使用，false：国内平台使用
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
