package godingtalk

import (
	"os"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	client := NewDingtalkClient(os.Getenv("APPKEY"), os.Getenv("APPSECRET"))
	err := client.setToken()
	if err != nil {
		t.Error(err)
	}
	err = storeToken(client.AccessToken)
	t.Error(err)
}
