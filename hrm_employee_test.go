package godingtalk

import (
	"fmt"
	"testing"
)

func TestOapiSmartworkHrmEmployeeListRequest(t *testing.T) {
	resp, err := client.OapiSmartworkHrmEmployeeListRequest("310541593537725929")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
