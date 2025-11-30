package srapi

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetEventQuestRooms(t *testing.T) {
	type args struct {
		client  *http.Client
		eventid string
		ib      int
		ie      int
	}
	tests := []struct {
		name    string
		args    args
		wantEqr *EventQuestRooms
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "popteen_akb48g_model",
			args: args{
				client:  &http.Client{},
				eventid: "popteen_akb48g_model",
				ib:      1,
				ie:      20,
				// ie:      200,
			},
			wantEqr: nil,
			wantErr: false,
		},
		{
			name: "donuttabetai",
			args: args{
				client:  &http.Client{},
				eventid: "donuttabetai",
				ib:      1,
				ie:      3,
				// ie:      200,
			},
			wantEqr: nil,
			wantErr: false,
		},
		{
			name: "mattari_fireworks255",
			args: args{
				client:  &http.Client{},
				eventid: "mattari_fireworks255",
				ib:      1,
				ie:      3,
				// ie:      200,
			},
			wantEqr: nil,
			wantErr: false,
		},
		{
			name: "weekday_start_00?block_id=75301",
			args: args{
				client:  &http.Client{},
				eventid: "weekday_start_00?block_id=75301",
				ib:      1,
				ie:      3,
				// ie:      200,
			},
			wantEqr: nil,
			wantErr: false,
		},
		{
			name: "giantpanda06",
			args: args{
				client:  &http.Client{},
				eventid: "giantpanda06",
				ib:      1,
				ie:      3,
				// ie:      200,
			},
			wantEqr: nil,
			wantErr: false,
		},
		{
			name: "listenerupupup_showroom260",
			args: args{
				client:  &http.Client{},
				eventid: "listenerupupup_showroom260",
				ib:      1,
				ie:      3,
				// ie:      200,
			},
			wantEqr: nil,
			wantErr: false,
		},
		{
			name: "hanakin_happy_night_007?block_id=85901",
			args: args{
				client:  &http.Client{},
				eventid: "hanakin_happy_night_007?block_id=85901",
				ib:      1,
				ie:      3,
				// ie:      200,
			},
			wantEqr: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEqr, err := GetEventQuestRoomsByApi(tt.args.client, tt.args.eventid, tt.args.ib, tt.args.ie)
			if err != nil {
				fmt.Printf("GetEventQuestRooms() error: %s\n", err.Error())
				return
			}
			if gotEqr == nil {
				err = fmt.Errorf("GetEventQuestRooms() gotEqr is nil")
				t.Logf("%s", err.Error())
				return
			}
			for i, rm := range gotEqr.EventQuestLevelRanges[0].Rooms {
				t.Logf("No.%4d RoomID: %7d  point: %10d  questlevel=%4d", i, rm.RoomID, rm.Point, rm.QuestLevel)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEventQuestRooms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEqr, tt.wantEqr) {
				// t.Errorf("GetEventQuestRooms() = %v, want %v", gotEqr, tt.wantEqr)
				t.Errorf("GetEventQuestRooms()  got != want ")
			}
		})
	}
}
