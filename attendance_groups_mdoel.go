package godingtalk

type AttendanceGetsimplegroupsResp struct {
	Base
	Result struct {
		HasMore bool `json:"has_more"`
		Groups  []struct {
			GroupID       int    `json:"group_id"`
			IsDefault     bool   `json:"is_default"`
			GroupName     string `json:"group_name"`
			SelectedClass []struct {
				Setting struct {
					ClassSettingID int `json:"class_setting_id"`
					RestBeginTime  struct {
						CheckTime string `json:"check_time"`
					} `json:"rest_begin_time"`
					PermitLateMinutes int `json:"permit_late_minutes"`
					WorkTimeMinutes   int `json:"work_time_minutes"`
					RestEndTime       struct {
						CheckTime string `json:"check_time"`
					} `json:"rest_end_time"`
					AbsenteeismLateMinutes int    `json:"absenteeism_late_minutes"`
					SeriousLateMinutes     int    `json:"serious_late_minutes"`
					IsOffDutyFreeCheck     string `json:"is_off_duty_free_check"`
				} `json:"setting"`
				ClassID  int `json:"class_id"`
				Sections []struct {
					Times []struct {
						CheckTime string `json:"check_time"`
						CheckType string `json:"check_type"`
						Across    int    `json:"across"`
					} `json:"times"`
				} `json:"sections"`
				ClassName string `json:"class_name"`
			} `json:"selected_class"`
			Type           string   `json:"type"`
			MemberCount    int      `json:"member_count"`
			DefaultClassID int      `json:"default_class_id"`
			WorkDayList    []string `json:"work_day_list"`
			ClassesList    []string `json:"classes_list"`
			ManagerList    []string `json:"manager_list"`
			DeptNameList   []string `json:"dept_name_list"`
		} `json:"groups"`
	} `json:"result"`
}

type AttendanceGetsimplegroupsReq struct {
	Offset int64 `json:"offset"`
	Size   int64 `json:"size"`
}

type AttendanceGroupMinimalismListResp struct {
	Result struct {
		HasMore bool `json:"has_more"`
		Cursor  int  `json:"cursor"`
		Result  []struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"result"`
	} `json:"result"`
	Base
}

type AttendanceGroupSearchResp struct {
	Result []struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"result"`
	Base
}

type AttendanceGroupQueryResp struct {
	Result struct {
		Name        string   `json:"name"`
		ID          int      `json:"id"`
		Wifis       []string `json:"wifis"`
		AddressList []string `json:"address_list"`
		WorkDayList []int    `json:"work_day_list"`
		MemberCount int      `json:"member_count"`
		Type        string   `json:"type"`
		URL         string   `json:"url"`
		ManagerList string   `json:"manager_list"`
	} `json:"result"`
	Success bool `json:"success"`
	Base
}

type AttendanceGetusergroupResp struct {
	Base
	Result struct {
		Name    string `json:"name"`
		GroupID int    `json:"group_id"`
		Type    string `json:"type"`
		Classes []struct {
			ClassID  int    `json:"class_id"`
			Name     string `json:"name"`
			Sections []struct {
				Times []struct {
					CheckTime string `json:"check_time"`
					CheckType string `json:"check_type"`
					Across    int    `json:"across"`
					BeginMin  int    `json:"begin_min"`
					EndMin    int    `json:"end_min"`
				} `json:"times"`
			} `json:"sections"`
			Setting struct {
				RestBeginTime struct {
					Across    int    `json:"across"`
					BeginMin  int    `json:"begin_min"`
					EndMin    int    `json:"end_min"`
					CheckTime string `json:"check_time"`
					CheckType string `json:"check_type"`
				} `json:"rest_begin_time"`
				RestEndTime struct {
					Across    int    `json:"across"`
					BeginMin  int    `json:"begin_min"`
					EndMin    int    `json:"end_min"`
					CheckTime string `json:"check_time"`
					CheckType string `json:"check_type"`
				} `json:"rest_end_time"`
			} `json:"setting"`
		} `json:"classes"`
	} `json:"result"`
}

type AttendanceGroupMemberuserListResp struct {
	Base
	Result struct {
		HasMore bool        `json:"has_more"`
		Cursor  interface{} `json:"cursor"`
		Result  []string    `json:"result"`
	} `json:"result"`
}


type AttendanceGroupMemberListbyidsResp struct {
	Base
	Result  []string `json:"result"`
}

type AttendanceGroupMemberUpdateResp struct {
	Base
}

type GroupMemberUpdateParam struct {
	AddDepts []string `json:"add_depts"`
	RemoveDepts []string `json:"remove_depts"`
	AddUsers []string `json:"add_users"`
	RemoveUsers []string `json:"remove_users"`
	AddExtraUsers []string `json:"add_extra_users"`
	RemoveExtraUser []string `json:"remove_extra_user"`
}

type AttendanceGroupMemberUpdateReq struct {
	OpUserID string `json:"op_user_id"`
	ScheduleFlag int64 `json:"schedule_flag"`
	GroupID int64 `json:"group_id"`
	UpdateParam GroupMemberUpdateParam `json:"Object"`
}

type AttendanceGroupMemberListResp struct {
	Base
	Result  struct {
		HasMore bool        `json:"has_more"`
		Cursor  interface{} `json:"cursor"`
		Result  []struct {
			AtcFlag  string `json:"atc_flag"`
			Type     string `json:"type"`
			MemberID string `json:"member_id"`
		} `json:"result"`
	} `json:"result"`
}