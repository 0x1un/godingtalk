package godingtalk

import (
	"fmt"
	"os"
	"testing"
)

var (
	client = NewDingtalkClient(os.Getenv("APPKEY"), os.Getenv("APPSECRET"))
)

func init() {
	if err := client.Init(); err != nil {
		panic(err)
	}
}

func TestOapiAuthScopesRequest(t *testing.T) {
	resp, err := client.OapiAuthScopesRequest()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

func TestOapiGetUserRequest(t *testing.T) {
	resp, err := client.OapiUserGetRequest("1519491135941375")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v", resp.Name)
}

// 105372678
func TestUserGetDeptMemberReques(t *testing.T) {
	resp, err := client.OapiUserGetDeptMemberRequest("105372678")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

func TestOapiUserSimplelistRequest(t *testing.T) {
	resp, err := client.OapiUserSimplelistRequest("105372678", "0", "10", "entry_asc")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

func TestOapiUserListbypageRequest(t *testing.T) {

	resp, err := client.OapiUserListbypageRequest("105372678", "0", "10", "entry_asc")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}
func TestOapiUserGetAdminRequest(t *testing.T) {
	resp, err := client.OapiUserGetAdminRequest()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

func TestOapiUserGetAdminScopeRequest(t *testing.T) {
	resp, err := client.OapiUserGetAdminScopeRequest("2749481918775803")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

func TestOapiUserGetUseridByUnionidReques(t *testing.T) {
	resp, err := client.OapiUserGetUseridByUnionidRequest("xxxx")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

func TestOapiUserGetByMobileRequest(t *testing.T) {
	resp, err := client.OapiUserGetByMobileRequest("17608035126")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

func TestOapiUserGetOrgUserCountRequest(t *testing.T) {
	resp, err := client.OapiUserGetOrgUserCountRequest("1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

func TestOapiInactiveUserGetRequest(t *testing.T) {
	resp, err := client.OapiInactiveUserGetRequest("20200503", "0", "100")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}
