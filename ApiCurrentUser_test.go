// Copyright © 2024 chouette.21.00@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package srapi

import (
	"net/http"
	"reflect"
	"testing"
)

func TestApiCurrentUser(t *testing.T) {
	type args struct {
		client *http.Client
	}
	tests := []struct {
		name    string
		args    args
		wantPcu *CurrentUser
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestApiCurrentUser",
			args: args{
				client: &http.Client{},
			},
			wantPcu: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPcu, err := ApiCurrentUser(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiCurrentUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPcu, tt.wantPcu) {
				t.Errorf("ApiCurrentUser() = %v, want %v", gotPcu, tt.wantPcu)
			}
		})
	}
}
