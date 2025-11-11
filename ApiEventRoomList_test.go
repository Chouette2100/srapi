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

	"net/http"
	"reflect"
	"testing"
)

func TestGetRoominfFromEventByApi(t *testing.T) {
	type args struct {
		client  *http.Client
		eventid int
		ib      int
		ie      int
	}
	tests := []struct {
		name            string
		args            args
		wantRoomlistinf *RoomListInf
		wantErr         bool
	}{
		// TODO: Add test cases.
		{
			"TestGetRoominfFromEventByApi",
			args{http.DefaultClient, 37539, 1, 300},
			&RoomListInf{},
			false,
		}, // Add test
		/*
			{
				"TestGetRoominfFromEventByApi",
				args{http.DefaultClient, 37539, 1, 300},
				&RoomListInf{},
				false,
			}, // Add test
			{
				"TestGetRoominfFromEventByApi",
				args{http.DefaultClient, 38384, 1, 300},
				&RoomListInf{},
				false,
			}, // Add test
			/*
			{
				"TestGetRoominfFromEventByApi",
				// args{http.DefaultClient, 37367, 1, 30},
				args{http.DefaultClient, 38256, 1, 30},
				&RoomListInf{},
				false,
			}, // Add test
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRoomlistinf, err := GetRoominfFromEventByApi(tt.args.client, tt.args.eventid, tt.args.ib, tt.args.ie)

			for _, rm := range gotRoomlistinf.RoomList {
				fmt.Printf("%3d %10d %s point=%d\n", rm.Rank, rm.Room_id, rm.Room_name, rm.Point)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("GetRoominfFromEventByApi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRoomlistinf, tt.wantRoomlistinf) {
				t.Errorf("GetRoominfFromEventByApi() = %v, want %v", gotRoomlistinf, tt.wantRoomlistinf)
			}
		})
	}
}
