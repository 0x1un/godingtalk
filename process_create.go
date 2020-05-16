package godingtalk

func (d *DingtalkClient) OapiProcessinstanceCreateRequest(reqData ProcessinstanceCreateReq) (ProcessinstanceCreateResp, error) {
	var respData ProcessinstanceCreateResp
	return respData, rpc(d, "topapi/processinstance/create", d.params, reqData, &respData)
}
