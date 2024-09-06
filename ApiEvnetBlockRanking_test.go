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
	"io"
	"os"

	"net/http"
	"reflect"
	"testing"

	"github.com/Chouette2100/exsrapi"
)

func TestGetEventBlockRanking(t *testing.T) {
	type args struct {
		client  *http.Client
		eventid int
		blockid int
		ib      int
		ie      int
	}

	logfile, err := exsrapi.CreateLogfile("TestGetEventBlockRanking")
	if err != nil {
		panic("cannnot open logfile: " + err.Error())
	}
	defer logfile.Close()
	//	log.SetOutput(logfile)
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))

	client, cookiejar, err := exsrapi.CreateNewClient("")
	if err != nil {
		log.Printf("exsrapi.CeateNewClient(): %s", err.Error())
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
		/*
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
			//	log.Printf("GetEventBlockRanking(): %+v\n%+v", err, gotEbr)
			log.Printf("eventid= %d, block_id= %d\n", tt.args.eventid, tt.args.eventid)
			for _, br := range(gotEbr.Block_ranking_list) {
			log.Printf("%10s%4d%10d\n", br.Room_id, br.Rank, br.Point)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEventBlockRanking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEbr, tt.wantEbr) {
				t.Errorf("GetEventBlockRanking() = %v, want %v", gotEbr, tt.wantEbr)
			}
		})
	}
}
