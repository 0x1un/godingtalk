package godingtalk

import (
	"fmt"
	"testing"
)

func TestOapiAttendanceListRecordRequest(t *testing.T) {
	resp, err := client.OapiAttendanceListRecordRequest([]string{"2749481918775803"}, "2020-05-15 00:00:00", "2020-05-22 00:00:00", false)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}

func TestOapiAttendanceListRequest(t *testing.T) {
	resp, err := client.OapiAttendanceListRequest([]string{"2749481918775803"}, "2020-05-22 00:00:00", "2020-05-22 00:00:00", 0, 50)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
