/*
!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

Ver. 0.0.0
Ver. 0.1.0 レベルイベントのRankとGapを−１とする。
*/
package srapi

import (
	// "log"

	"net/http"
	"testing"
)

func TestGetPointByApi(t *testing.T) {

	/*
		client, cookiejar, err := CreateNewClient("")
		if err != nil {
			log.Printf("CeateNewClient(): %s", err.Error())
			return //	エラーがあれば、ここで終了
		}
		defer cookiejar.Save()
	*/

	client := &http.Client{}

	type args struct {
		client *http.Client
		roomid int
	}
	tests := []struct {
		name          string
		args          args
		wantPoint     int
		wantRank      int
		wantGap       int
		wantEventid   int
		wantEventurl  string
		wantEventname string
		wantBlockid   int
		wantErr       bool
	}{
		{
			name: "Test_552318",
			args: args{
				client: client,
				roomid: 552318, // みう
			},
			wantPoint:     0,
			wantRank:      -1,
			wantGap:       -1,
			wantEventid:   0,
			wantEventurl:  "",
			wantEventname: "",
			wantBlockid:   0,
			wantErr:       false,
		},
		{
			name: "Test_239199",
			args: args{
				client: client,
				roomid: 239199, // 風花ゆらぎ
			},
			wantPoint:     0,
			wantRank:      -1,
			wantGap:       -1,
			wantEventid:   0,
			wantEventurl:  "",
			wantEventname: "",
			wantBlockid:   0,
			wantErr:       false,
		},
		/*
			{
				name: "Test_96747",
				args: args{
					client: client,
					roomid: 96747, // 無言くん
				},
				wantPoint:     0,
				wantRank:      -1,
				wantGap:       -1,
				wantEventid:   0,
				wantEventurl:  "",
				wantEventname: "",
				wantBlockid:   0,
				wantErr:       false,
			},
			{
				name: "Test_87911",
				args: args{
					client: client,
					roomid: 87911, // Annnnnaの空
				},
				wantPoint:     0,
				wantRank:      -1,
				wantGap:       -1,
				wantEventid:   0,
				wantEventurl:  "",
				wantEventname: "",
				wantBlockid:   0,
				wantErr:       false,
			},
		*/
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPoint, gotRank, gotGap, gotEventid, gotEventurl, gotEventname, gotBlockid, err := GetPointByApi(tt.args.client, tt.args.roomid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPointByApi() error = %v, wantErr %v", err, tt.wantErr)
				//	return
			}
			if gotPoint != tt.wantPoint {
				t.Errorf("GetPointByApi() gotPoint = %v, want %v", gotPoint, tt.wantPoint)
			}
			if gotRank != tt.wantRank {
				t.Errorf("GetPointByApi() gotRank = %v, want %v", gotRank, tt.wantRank)
			}
			if gotGap != tt.wantGap {
				t.Errorf("GetPointByApi() gotGap = %v, want %v", gotGap, tt.wantGap)
			}
			if gotEventid != tt.wantEventid {
				t.Errorf("GetPointByApi() gotEventid = %v, want %v", gotEventid, tt.wantEventid)
			}
			if gotEventurl != tt.wantEventurl {
				t.Errorf("GetPointByApi() gotEventurl = %v, want %v", gotEventurl, tt.wantEventurl)
			}
			if gotEventname != tt.wantEventname {
				t.Errorf("GetPointByApi() gotEventname = %v, want %v", gotEventname, tt.wantEventname)
			}
			if gotBlockid != tt.wantBlockid {
				t.Errorf("GetPointByApi() gotBlockid = %v, want %v", gotBlockid, tt.wantBlockid)
			}
		})
	}
}
