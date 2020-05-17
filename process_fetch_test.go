// Copyright 2020 aumujun
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package godingtalk

import (
	"fmt"
	"testing"
	"time"
)

func TestOapiProcessinstanceListidsRequest(t *testing.T) {
	loc, _ := time.LoadLocation("Local")
	startTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-02-01 00:00:00", loc)
	if err != nil {
		t.Fatal(err)
	}
	reqData := ProcessinstanceListidsReq{
		ProcessCode: "RPOC-DF144E91-AC5E-4938-A351-22338FF3B7D9",
		UseridList:  "2749481918775803",
		StartTime:   startTime.UnixNano() / 1e6,
		EndTime:     time.Now().UnixNano() / 1e6,
		Cursor:      0,
		Size:        20,
	}
	for {
		resp, err := client.OapiProcessinstanceListidsRequest(reqData)
		if err == nil && resp.Result.NextCursor != 0 {
			reqData.Cursor = resp.Result.NextCursor
			fmt.Println(resp)
			continue
		}
		break
	}
}