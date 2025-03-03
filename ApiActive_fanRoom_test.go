package srapi

import (
	"log"

	"net/http"
	"reflect"
	"testing"
)

func TestApiActivefanRoom(t *testing.T) {

	client, cookiejar, err := CreateNewClient("")
	if err != nil {
		log.Printf("CeateNewClient(): %s", err.Error())
		return //	エラーがあれば、ここで終了
	}
	defer cookiejar.Save()

	type args struct {
		client  *http.Client
		room_id string
		ym      string
	}
	tests := []struct {
		name     string
		args     args
		wantPafr *ActivefanRoom
		wantErr  bool
	}{
		{
			name: "test1",
			args: args{
				client:  client,
				room_id: "87911",
				ym:      "202404",
			},
			wantPafr: &ActivefanRoom{},
			wantErr:  false,
		},
		{
			name: "test2",
			args: args{
				client:  client,
				room_id: "87911",
				ym:      "202403",
			},
			wantPafr: &ActivefanRoom{},
			wantErr:  false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPafr, err := ApiActivefanRoom(tt.args.client, tt.args.room_id, tt.args.ym)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiActivefanRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPafr, tt.wantPafr) {
				t.Errorf("ApiActivefanRoom() = %v, want %v", gotPafr, tt.wantPafr)
			}
		})
	}
}
