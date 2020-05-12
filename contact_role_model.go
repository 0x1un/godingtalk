package godingtalk

type RoleListResp struct {
	Base
	Result struct {
		HasMore bool `json:"hasMore"`
		List    []struct {
			Name    string `json:"name"`
			GroupID int    `json:"groupId"`
			Roles   []struct {
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"roles"`
		} `json:"list"`
	} `json:"result"`
}

type RoleSimplelistResp struct {
	Base
	Result struct {
		HasMore bool `json:"hasMore"`
		List    []struct {
			Userid       string `json:"userid"`
			Name         string `json:"name"`
			ManageScopes []struct {
				DeptID int    `json:"dept_id"`
				Name   string `json:"name"`
			} `json:"manageScopes"`
		} `json:"list"`
	} `json:"result"`
}
