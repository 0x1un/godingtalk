package godingtalk

import (
	"fmt"
	"testing"
	"time"
)

func TestOapiAttendanceListscheduleRequest(t *testing.T) {
	ts := time.Now().Format("2006-01-02 15:04:05")
	resp, err := client.OapiAttendanceListscheduleRequest(ts, 0, 200)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}

func TestOapiAttendanceScheduleListbydayRequest(t *testing.T) {
	ts := time.Now().AddDate(0, 0, 1)
	resp, err := client.OapiAttendanceScheduleListbydayRequest("2749481918775803", "1519491135941375", ts.UnixNano()/1e6)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}

func TestOapiAttendanceScheduleListbyusersRequest(t *testing.T) {
	ts := time.Now().AddDate(0, 0, -1)
	resp, err := client.OapiAttendanceScheduleListbyusersRequest("2749481918775803", "2749481918775803,04553731381000121,6927611729781291,02155716671265668,2617504507699370,29616845081220641,322323373831091186,21604948491168487,095931334621426867,15163867091053091,0141304625714090,1519491135941375", ts.UnixNano()/1e6, ts.UnixNano()/1e6)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range resp.Result {
		fmt.Println(v)
	}
}

func TestOapiAttendanceScheduleResultListbyidsRequest(t *testing.T) {

	resp, err := client.OapiAttendanceScheduleResultListbyidsRequest("1519491135941375", "113783017330,113783017331,113783016878,113783017280,113783016979,113783017179,113783017180,113783017130,113783016929,113783017031,113783017080,113783017081,113783017230")
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range resp.Result {
		fmt.Println(v)
	}
}
