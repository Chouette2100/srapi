/*
!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

Ver. 0.1.0
*/
package srapi

import (
	"log"

	"net/http"
	"reflect"
	"testing"

	"github.com/chouette2100/exsrapi"
)

func TestApiEventsRanking(t *testing.T) {
	type args struct {
		client   *http.Client
		ieventid int
		roomid   int
		blockid  int
	}
	tests := []struct {
		name         string
		args         args
		wantPranking *Eventranking
		wantErr      bool
	}{
		// TODO: Add test cases.
		/*
		{
			name: "test4",
			args: args{
				client:   nil,
				ieventid: 36450,
				roomid:   507716,
				blockid:  0,
			},
			wantErr: true,
		},
		*/
		{
			name: "test3",
			args: args{
				client:   nil,
				ieventid: 36142,
				roomid:   192641,
				blockid:  0,
			},
			wantErr: true,
		},
		/*
		{
			name: "test1",
			args: args{
				client:   nil,
				ieventid: 35221,  //	【新規枠・2nd Stage】SR限定『ホークス応援モデルオーディション～5/19～』
				roomid:   501854, //	なおのへや
				blockid:  0,
			},
			wantErr: true,
		},
		{
			name: "test2",
			args: args{
				client:   nil,
				ieventid: 35074,  //	iitoJAPAN スタートダッシュイベント vol.18
				roomid:   500695, //	椿ミヤ
				blockid:  18101,  //	[SS-5~C-1]Aブロック
			},
			wantErr: true,
		},
		*/
	}

	client, cookiejar, err := exsrapi.CreateNewClient("")
	if err != nil {
		log.Printf("exsrapi.CeateNewClient(): %s", err.Error())
		return //	エラーがあれば、ここで終了
	}
	defer cookiejar.Save()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.client = client
			gotPranking, err := ApiEventsRanking(tt.args.client, tt.args.ieventid, tt.args.roomid, tt.args.blockid)
			for i, v := range gotPranking.Ranking {
				log.Printf("Ranking[%2d]: %7d, %3d, %8d, %s\n", i, v.Room.RoomID, v.Rank, v.Point, v.Room.Name)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiRoomStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPranking, tt.wantPranking) {
				t.Errorf("ApiEventsRanking() = %v, want %v", gotPranking, tt.wantPranking)
			}
		})
	}
}
