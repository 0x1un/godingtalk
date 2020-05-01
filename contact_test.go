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
	resp, err := client.OapiUserGetRequest("2749481918775803")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

func TestGetTagName(t *testing.T) {
}
