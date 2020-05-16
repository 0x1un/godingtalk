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

import "encoding/json"

type ProcessinstanceListidsReq struct {
	Cursor      int    `json:"cursor"`       // 分页查询的游标，最开始传0，后续传返回参数中的next_cursor值，默认值：0
	StartTime   int64  `json:"start_time"`   // * 开始时间。Unix时间戳，单位毫秒。
	EndTime     int64  `json:"end_time"`     // 结束时间。不传该参数则默认取当前时间。Unix时间戳，单位毫秒。
	Size        int64  `json:"size"`         // 分页参数，每页大小，最多传20，默认值：20
	UseridList  string `json:"userid_list"`  // 发起人用户id列表，用逗号分隔，最大列表长度：10
	ProcessCode string `json:"process_code"` // * 流程模板唯一标识，可在OA管理后台编辑审批表单部分查询
}

type ProcessinstanceListidsResp struct {
	Base
	Result struct {
		List       []string `json:"list"`
		NextCursor int      `json:"next_cursor"`
	} `json:"result"`
}

func (p ProcessinstanceListidsReq) ToBytes() ([]byte, error) {
	return json.Marshal(p)
}
