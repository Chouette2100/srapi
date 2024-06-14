/*
!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

Ver. 0.1.0
*/
package srapi

import (
	"reflect"
	"testing"
)

func TestCrawlRoomRanking(t *testing.T) {
	type args struct {
		mode string
	}
	tests := []struct {
		name    string
		args    args
		wantRr  *[]RoomRanking
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test",
			args{
				"daily",
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRr, err := CrawlRoomRanking(tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("CrawlRoomRanking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRr, tt.wantRr) {
				t.Errorf("CrawlRoomRanking() = %v, want %v", gotRr, tt.wantRr)
			}
		})
	}
}
