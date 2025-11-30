// Copyright © 2025 chouette.21.00@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php

// package srapi_test
package srapi

import (
	"log"

	"net/http"
	"testing"
	// "github.com/Chouette2100/srapi/v2"
)

func TestGetEventRoomsByApi(t *testing.T) {

	client := http.DefaultClient
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		client      *http.Client
		eventUrlKey string
		ib          int
		ie          int
		want        *EventRooms
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			"mattari_fireworks255",
			client,
			"mattari_fireworks255", // Example event URL key
			1,
			500,
			&EventRooms{},
			false,
		}, // Add test case
	}

	// ログファイルの作成
	logfile, err := CreateLogfile(Version, Version)
	if err != nil {
		log.Printf("ログファイルの作成に失敗しました。%v\n", err)
		return
	}
	defer logfile.Close()
	// --------------------------------
	log.Printf("TestGetEventRoomsByApi start\n")
	log.Printf("%+v\n", tests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := GetEventRoomsByApi(tt.client, tt.eventUrlKey, tt.ib, tt.ie)
			for i, room := range got.Rooms {
				log.Printf("Room %d: ID=%d, Name=%s, Point=%d", i+1, room.RoomID, room.RoomName, room.Point)
			}
			if gotErr != nil {
				if !tt.wantErr {
					// t.Errorf("GetEventRoomsByApi() failed: %v", gotErr)
					t.Errorf("GetEventRoomsByApi() failed: ")
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetEventRoomsByApi() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				// t.Errorf("GetEventRoomsByApi() = %v, want %v", got, tt.want)
				t.Errorf("GetEventRoomsByApi() ")
			}
		})
	}
}
