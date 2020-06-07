package godingtalk

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
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

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func TestOapiGetUserRequest(t *testing.T) {
	for i := 0; i < 20 ; i++ {

		if i == 4 {
			client.AccessToken.ExpiresTime = 10000000
		}
		time.Sleep(1 * time.Second)
		resp, err := client.OapiUserGetRequest("1519491135941375")
		if err != nil {
			t.Error(err.Error() + "ssssss")
		}
		fmt.Printf("%#v\n", resp.Name)
	}
}

// 105372678
func TestUserGetDeptMemberReques(t *testing.T) {
	for i := 0; i < 20; i++ {
		if i == 4 {
			client.AccessToken.ExpiresTime = client.AccessToken.ExpiresTime - 7200 - 100 - 100
		}
		if i == 16 {
			client.AccessToken.ExpiresTime = client.AccessToken.ExpiresTime - 7200 - 100 - 100
		}
		resp, err := client.OapiUserGetDeptMemberRequest("105372678")
		if err != nil {
			t.Error(err)
		}
		fmt.Println(resp)
	}
}

func TestOapiUserSimplelistRequest(t *testing.T) {
	for i:=0; i < 20 ; i ++ {
		if i == 1 || i == 3 || i == 15 {
			client.AccessToken.ExpiresTime = 0
		}
		resp, err := client.OapiUserSimplelistRequest("105372678", 0, 10, "entry_asc")
		if err != nil {
			t.Error(err)
		}
		fmt.Println(resp)
	}
}

func TestOapiUserListbypageRequest(t *testing.T) {

	resp, err := client.OapiUserListbypageRequest("105372678", 0, 10, "entry_asc")
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
