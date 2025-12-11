// Copyright © 2025 chouette.21.00@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php

// package srapi_test
package srapi_test

import (
	"log"
	"net/http"

	"testing"

	"github.com/Chouette2100/srapi/v2"
	"github.com/Chouette2100/srcom"
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
		want        *srapi.EventRooms
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			"enjoykaraoke_vol179",
			client,
			"enjoykaraoke_vol179", // Example event URL key
			1,
			500,
			&srapi.EventRooms{},
			false,
		},
		{
			"mattari_fireworks258",
			client,
			"mattari_fireworks258", // Example event URL key
			1,
			500,
			&srapi.EventRooms{},
			false,
		},
		{
			"heisei_karaoke?block_id=89801",
			client,
			"heisei_karaoke?block_id=89801", // Example event URL key
			1,
			500,
			&srapi.EventRooms{},
			false,
		},
		/*
			{
				"mattari_fireworks255",
				client,
				"mattari_fireworks255", // Example event URL key
				1,
				500,
				&EventRooms{},
				false,
			}, // Add test case
		*/
	}

	// ログファイルの作成
	logfile, err := srcom.CreateLogfile3("ApiEventRooms", srapi.Version)
	if err != nil {
		log.Printf("ログファイルの作成に失敗しました。%v\n", err)
		return
	}
	defer logfile.Close()
	// --------------------------------
	log.Printf("TestGetEventRoomsByApi start\n")

	for j, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := srapi.GetEventRoomsByApi(tt.client, tt.eventUrlKey, tt.ib, tt.ie)
			log.Printf("Test case %d: EventURLKey=%s, ib=%d, ie=%d\n", j+1, tt.eventUrlKey, tt.ib, tt.ie)
			if got != nil {
				log.Printf(" TotalRooms: %d\n", got.TotalEntries)
				for i, room := range got.Rooms {
					log.Printf("Room %d: ID=%d, Name=%s, Point=%d", i+1, room.RoomID, room.RoomName, room.Point)
				}
				log.Printf("Error: %+v\n", got.Errors)
			}
			if gotErr != nil {
				if !tt.wantErr {
					// t.Errorf("GetEventRoomsByApi() failed: %v", gotErr)
					// t.Errorf("GetEventRoomsByApi() failed: ")
					log.Printf("GetEventRoomsByApi() failed: %s", gotErr.Error())
				}
				return
			}
			if tt.wantErr {
				// t.Fatal("GetEventRoomsByApi() succeeded unexpectedly")
				log.Printf("GetEventRoomsByApi() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				// t.Errorf("GetEventRoomsByApi() = %v, want %v", got, tt.want)
				// t.Errorf("GetEventRoomsByApi() ")
				log.Printf("GetEventRoomsByApi() ")
			}
		})
	}
}
