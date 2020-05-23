// doc: https://ding-doc.dingtalk.com/doc#/serverapi2/potcn9b
// doc: https://ding-doc.dingtalk.com/doc#/serverapi2/ul33mm
package godingtalk

type AttendanceListRecordReq struct {
	UserIds       []string `json:"userIds"`
	CheckDateFrom string   `json:"checkDateFrom"`
	CheckDateTo   string   `json:"checkDateTo"`
	IsI18N        bool     `json:"isI18n"`
}

type AttendanceListRecordResp struct {
	Base
	Recordresult []struct {
		// 是否合法，当timeResult和locationResult都为Normal时，该值为Y；否则为N
		IsLegal string `json:"isLegal"`
		// 排班打卡时间
		BaseCheckTime int64 `json:"baseCheckTime"`
		// 唯一标识ID
		ID int `json:"id"`
		// 用户打卡地址
		UserAddress string `json:"userAddress"`
		// 用户ID
		UserID string `json:"userId"`
		//  考勤类型，OnDuty：上班 OffDuty：下班
		CheckType string `json:"checkType"`
		// 时间结果，Normal：正常; Early：早退; Late：迟到; SeriousLate：严重迟到； Absenteeism：旷工迟到； NotSigned：未打卡
		TimeResult string `json:"timeResult"`
		// 设备id
		DeviceID string `json:"deviceId"`
		// 企业ID
		CorpID string `json:"corpId"`
		// 数据来源，ATM：考勤机;BEACON：IBeacon;DING_ATM：钉钉考勤机;USER：用户打卡;BOSS：老板改签;APPROVE：审批系统;SYSTEM：考勤系统;AUTO_CHECK：自动打卡
		SourceType string `json:"sourceType"`
		// 工作日
		WorkDate int64 `json:"workDate"`
		// 排班打卡时间
		PlanCheckTime int64 `json:"planCheckTime"`
		// 定位方法
		LocationMethod string `json:"locationMethod"`
		// 位置结果，Normal：范围内 Outside：范围外，外勤打卡时为这个值
		LocationResult string `json:"locationResult"`
		// 用户打卡经度
		UserLongitude float64 `json:"userLongitude"`
		// 排班ID
		PlanID int64 `json:"planId"`
		// 考勤组ID
		GroupID int64 `json:"groupId"`
		// 用户打卡定位精度
		UserAccuracy float64 `json:"userAccuracy"`
		// 实际打卡时间
		UserCheckTime int64 `json:"userCheckTime"`
		// 用户打卡纬度
		UserLatitude float64 `json:"userLatitude"`
		// 关联的审批实例id，当该字段非空时，表示打卡记录与请假、加班等审批有关。可以与获取单个审批数据配合使用
		ProcInstID string `json:"procInstId"`
	} `json:"recordresult"`
}

type AttendanceListResp struct {
	Base
	HasMore      bool `json:"hasMore"`
	Recordresult []struct {
		// 计算迟到和早退，基准时间
		BaseCheckTime int64 `json:"baseCheckTime"`
		// 考勤类型 OnDuty：上班 OffDuty：下班
		CheckType string `json:"checkType"`
		// 企业ID
		CorpID string `json:"corpId"`
		// 考勤组ID
		GroupID int `json:"groupId"`
		// 唯一标识ID
		ID int `json:"id"`
		//  	位置结果 Normal：范围内；Outside：范围外；NotSigned：未打卡
		LocationResult string `json:"locationResult"`
		// 排班ID
		PlanID int `json:"planId"`
		// 打卡记录ID
		RecordID int `json:"recordId,omitempty"`
		// 时间结果 Normal：正常; Early：早退; Late：迟到; SeriousLate：严重迟到； Absenteeism：旷工迟到； NotSigned：未打卡
		TimeResult string `json:"timeResult"`
		// 实际打卡时间,  用户打卡时间的毫秒数
		UserCheckTime int64 `json:"userCheckTime"`
		// 用户ID
		UserID string `json:"userId"`
		// 工作日
		WorkDate int64 `json:"workDate"`
		// 关联的审批实例id，当该字段非空时，表示打卡记录与请假、加班等审批有关。可以与获取单个审批数据配合使用
		ProcInstID string `json:"procInstId"`
	} `json:"recordresult"`
}

type AttendanceListReq struct {
	WorkDateFrom string   `json:"workDateFrom"`
	WorkDateTo   string   `json:"workDateTo"`
	UserIDList   []string `json:"userIdList"`
	Offset       int64    `json:"offset"`
	Limit        int64    `json:"limit"`
}
