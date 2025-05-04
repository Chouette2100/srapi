// Copyright © 2025 chouette.21.00@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php

package srapi

import (
	"log"

	"net/http"
	"reflect"
	"testing"
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
			name: "TestGetEventRankingByApi",
			args: args{
				client:      &http.Client{},
				eventUrlKey: "ojisan_kobun_kawaii",
				// eventUrlKey: "omusubiyokochoo_tokyo",
				ib: 1,
				ie: 32,
			},
			wantPranking: nil,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPranking, err := GetEventRankingByApi(tt.args.client, tt.args.eventUrlKey, tt.args.ib, tt.args.ie)
			for _, ranking := range gotPranking.Ranking {
				log.Printf("Rank: %d, RoomID: %d, Point: %d", ranking.Rank, ranking.RoomID, ranking.Point)
			}
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
