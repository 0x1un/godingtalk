package godingtalk

import (
	"fmt"
	"testing"
)

func TestOapiAttendanceShiftListRequest(t *testing.T) {
	resp, err := client.OapiAttendanceShiftListRequest("2749481918775803", 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}

func TestOapiAttendanceShiftQueryRequest(t *testing.T) {
	resp, err := client.OapiAttendanceShiftQueryRequest("2749481918775803", 144280064)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}

func TestOapiAttendanceShiftSearchRequest(t *testing.T) {
	resp, err := client.OapiAttendanceShiftSearchRequest("2749481918775803", "C")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
