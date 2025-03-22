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

func TestApiEventAbstraction(t *testing.T) {
	type args struct {
		client    *http.Client
		eventid   string
	}
	tests := []struct {
		name    string
		args    args
		wantEa  *EventAbstraction
		wantErr bool
	}{
		{
			name: "sr_tsutsuzyuku_geinin_2",
			args: args{
				client:    &http.Client{},
				eventid:   "sr_tsutsuzyuku_geinin_2",
			},
			wantEa: &EventAbstraction{},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEa, err := ApiEventAbstraction(tt.args.client, tt.args.eventid)
			log.Printf("ApiEventAbstraction() = %v, err = %+v", gotEa, err)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiEventAbstraction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEa, tt.wantEa) {
				t.Errorf("ApiEventAbstraction() = %v, want %v", gotEa, tt.wantEa)
			}
		})
	}
}
