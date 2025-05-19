package srapi

import (
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
			name: "TestGetEventQuestRooms",
			args: args{
				client:  &http.Client{},
				eventid: "mattari_fireworks229",
				ib:      1,
				ie:      1,
				// ie:      200,
			},
			wantEqr: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEqr, err := GetEventQuestRooms(tt.args.client, tt.args.eventid, tt.args.ib, tt.args.ie)
			for i, rm := range gotEqr.EventQuestLevelRanges[0].Rooms {
				t.Logf("No.%4d RoomID: %d", i, rm.RoomID)
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
