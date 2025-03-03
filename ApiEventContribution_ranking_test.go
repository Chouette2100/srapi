/*
!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

Ver. 1.1.1	Contribution_ranking.Me のtypeをanyからinterface{}に変更する。テスト用データを変更する。

*/
package srapi

import (
	"log"

	"net/http"
	"reflect"
	"testing"
)

func TestApiContribution_ranking(t *testing.T) {
	type args struct {
		client   *http.Client
		ieventid int
		roomid   int
	}
	tests := []struct {
		name         string
		args         args
		wantPranking *Eventranking
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "test3",
			args: args{
				client:   nil,
				ieventid: 36159,  //	
				roomid:   518083, //	
			},
			wantErr: true,
		},
		/*
		{
			name: "test1",
			args: args{
				client:   nil,
				ieventid: 35315,  //	
				roomid:   417115, //	
			},
			wantErr: true,
		},
		{
			name: "test2",
			args: args{
				client:   nil,
				ieventid: 35183,  //	まったり配信したいあなたへ♡みんなで花火を楽しもう！vol.166 
				roomid:   408389, //	日向端ひな（高嶺のなでしこ）
			},
			wantErr: true,
		},
		*/
	}

	client, cookiejar, err := CreateNewClient("")
	if err != nil {
		log.Printf("CeateNewClient(): %s", err.Error())
		return //	エラーがあれば、ここで終了
	}
	defer cookiejar.Save()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.client = client
			gotPranking, err := ApiEventContribution_ranking(tt.args.client, tt.args.ieventid, tt.args.roomid)
			for i, v := range gotPranking.Ranking {
				log.Printf("Ranking[%2d]: %3d %8d, %8d, %s\n", i, v.Rank, v.UserID, v.Point, v.Name)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiEventContribution_ranking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPranking, tt.wantPranking) {
				t.Errorf("ApiEventContribution_rankingApi() = %v, want %v", gotPranking, tt.wantPranking)
			}
		})
	}
}
