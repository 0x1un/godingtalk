package godingtalk

import "encoding/json"

// AuthScopesResp 权限范围响应
type AuthScopesResp struct {
	Base
	AuthUserField []string `json:"auth_user_field"`
	AuthOrgScopes struct {
		AuthedDept []int    `json:"authed_dept"`
		AuthedUser []string `json:"authed_user"`
	} `json:"auth_org_scopes"`
}

// UserCreateResp 用户创建响应
type UserCreateResp struct {
	Base
	Userid string `json:"userid"`
}

// UserUpdateResp 用户更新响应
type UserUpdateResp struct {
	Base
}

// UserDeleteResp 用户删除响应
type UserDeleteResp struct {
	Base
}

// UserCreateReq 用户创建
type UserCreateReq struct {
	Mobile string `json:"mobile"` // 手机号码，企业内必须唯一，不可重复。如果是国际号码，请使用+xx-xxxxxx的格式
	UserUpdateReq
}

// UserUpdateReq 用户信息更新
type UserUpdateReq struct {
	Userid       string `json:"userid"`       // 员工在当前企业内的唯一标识，也称staffId。可由企业在创建时指定，并代表一定含义比如工号，创建后不可修改，企业内必须唯一。长度为1~64个字符，如果不传，服务器将自动生成一个userid
	Name         string `json:"name"`         // 必要, 用户名称
	OrderInDepts string `json:"orderInDepts"` // 在对应的部门中的排序，Map结构的json字符串，key是部门的Id, value是人员在这个部门的排序值
	Department   []int  `json:"department"`   // 数组类型，数组里面值为整型，成员所属部门id列表
	Position     string `json:"position"`     // 职位信息。长度为0~64个字符
	Tel          string `json:"tel"`          // 必要, 分机号，长度为0~50个字符，企业内必须唯一，不可重复
	WorkPlace    string `json:"workPlace"`    // 办公地点，长度为0~50个字符
	Remark       string `json:"remark"`       // 备注，长度为0~1000个字符
	Email        string `json:"email"`        // 邮箱。长度为0~64个字符。企业内必须唯一，不可重复
	OrgEmail     string `json:"orgEmail"`     // 员工的企业邮箱，员工的企业邮箱已开通，才能增加此字段， 否则会报错
	Jobnumber    string `json:"jobnumber"`    // 员工工号。对应显示到OA后台和客户端个人资料的工号栏目。长度为0~64个字符
	IsHide       bool   `json:"isHide"`       // 是否号码隐藏，true表示隐藏，false表示不隐藏。隐藏手机号后，手机号在个人资料页隐藏，但仍可对其发DING、发起钉钉免费商务电话。
	IsSenior     bool   `json:"isSenior"`     // 是否高管模式，true表示是，false表示不是
	Extattr      struct {
		Interest string `json:"爱好"`
		Age      string `json:"年龄"`
	} `json:"extattr"` // 扩展属性，可以设置多种属性（手机上最多显示10个扩展属性.具体显示哪些属性，请到OA管理后台->设置->通讯录信息设置和OA管理后台->设置->手机端显示信息设置）。该字段的值支持链接类型填写，同时链接支持变量通配符自动替换，目前支持通配符有：userid，corpid。示例： [工位地址](http://www.dingtalk.com?userid=#userid#&corpid=#corpid#)
}

type UserGetResp struct {
	Base
	Unionid         string `json:"unionid"`
	Remark          string `json:"remark"`
	Userid          string `json:"userid"`
	IsLeaderInDepts string `json:"isLeaderInDepts"`
	IsBoss          bool   `json:"isBoss"`
	HiredDate       int64  `json:"hiredDate"`
	IsSenior        bool   `json:"isSenior"`
	Tel             string `json:"tel"`
	Department      []int  `json:"department"`
	WorkPlace       string `json:"workPlace"`
	Email           string `json:"email"`
	OrderInDepts    string `json:"orderInDepts"`
	Mobile          string `json:"mobile"`
	Active          bool   `json:"active"`
	Avatar          string `json:"avatar"`
	IsAdmin         bool   `json:"isAdmin"`
	IsHide          bool   `json:"isHide"`
	Jobnumber       string `json:"jobnumber"`
	Name            string `json:"name"`
	Extattr         struct {
	} `json:"extattr"`
	StateCode string `json:"stateCode"`
	Position  string `json:"position"`
	Roles     []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		GroupName string `json:"groupName"`
	} `json:"roles"`
}

type UserGetDeptMemberResp struct {
	Base
	UserIds []string `json:"userIds"`
}

type UserSimplelistResp struct {
	Base
	HasMore  bool `json:"hasMore"`
	Userlist []struct {
		Userid string `json:"userid"`
		Name   string `json:"name"`
	} `json:"userlist"`
}

type UserListbypageResp struct {
	Base
	HasMore  bool `json:"hasMore"`
	Userlist []struct {
		Userid     string `json:"userid"`
		Unionid    string `json:"unionid"`
		Mobile     string `json:"mobile"`
		Tel        string `json:"tel"`
		WorkPlace  string `json:"workPlace"`
		Remark     string `json:"remark"`
		Order      int    `json:"order"`
		IsAdmin    bool   `json:"isAdmin"`
		IsBoss     bool   `json:"isBoss"`
		IsHide     bool   `json:"isHide"`
		IsLeader   bool   `json:"isLeader"`
		Name       string `json:"name"`
		Active     bool   `json:"active"`
		Department []int  `json:"department"`
		Position   string `json:"position"`
		Email      string `json:"email"`
		Avatar     string `json:"avatar"`
		Jobnumber  string `json:"jobnumber"`
		Extattr    struct {
			Age      string `json:"爱好"`
			Interest string `json:"年龄"`
		} `json:"extattr"`
	} `json:"userlist"`
}

type UserGetAdminResp struct {
	Base
	AdminList []struct {
		SysLevel int    `json:"sys_level"`
		Userid   string `json:"userid"`
	} `json:"admin_list"`
}

type UserGetAdminScopeResp struct {
	Base
	DeptIds []int `json:"dept_ids"`
}

type UserGetUseridByUnionidResp struct {
	Base
	ContactType int    `json:"contactType"`
	Userid      string `json:"userid"`
}

type UserGetByMobileResp struct {
	Base
	Userid string `json:"userid"`
}

type UserGetOrgUserCountResp struct {
	Count int `json:"count"`
	Base
}

type InactiveUserGetResp struct {
	Base
	Result struct {
		HasMore bool     `json:"has_more"`
		List    []string `json:"list"`
	} `json:"result"`
}

func (u UserUpdateReq) ToBytes() ([]byte, error) {
	return json.Marshal(u)
}
