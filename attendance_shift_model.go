package godingtalk

type AttendanceShiftListResp struct {
	Result struct {
		HasMore bool        `json:"has_more"`
		Cursor  interface{} `json:"cursor"`
		Result  []struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"result"`
	} `json:"result"`
	Base
}

type AttendanceShiftListReq struct {
	OpUserID string `json:"op_user_id"`
	Cursor   int64  `json:"cursor"`
}

type AttendanceShiftQueryResp struct {
	Result struct {
		ShiftGroupName string `json:"shift_group_name"`
		CorpID         string `json:"corp_id"`
		ShiftSetting   struct {
			ShiftID         int    `json:"shift_id"`
			GmtModified     string `json:"gmt_modified"`
			CorpID          string `json:"corp_id"`
			IsDeleted       string `json:"is_deleted"`
			WorkTimeMinutes int    `json:"work_time_minutes"`
			ID              int    `json:"id"`
			AttendDays      string `json:"attend_days"`
			GmtCreate       string `json:"gmt_create"`
		} `json:"shift_setting"`
		Name     string `json:"name"`
		ID       int    `json:"id"`
		Sections []struct {
			Punches []struct {
				CheckType     string `json:"check_type"`
				EndMin        int    `json:"end_min"`
				Across        int    `json:"across"`
				CheckTime     string `json:"check_time"`
				PermitMinutes int    `json:"permit_minutes"`
				FreeCheck     bool   `json:"free_check"`
				ID            int    `json:"id"`
				BeginMin      int    `json:"begin_min"`
			} `json:"punches"`
			WorkTimeMinutes int `json:"work_time_minutes"`
			Rests           []struct {
				CheckType string `json:"check_type"`
				Across    int    `json:"across"`
				CheckTime string `json:"check_time"`
				ID        int    `json:"id"`
			} `json:"rests"`
			ID int `json:"id"`
		} `json:"sections"`
		ShiftGroupID int `json:"shift_group_id"`
	} `json:"result"`
	Base
}

type AttendanceShiftQueryReq struct {
	OpUserID string `json:"op_user_id"`
	ShiftID  int64  `json:"shift_id"`
}

type AttendanceShiftSearchResp struct {
	Result []struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"result"`
	Base
}

type AttendanceShiftSearchReq struct {
	OpUserID  string `json:"op_user_id"`
	ShiftName string `json:"shift_name"`
}
