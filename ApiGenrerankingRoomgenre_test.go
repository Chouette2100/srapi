// Copyright © 2024 chouette.21.00@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package srapi

import (
	"net/http"
	"reflect"
	"testing"
)

func TestApiGenrerankingRoomGenre(t *testing.T) {
	type args struct {
		client *http.Client
	}
	tests := []struct {
		name       string
		args       args
		wantPgenre *RoomGenreList
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "TestApiGenrerankingRoomGenre",
			args: args{
				client: &http.Client{},
			},
			wantPgenre: nil,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPgenre, err := ApiGenrerankingRoomGenre(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiGenrerankingRoomGenre() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPgenre, tt.wantPgenre) {
				t.Errorf("ApiGenrerankingRoomGenre() = %v, want %v", gotPgenre, tt.wantPgenre)
			}
		})
	}
}
