package godingtalk

import (
	"fmt"
	"testing"
)

func TestOapiRoleListRequest(t *testing.T) {
	resp, err := client.OapiRoleListRequest(RoleListRequest{Size: 20, Offset: 0})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}
