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

func (d *DingtalkClient) OapiProcessinstanceListidsRequest(reqData ProcessinstanceListidsReq) (ProcessinstanceListidsResp, error) {
	var respData ProcessinstanceListidsResp
	return respData, rpc(d, "topapi/processinstance/listids", d.params, reqData, &respData)
}

func (d *DingtalkClient) OapiProcessinstanceGetRequest(processInstanceID string) (ProcessinstanceGetResp, error) {
	reqData := ProcessinstanceGetReq{
		ProcessInstanceID: processInstanceID,
	}
	var respData ProcessinstanceGetResp
	return respData, rpc(d, "topapi/processinstance/get", d.params, reqData, &respData)
}
