// Copyright 2020 aumujun
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package godingtalk

// OapiProcessinstanceListidsRequest 批量获取审批实例ID
func (d *DingtalkClient) OapiProcessinstanceListidsRequest(reqData ProcessinstanceListidsReq) (ProcessinstanceListidsResp, error) {
	var respData ProcessinstanceListidsResp
	return respData, rpc(d, "topapi/processinstance/listids", d.params, reqData, &respData)
}

// OapiProcessinstanceGetRequest 获取审批实例详情
func (d *DingtalkClient) OapiProcessinstanceGetRequest(processInstanceID string) (ProcessinstanceGetResp, error) {
	reqData := ProcessinstanceGetReq{
		ProcessInstanceID: processInstanceID,
	}
	var respData ProcessinstanceGetResp
	return respData, rpc(d, "topapi/processinstance/get", d.params, reqData, &respData)
}

// OapiProcessGettodonumRequest 获取用户待审批数量
func (d *DingtalkClient) OapiProcessGettodonumRequest(userID string) (ProcessGettodonumResp, error) {
	reqData := ProcessGettodonumReq{
		UserID: userID,
	}
	var respData ProcessGettodonumResp

	return respData, rpc(d, "topapi/process/gettodonum", d.params, reqData, &respData)
}

// OapiProcessListbyuseridRequest 获取用户可见的审批模板
func (d *DingtalkClient) OapiProcessListbyuseridRequest(userID string, offset, size int64) (ProcessListbyuseridResp, error) {

	reqData := ProcessListbyuseridReq{
		UserID: userID,
		Offset: offset,
		Size:   size,
	}
	var respData ProcessListbyuseridResp
	return respData, rpc(d, "topapi/process/listbyuserid", d.params, reqData, &respData)
}

func (d *DingtalkClient) OapiProcessinstanceCspaceInfoRequest(userID string) (ProcessinstanceCspaceInfoResp, error) {
	reqData := ProcessinstanceCspaceInfoReq{
		UserID: userID,
	}
	var respData ProcessinstanceCspaceInfoResp
	return respData, rpc(d, "topapi/processinstance/cspace/info", d.params, reqData, &respData)
}
