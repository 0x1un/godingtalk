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
		IsLegal        string  `json:"isLegal"`
		BaseCheckTime  int64   `json:"baseCheckTime"`
		ID             int     `json:"id"`
		UserAddress    string  `json:"userAddress"`
		UserID         string  `json:"userId"`
		CheckType      string  `json:"checkType"`
		TimeResult     string  `json:"timeResult"`
		DeviceID       string  `json:"deviceId"`
		CorpID         string  `json:"corpId"`
		SourceType     string  `json:"sourceType"`
		WorkDate       int64   `json:"workDate"`
		PlanCheckTime  int64   `json:"planCheckTime"`
		LocationMethod string  `json:"locationMethod"`
		LocationResult string  `json:"locationResult"`
		UserLongitude  float64 `json:"userLongitude"`
		PlanID         int64   `json:"planId"`
		GroupID        int64   `json:"groupId"`
		UserAccuracy   float64 `json:"userAccuracy"`
		UserCheckTime  int64   `json:"userCheckTime"`
		UserLatitude   float64 `json:"userLatitude"`
		ProcInstID     string  `json:"procInstId"`
	} `json:"recordresult"`
}

type AttendanceListResp struct {
	Base
	HasMore      bool `json:"hasMore"`
	Recordresult []struct {
		BaseCheckTime  int64  `json:"baseCheckTime"`
		CheckType      string `json:"checkType"`
		CorpID         string `json:"corpId"`
		GroupID        int    `json:"groupId"`
		ID             int    `json:"id"`
		LocationResult string `json:"locationResult"`
		PlanID         int    `json:"planId"`
		RecordID       int    `json:"recordId,omitempty"`
		TimeResult     string `json:"timeResult"`
		UserCheckTime  int64  `json:"userCheckTime"`
		UserID         string `json:"userId"`
		WorkDate       int64  `json:"workDate"`
		ProcInstID     string `json:"procInstId"`
	} `json:"recordresult"`
}

type AttendanceListReq struct {
	WorkDateFrom string   `json:"workDateFrom"`
	WorkDateTo   string   `json:"workDateTo"`
	UserIDList   []string `json:"userIdList"`
	Offset       int64    `json:"offset"`
	Limit        int64    `json:"limit"`
}
