// Copyright © 2025 chouette.21.00@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php

package srapi

import (
	"log"
	"net/http"

	"reflect"
	"testing"

	"github.com/Chouette2100/srcom"
)

func TestGetEventRankingByApi(t *testing.T) {
	type args struct {
		client      *http.Client
		eventUrlKey string
		ib          int
		ie          int
	}
	tests := []struct {
		name         string
		args         args
		wantPranking *EventRanking
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "enjoykaraoke_vol178",
			args: args{
				client:      &http.Client{},
				eventUrlKey: "enjoykaraoke_vol178",
				// eventUrlKey: "omusubiyokochoo_tokyo",
				ib: 1,
				ie: 500,
			},
			wantPranking: nil,
			wantErr:      false,
		},
	}

	// ログファイルの作成
	logfile, err := srcom.CreateLogfile3(Version, "ApiEventRanking")
	if err != nil {
		log.Printf("ログファイルの作成に失敗しました。%v\n", err)
		return
	}
	defer logfile.Close()
	// --------------------------------

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPranking, err := GetEventRankingByApi(tt.args.client, tt.args.eventUrlKey, tt.args.ib, tt.args.ie)
			log.Printf("EventURLKey: %s", tt.args.eventUrlKey)
			log.Printf("TotalEntries: %d", gotPranking.TotalEntries)
			for i, ranking := range gotPranking.Ranking {
				log.Printf("No. %d Rank: %d, RoomID: %d, Point: %d", i+1, ranking.Rank, ranking.RoomID, ranking.Point)
			}
			log.Printf("End of EventURLKey: %s", tt.args.eventUrlKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEventRankingByApi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPranking, tt.wantPranking) {
				t.Errorf("GetEventRankingByApi() = %v, want %v", gotPranking, tt.wantPranking)
			}
		})
	}
}
