package godingtalk

import (
	"fmt"
	"os"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	client := NewDingtalkClient(os.Getenv("APPKEY"), os.Getenv("APPSECRET"))
	if err := client.Init(); err != nil {
		t.Error(err)
	}
	fmt.Println(client.params)
}
