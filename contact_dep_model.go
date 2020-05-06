package godingtalk

import "encoding/json"

// DepartmentCreateReq 部门创建请求
type DepartmentCreateReq struct {
	Name             string `json:"name"`             // 部门名称，长度限制为1~64个字符，不允许包含字符‘-’‘，’以及‘,’
	Parentid         string `json:"parentid"`         // 父部门id，根部门id为1
	Order            string `json:"order"`            // 在父部门中的排序值，order值小的排序靠前
	CreateDeptGroup  bool   `json:"createDeptGroup"`  // 是否创建一个关联此部门的企业群，默认为false
	DeptHiding       bool   `json:"deptHiding"`       // 是否隐藏部门，true表示隐藏 false表示显示
	DeptPermits      string `json:"deptPermits"`      // 可以查看指定隐藏部门的其他部门列表，如果部门隐藏，则此值生效，取值为其他的部门id组成的字符串，使用“\|”符号进行分割。总数不能超过200
	UserPermits      string `json:"userPermits"`      // 可以查看指定隐藏部门的其他人员列表，如果部门隐藏，则此值生效，取值为其他的人员userid组成的字符串，使用“\|”符号进行分割。总数不能超过200
	OuterDept        bool   `json:"outerDept"`        // 限制本部门成员查看通讯录，限制开启后，本部门成员只能看到限定范围内的通讯录。true表示限制开启
	OuterPermitDepts string `json:"outerPermitDepts"` // outerDept为true时，可以配置额外可见部门，值为部门id组成的的字符串，使用“\|”符号进行分割。总数不能超过200
	OuterPermitUsers string `json:"outerPermitUsers"` // outerDept为true时，可以配置额外可见人员，值为userid组成的的字符串，使用“\|”符号进行分割。总数不能超过200
	SourceIdentifier string `json:"sourceIdentifier"` // 部门标识字段，开发者可用该字段来唯一标识一个部门，并与钉钉外部通讯录里的部门做映射
	Ext              string `json:"ext"`              // 部门自定义字段，格式为文本类型的Json格式
}

// DepartmentCreateResp 部门创建响应
type DepartmentCreateResp struct {
	Base
	ID int `json:"id"`
}

// DepartmentDeleteResp 部门删除响应
type DepartmentDeleteResp struct {
	Base
}

// DepartmentListIdsResp 子部门ID列表响应
type DepartmentListIdsResp struct {
	Base
	SubDeptIDList []int `json:"sub_dept_id_list"`
}

// ToBytes 实现Marshallable接口
func (d DepartmentCreateReq) ToBytes() ([]byte, error) {
	return json.Marshal(d)
}
