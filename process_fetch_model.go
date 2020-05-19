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

type ProcessinstanceGetReq struct {
	ProcessInstanceID string `json:"process_instance_id"`
}

type ProcessinstanceGetResp struct {
	Base
	ProcessInstance struct {
		Title               string   `json:"title"`
		CreateTime          string   `json:"create_time"`
		FinishTime          string   `json:"finish_time"`
		OriginatorUserid    string   `json:"originator_userid"`
		OriginatorDeptID    string   `json:"originator_dept_id"`
		Status              string   `json:"status"`
		CcUserids           []string `json:"cc_userids"`
		FormComponentValues []struct {
			Name     string `json:"name"`
			Value    string `json:"value"`
			ExtValue string `json:"ext_value"`
		} `json:"form_component_values"`
		Result           string `json:"result"`
		BusinessID       string `json:"business_id"`
		OperationRecords []struct {
			Userid          string `json:"userid"`
			Date            string `json:"date"`
			OperationType   string `json:"operation_type"`
			OperationResult string `json:"operation_result"`
			Remark          string `json:"remark"`
		} `json:"operation_records"`
		Tasks []struct {
			Userid     string `json:"userid"`
			TaskStatus string `json:"task_status"`
			TaskResult string `json:"task_result"`
			CreateTime string `json:"create_time"`
			FinishTime string `json:"finish_time"`
			Taskid     string `json:"taskid"`
		} `json:"tasks"`
		OriginatorDeptName         string   `json:"originator_dept_name"`
		BizAction                  string   `json:"biz_action"`
		AttachedProcessInstanceIds []string `json:"attached_process_instance_ids"`
	} `json:"process_instance"`
}

type ProcessGettodonumResp struct {
	Count int `json:"count"`
	Base
}

type ProcessGettodonumReq struct {
	UserID string `json:"userid"`
}

type ProcessListbyuseridReq struct {
	UserID string `json:"userid"`
	Offset int64  `json:"offset"`
	Size   int64  `json:"size"`
}

type ProcessListbyuseridResp struct {
	Base
	Result struct {
		ProcessList []struct {
			Name        string `json:"name"`
			IconURL     string `json:"icon_url"`
			ProcessCode string `json:"process_code"`
			URL         string `json:"url"`
		} `json:"process_list"`
		NextCursor int `json:"next_cursor"`
	} `json:"result"`
}

type ProcessinstanceCspaceInfoReq struct {
	UserID string `json:"user_id"`
}

type ProcessinstanceCspaceInfoResp struct {
	Base
	Result struct {
		// 审批钉盘空间id
		SpaceID int `json:"space_id"`
	} `json:"result"`
	Success bool `json:"success"`
}
