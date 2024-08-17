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

	"github.com/Chouette2100/exsrapi"
)

func TestApiRoomStatus(t *testing.T) {
	type args struct {
		client       *http.Client
		room_url_key string
	}
	tests := []struct {
		name           string
		args           args
		wantRoomstatus *RoomStatus
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				client:       nil,
				room_url_key: "mayachacha",
			},
			wantErr: true,
		},
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
			gotRoomstatus, err := ApiRoomStatus(tt.args.client, tt.args.room_url_key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiRoomStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRoomstatus, tt.wantRoomstatus) {
				t.Errorf("ApiRoomStatus() = %v, want %v", gotRoomstatus, tt.wantRoomstatus)
			}
		})
	}
}
