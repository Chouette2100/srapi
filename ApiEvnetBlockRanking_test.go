/*
!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

Ver. 0.1.0
*/
package srapi

import (
	"fmt"
	"io"
	"log"
	"os"

	"net/http"
	// "reflect"
	"testing"
	// "golang.org/x/tools/go/analysis/passes/defers"
)

func TestGetEventBlockRanking(t *testing.T) {
	type args struct {
		client  *http.Client
		eventid int
		blockid int
		ib      int
		ie      int
	}

	logfile, err := CreateLogfile("TestGetEventBlockRanking")
	if err != nil {
		panic("cannnot open logfile: " + err.Error())
	}
	defer logfile.Close()
	//	log.SetOutput(logfile)
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))

	client, cookiejar, err := CreateNewClient("")
	if err != nil {
		log.Printf("CeateNewClient(): %s", err.Error())
		return //	エラーがあれば、ここで終了
	}
	defer cookiejar.Save()

	tests := []struct {
		name    string
		args    args
		wantEbr *EventBlockRanking
		wantErr bool
	}{
		// TODO: Add test cases.
		// https://showroom-live.com/event/ojisan_kobun_kawaii?block_id=0
		{
			name: "weekday_start_006",
			args: args{
				client:  client,
				eventid: 40735,
				blockid: 75301,
				ib:      1,
				ie:      3,
			},
			wantEbr: nil,
			wantErr: false,
		},
		{
			name: "hanakin_happy_night_007",
			args: args{
				client:  client,
				eventid: 41011,
				blockid: 85901,
				ib:      1,
				ie:      3,
			},
			wantEbr: nil,
			wantErr: false,
		},
		/*
			{
				name: "little_love_valentine?block_id=30801",
				args: args{
					client:  client,
					eventid: 38246,
					blockid: 30801,
					ib:      1,
					ie:      100,
				},
				wantEbr: nil,
				wantErr: false,
			},
			{
				name: "wdebutf_s1?block_id=25501",
				args: args{
					client:  client,
					eventid: 37430,
					blockid: 25501,
					ib:      1,
					ie:      100,
				},
				wantEbr: nil,
				wantErr: false,
			},
			{
				name: "TestGetEventBlockRanking",
				args: args{
					client:  client,
					eventid: 36695,
					blockid: 0,
					ib:      1,
					ie:      100,
				},
				wantEbr: nil,
				wantErr: false,
			},
			{
				name: "TestGetEventBlockRanking",
				args: args{
					client:  client,
					eventid: 36695,
					blockid: 0,
					ib:      1,
					ie:      20,
				},
				wantEbr: nil,
				wantErr: false,
			},
					{
				name: "TestGetEventBlockRanking",
				args: args{
					client:  client,
					eventid: 36695,
					blockid: 20901,
					ib:      1,
					ie:      200,
				},
				wantEbr: nil,
				wantErr: false,
			},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEbr, err := GetEventBlockRanking(tt.args.client, tt.args.eventid, tt.args.blockid, tt.args.ib, tt.args.ie)
			if err != nil {
				log.Printf("GetEventBlockRanking(): %s", err.Error())
				return
			}
			//	log.Printf("GetEventBlockRanking(): %+v\n%+v", err, gotEbr)
			log.Printf("eventid= %d, block_id= %d\n", tt.args.eventid, tt.args.blockid)
			f, err := os.OpenFile("TestGetEventBlockRanking.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
			if err != nil {
				log.Printf("cannot open TestGetEventBlockRanking.txt: %s", err.Error())
				return
			}
			defer f.Close()
			lng := len(gotEbr.Block_ranking_list)
			// for _, br := range(gotEbr.Block_ranking_list) {
			for i := lng - 1; i >= 0; i-- {
				br := gotEbr.Block_ranking_list[i]
				// log.Printf("%10s%4d%10d\n", br.Room_id, br.Rank, br.Point)
				fmt.Fprintf(f, "%s\n", br.Room_id)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEventBlockRanking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for i, br := range gotEbr.Block_ranking_list {
				log.Printf("No.%4d Room_id: %s, Rank: %d, Point: %d", i+1, br.Room_id, br.Rank, br.Point)
			}
			// if !reflect.DeepEqual(gotEbr, tt.wantEbr) {
			// 	t.Errorf("GetEventBlockRanking() = %v, want %v", gotEbr, tt.wantEbr)
			// }
		})
	}
}
