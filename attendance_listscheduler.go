package godingtalk

import "strings"

// OapiAttendanceListscheduleRequest 查询企业考勤排班详情
// workDate 取值为日期 2020-01-01 00:00:00
// offset 偏移位置，从0开始，后续传offset+size的值。当返回结果中的has_more为false时，表示没有多余的数据了。
// size 分页大小，最大200，默认值：200
func (d *DingtalkClient) OapiAttendanceListscheduleRequest(workDate string, offset, size int) (AttendanceListscheduleResp, error) {
	reqData := AttendanceListscheduleReq{
		WorkDate: workDate,
		Offset:   offset,
		Size:     size,
	}
	var respData AttendanceListscheduleResp
	return respData, rpc(d, "topapi/attendance/listschedule", d.params, reqData, &respData)
}

// OapiAttendanceScheduleListbydayRequest 查询成员排班信息
// opUserID 操作人userId
// userID 用户userId
// dateTime 查询某个工作日的数据，Unix时间戳 (毫秒)
func (d *DingtalkClient) OapiAttendanceScheduleListbydayRequest(opUserID, userID string, dateTime int64) (AttendanceScheduleListbydayResp, error) {
	reqData := AttendanceScheduleListbydayReq{
		OpUserID: opUserID,
		UserID:   userID,
		DateTime: dateTime,
	}
	var respData AttendanceScheduleListbydayResp
	return respData, rpc(d, "topapi/attendance/schedule/listbyday", d.params, reqData, &respData)

}

// OapiAttendanceScheduleListbyusersRequest 批量查询成员排班信息
// op_user_id 操作人userId
// userids 用户userId列表
// from_date_time 起始日期(毫秒)
// to_date_time 结束日期(毫秒)
func (d *DingtalkClient) OapiAttendanceScheduleListbyusersRequest(opUserID, userIDs string, fromDateTime, toDateTime int64) (AttendanceScheduleListbyusersResp, error) {
	reqData := AttendanceScheduleListbyusersReq{
		OpUserID:     opUserID,
		UserIDs:      userIDs,
		FromDateTime: fromDateTime,
		ToDateTime:   toDateTime,
	}
	var respData AttendanceScheduleListbyusersResp
	return respData, rpc(d, "topapi/attendance/schedule/listbyusers", d.params, reqData, &respData)
}

// OapiAttendanceScheduleResultListbyidsRequest 查询排班打卡结果
func (d *DingtalkClient) OapiAttendanceScheduleResultListbyidsRequest(opUserID string, scheduleIDs string) (AttendanceScheduleResultListbyidsResp, error) {
	reqData := AttendanceScheduleResultListbyidsReq{
		OpUserID:    opUserID,
		ScheduleIDs: scheduleIDs,
	}
	var respData AttendanceScheduleResultListbyidsResp
	return respData, rpc(d, "topapi/attendance/schedule/result/listbyids", d.params, reqData, &respData)
}

func (d *DingtalkClient) OapiAttendanceGetleavestatusRequest(useridList []string, startTime, endTime, offset, size int64) (AttendanceGetleavestatusResp, error) {
	reqData := struct {
		UseridList string `json:"userid_list"`
		StartTime int64 `json:"start_time"`
		EndTime int64 `json:"end_time"`
		Offset int64 `json:"offset"`
		Size int64 `json:"size"`
	}{
		strings.Join(useridList, ","),
		startTime,
		endTime,
		offset,
		size,
	}
	var respData AttendanceGetleavestatusResp
	return respData, rpc(d, "topapi/attendance/getleavestatus", d.params, reqData, &respData)
}
