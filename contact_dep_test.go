package godingtalk

import (
	"fmt"
	"testing"
)

// func TestOapiDepartmentCreateRequest(t *testing.T) {
// 	resp, err := client.OapiDepartmentCreateRequest()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Println(resp)
// }

func TestOapiDepartmentDeleteRequest(t *testing.T) {
	resp, err := client.OapiDepartmentDeleteRequest("xxx")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

func TestOapiDepartmentListIdsRequest(t *testing.T) {
	resp, err := client.OapiDepartmentListIdsRequest("105372678")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}
