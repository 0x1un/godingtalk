package godingtalk

type SmartworkHrmEmployeeListResp struct {
	Result []struct {
		Userid    string `json:"userid"`
		Partner   bool   `json:"partner"`
		FieldList []struct {
			FieldName string `json:"fieldName"`
			FieldCode string `json:"fieldCode"`
			GroupID   string `json:"groupId"`
			Label     string `json:"label,omitempty"`
			Value     string `json:"value,omitempty"`
		} `json:"field_list"`
	} `json:"result"`
	Base
	Success bool `json:"success"`
}

type SmartworkHrmEmployeeListReq struct {
	UserIDList string `json:"userid_list"`
	//FieldFilterList string `json:"field_filter_list"`
}

func (d *DingtalkClient) OapiSmartworkHrmEmployeeListRequest(useridList string) (SmartworkHrmEmployeeListResp, error) {
	reqData := SmartworkHrmEmployeeListReq{
		UserIDList: useridList,
		//FieldFilterList: strings.Join(fieldFilterList, ","),
	}
	var respData SmartworkHrmEmployeeListResp
	return respData, rpc(d, "topapi/smartwork/hrm/employee/list", d.params, reqData, &respData)
}
