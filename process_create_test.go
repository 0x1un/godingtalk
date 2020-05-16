package godingtalk

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestOapiProcessinstanceCreateRequest(t *testing.T) {
	var fvs FormComponentValues
	fvs.Add("城市", "高朋")
	fvs.Add("台席号", "343")
	fvs.Add("域账号", "ssss")
	fvs.Add("联系方式", "13800138000")
	fvs.Add("故障类型", "业务系统故障")
	fvs.Add("故障范围", "单个台席")
	fvs.Add("故障现象", "xdjfjdk")
	agentID, _ := strconv.Atoi(os.Getenv("APP_AGENT_ID"))
	depID, _ := strconv.Atoi(os.Getenv("DEPID"))
	reqData := ProcessinstanceCreateReq{
		AgentID:             agentID,
		ProcessCode:         "RPOC-DF144E91-AC5E-4938-A351-22338FF3B7D9",
		OriginatorUserID:    os.Getenv("USERID"),
		DeptID:              depID,
		FormComponentValues: fvs,
	}
	resp, err := client.OapiProcessinstanceCreateRequest(reqData)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}
