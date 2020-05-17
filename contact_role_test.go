package godingtalk

import (
	"fmt"
	"testing"
)

func TestOapiRoleListRequest(t *testing.T) {
	resp, err := client.OapiRoleListRequest(20, 0)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}
