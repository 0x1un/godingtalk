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
