package godingtalk

import (
	"fmt"
	"os"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	client := NewDingtalkClient(os.Getenv("APPKEY"), os.Getenv("APPSECRET"))
	err := client.setToken()
	if err != nil {
		t.Error(err)
	}
	tr, err := readToken()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(tr)
}
