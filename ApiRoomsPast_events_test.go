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
	"time"

	"github.com/Chouette2100/exsrapi"
)


const room001 = 75721
const page001 = 8

func TestApiRoomsPast_events(t *testing.T) {

	client, cookiejar, err := exsrapi.CreateNewClient("")
	if err != nil {
		log.Printf("exsrapi.CeateNewClient(): %s", err.Error())
		return //	エラーがあれば、ここで終了
	}
	defer cookiejar.Save()

	type args struct {
		client *http.Client
		roomid int
		page   int
	}
	tests := []struct {
		name    string
		args    args
		wantRpe *RoomsPastevents
		wantErr bool
	}{

		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				client: client,
				roomid: room001,
				page:   1,
			},
		},
		{
			name: "Test2",
			args: args{
				client: client,
				roomid: room001,
				page:   page001,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRpe, err := ApiRoomsPast_events(tt.args.client, tt.args.roomid, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiRoomsPast_events() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRpe, tt.wantRpe) {
				t.Errorf("ApiRoomsPast_events() = %v, want %v", gotRpe, tt.wantRpe)
			}
		})
	}
}

func TestGetRoomsPasteventsByApi(t *testing.T) {

	client, cookiejar, err := exsrapi.CreateNewClient("")
	if err != nil {
		log.Printf("exsrapi.CeateNewClient(): %s", err.Error())
		return //	エラーがあれば、ここで終了
	}
	defer cookiejar.Save()

	type args struct {
		client *http.Client
		roomid int
	}
	tests := []struct {
		name                string
		args                args
		wantRoomspastevents *RoomsPastevents
		wantErr             bool
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				client: client,
				// roomid: 75721,
				// roomid: 87911,
				roomid: 529960,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRoomspastevents, err := GetRoomsPasteventsByApi(tt.args.client, tt.args.roomid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRoomsPasteventsByApi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRoomspastevents, tt.wantRoomspastevents) {
				// t.Errorf("GetRoomsPasteventsByApi() = %v, want %v", gotRoomspastevents, tt.wantRoomspastevents)
				for i, v := range gotRoomspastevents.Events {
					log.Printf("Event[%d]: %s %s\n", i, v.EventName, time.Unix(int64(v.StartedAt), 0).Format("2006-01-02 15:04:05"))	
				}
			}
		})
	}
}
