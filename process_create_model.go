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

type ProcessInstanceApproverVo struct {
	UserIDs        []string `json:"user_ids"`         // 0审批人userid列表。会签/或签列表长度必须大于1，非会签/或签列表长度只能为1
	TaskActionType string   `json:"task_action_type"` //审批类型。AND表示会签，OR表示或签，NONE表示单人
}

type FormComponentValues []FormComponentValue

type FormComponentValue struct {
	Name     interface{} `json:"name"`      // 表单控件每一栏的名称
	Value    interface{} `json:"value"`     // 表单控件每一栏的值
	ExtValue interface{} `json:"ext_value"` // 扩展值
}

type ProcessinstanceCreateReq struct {
	AgentID          string `json:"agent_id"`           // 企业应用标识(ISV调用必须设置)
	ProcessCode      string `json:"process_code"`       // 审批流的唯一码，process_code就在审批流编辑的页面URL中
	OriginatorUserID string `json:"originator_user_id"` // 审批实例发起人的userid
	DeptID           string `json:"dept_id"`            // 发起人所在的部门 如果发起人属于根部门，传-1
	// Approvers           string                      `json:"approvers"`          // 审批人userid列表，最大列表长度20。多个审批人用逗号分隔，按传入的顺序依次审批
	// ApproversV2         []ProcessInstanceApproverVo `json:"approvers_v2"`       // 审批人列表, 支持会签/或签，优先级高于approvers变量
	// CCList              string                      `json:"cc_list"`            // 抄送人userid列表，最大列表长度：20。多个抄送人用逗号分隔。该参数需要与cc_position参数一起传，抄送人才会生效；该参数需要与approvers或approvers_v2参数一起传，抄送人才会生效；
	// CCPosition          string                      `json:"cc_position"`
	FormComponentValues []FormComponentValue `json:"form_component_values"` // 审批流表单控件，最大列表长度20。控件仅支持本文档中列出的表单控件，使用不支持的控件会提示错误, https://ding-doc.dingtalk.com/doc#/serverapi2/cmct1a
}

type ProcessinstanceCreateResp struct {
	Base
	ProcessInstanceID string `json:"process_instance_id"`
}

func (f *FormComponentValues) Add(key, value interface{}) {
	*f = append(*f, FormComponentValue{
		Name:  key,
		Value: value,
	})
}
