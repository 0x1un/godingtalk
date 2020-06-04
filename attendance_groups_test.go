package godingtalk

import (
	"fmt"
	"testing"
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
	resp, err := client.OapiAttendanceGroupSearchRequest("2749481918775803", "IT")
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

func TestOapiAttendanceGetusergroupRequest(t *testing.T) {
	resp, err := client.OapiAttendanceGetusergroupRequest("2749481918775803")
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range resp.Result.Classes {
		for _, v2 := range v.Sections {
			fmt.Println(v.Name, v.ClassID, v2.Times[0].CheckTime, v2.Times[1].CheckTime)
		}
	}
}

func TestOapiAttendanceGroupMemberListbyidsRequest(t *testing.T) {
	resp, err := client.OapiAttendanceGroupMemberListbyidsRequest(
		"2749481918775803",
		[]string{"2749481918775803","04553731381000121",
			"6927611729781291","02155716671265668",
			"2617504507699370"}, 0, 358535036)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp.Result)
}

func TestOapiAttendanceGroupMemberUpdateRequest(t *testing.T) {
	// 没有权限。。
}

func TestOapiAttendanceGroupMemberListRequest(t *testing.T) {
	resp, err := client.OapiAttendanceGroupMemberListRequest("2749481918775803", 0, 358535036)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}