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

func TestOapiDepartmentListRequest(t *testing.T) {
	resp, err := client.OapiDepartmentListRequest("1", true)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

func TestOapiDepartmentGetRequest(t *testing.T) {

	resp, err := client.OapiDepartmentGetRequest("105372678")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

func TestOapiDepartmentListParentDeptsByDeptRequest(t *testing.T) {
	resp, err := client.OapiDepartmentListParentDeptsByDeptRequest("105372678")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

func TestOapiDepartmentListParentDeptsRequest(t *testing.T) {
	resp, err := client.OapiDepartmentListParentDeptsRequest("2749481918775803")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}
