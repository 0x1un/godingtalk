package godingtalk

import (
	"testing"
	"fmt"
)

func TestOapiAttendanceGetsimplegroupsRequest(t *testing.T) {
	resp, err := client.OapiAttendanceGetsimplegroupsRequest(0, 10)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}

func TestOapiAttendanceGroupMinimalismListRequest(t *testing.T) {
	resp, err := client.OapiAttendanceGroupMinimalismListRequest("2749481918775803", 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}

func TestOapiAttendanceGroupSearchRequest(t *testing.T) {
	resp, err := client.OapiAttendanceGroupSearchRequest("2749481918775803", "运维")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}

func TestOapiAttendanceGroupQueryRequest(t *testing.T) {
	resp, err := client.OapiAttendanceGroupQueryRequest("2749481918775803", 358535036)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}