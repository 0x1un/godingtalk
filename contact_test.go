package godingtalk

import (
	"fmt"
	"os"
	"testing"
)

func TestOapiAuthScopesRequest(t *testing.T) {
	client := NewDingtalkClient(os.Getenv("APPKEY"), os.Getenv("APPSECRET"))
	if err := client.Init(); err != nil {
		t.Error(err)
	}
	resp, err := client.OapiAuthScopesRequest()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}
