package godingtalk

import (
	"encoding/json"
)

type AttendanceListscheduleResp struct {
	Result struct {
		Schedules []struct {
			PlanID           int    `json:"plan_id"`            // 排班id
			CheckType        string `json:"check_type"`         // 打卡类型，OnDuty表示上班打卡，OffDuty表示下班打卡
			ApproveID        int    `json:"approve_id"`         // 审批id，结果集中没有的话表示没有审批单
			Userid           string `json:"userid"`             // userId
			ClassID          int    `json:"class_id"`           // 考勤班次id
			ClassSettingID   int    `json:"class_setting_id"`   // 班次配置id，结果集中没有的话表示使用全局班次配置
			PlanCheckTime    string `json:"plan_check_time"`    // 打卡时间
			GroupID          int    `json:"group_id"`           // 调整后的打卡时间
			ChangedCheckTime string `json:"changed_check_time"` // 考勤组id
		} `json:"schedules"` // 排班列表
		HasMore bool `json:"has_more"` // 分页用，表示是否还有下一页
	} `json:"result"`
	Base
}

type AttendanceListscheduleReq struct {
	WorkDate string `json:"workDate"`
	Offset   int    `json:"offset"`
	Size     int    `json:"size"`
}

func (a AttendanceListscheduleReq) ToBytes() ([]byte, error) {
	return json.Marshal(a)
}

type AttendanceScheduleListbydayResp struct {
	Result []struct {
		// 卡点类型 OnDuty：上班 OffDuty：下班
		CheckType string `json:"check_type"`
		// 最后更新时间
		GmtModified string `json:"gmt_modified"`
		// 计划打卡时间
		PlanCheckTime string `json:"plan_check_time"`
		// 企业id
		CorpID string `json:"corp_id"`
		// 打卡时间
		CheckDateTime string `json:"check_date_time"`
		// 允许迟到早退等规则调整后的计划打卡时间
		BaseCheckTime string `json:"base_check_time"`
		// 考勤组id
		GroupID int `json:"group_id"`
		// 班次名称
		ClassName string `json:"class_name"`
		// 创建时间
		GmtCreate string `json:"gmt_create"`
		// 员工id
		UserID string `json:"user_id"`
		// 排班绑定的假勤审批类型 1：加班，2：出差，3：请假
		ApproveBizType int `json:"approve_biz_type"`
		// 排班绑定的审批单id
		ApproveID int `json:"approve_id"`
		//  排班关联的班次设置id
		ClassSettingID int `json:"class_setting_id"`
		// 排班绑定的假勤审批单名称
		ApproveTagName string `json:"approve_tag_name"`
		// 扩展字段
		Features string `json:"features"`
		// 排班绑定的班次id，该字段为空或者小于0时代表当天未排班
		ClassID int `json:"class_id"`
		// 卡点状态 Init：未打卡 Checked：已打卡 Timeout：缺卡
		CheckStatus string `json:"check_status"`
		// 工作日，代表具体哪一天的排班
		WorkDate string `json:"work_date"`
		// 结束打卡时间
		CheckEndTime string `json:"check_end_time"`
		// 开始打卡时间
		IsRest string `json:"is_rest"`
		// Y:当天排休
		CheckBeginTime string `json:"check_begin_time"`
		// 排班id
		ID int `json:"id"`
	} `json:"result"`
	Success bool `json:"success"`
	Base
}

type AttendanceScheduleListbydayReq struct {
	OpUserID string `json:"op_user_id"`
	UserID   string `json:"user_id"`
	DateTime int64  `json:"date_time"`
}

func (a AttendanceScheduleListbydayReq) ToBytes() ([]byte, error) {
	return json.Marshal(a)
}

type AttendanceScheduleListbyusersResp AttendanceScheduleListbydayResp

type AttendanceScheduleListbyusersReq struct {
	OpUserID     string `json:"op_user_id"`
	UserIDs      string `json:"userids"`
	FromDateTime int64  `json:"from_date_time"`
	ToDateTime   int64  `json:"to_date_time"`
}

func (a AttendanceScheduleListbyusersReq) ToBytes() ([]byte, error) {
	return json.Marshal(a)
}

type AttendanceScheduleResultListbyidsResp struct {
	Result []struct {
		CheckType      string `json:"check_type"`
		GmtModified    string `json:"gmt_modified"`
		PlanCheckTime  string `json:"plan_check_time"`
		CorpID         string `json:"corp_id"`
		BaseCheckTime  string `json:"base_check_time"`
		GroupID        int    `json:"group_id"`
		GmtCreate      string `json:"gmt_create"`
		UserID         string `json:"user_id"`
		WorkDate       string `json:"work_date"`
		ID             int64  `json:"id"`
		LocationResult string `json:"location_result"`
		IsLegal        string `json:"is_legal"`
		TimeResult     string `json:"time_result"`
		RecordID       int    `json:"record_id"`
		UserCheckTime  string `json:"user_check_time"`
		ScheduleID     int64  `json:"schedule_id"`
	} `json:"result"`
	Success bool `json:"success"`
	Base
}

type AttendanceScheduleResultListbyidsReq struct {
	OpUserID    string `json:"op_user_id"`
	ScheduleIDs string `json:"schedule_ids"`
}

func (a AttendanceScheduleResultListbyidsReq) ToBytes() ([]byte, error) {
	return json.Marshal(a)
}
