// Copyright © 2024 chouette.21.00@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package srapi

import (
	"log"
	"io"
	"os"

	"github.com/Chouette2100/exsrapi"
	
	"net/http"
	"reflect"
	"testing"
)

func TestApiEventGiftRanking(t *testing.T) {
	type args struct {
		client  *http.Client
		gift_id int
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
		name         string
		args         args
		wantPranking *EventGiftRanking
		wantErr      bool
	}{
		{
			name: "1497",
			args: args{
				client:  client,
				gift_id: 1497,
			},
			wantPranking: nil,
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPranking, err := ApiEventGiftRanking(tt.args.client, tt.args.gift_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiEventGiftRanking() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPranking, tt.wantPranking) {
				//	t.Errorf("ApiEventGiftRanking() = %+v, want %+v", gotPranking, tt.wantPranking)
				//	t.Errorf("ApiEventGiftRanking() error = %+v, wantErr %+v", err, tt.wantErr)
				log.Printf("%s(%d)\n", gotPranking.GiftData[0].Name, gotPranking.GiftData[0].GiftID)
				for _, v := range gotPranking.RankingList {
					log.Printf("%3d: %9s %10d %s\n", v.OrderNo, v.Room.RoomID, v.Score, v.Room.RoomName)
				}
			}
		})
		
	}
}
