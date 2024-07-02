package srapi

import (
	"log"
	"time"
	"reflect"
	"testing"

	"net/http"

	"github.com/chouette2100/exsrapi"
)

func TestApiRoomProfile(t *testing.T) {

	client, cookiejar, err := exsrapi.CreateNewClient("")
	if err != nil {
		log.Printf("exsrapi.CeateNewClient(): %s", err.Error())
		return //	エラーがあれば、ここで終了
	}
	defer cookiejar.Save()

	type args struct {
		client  *http.Client
		room_id string
	}
	tests := []struct {
		name        string
		args        args
		wantRoominf RoomInfAll
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				client:  client,
				room_id: "239199",
			},
			wantRoominf: RoomInfAll{},
			wantErr:     false,
		},
		{
			name: "test2",
			args: args{
				client:  client,
				room_id: "999999",
			},
			wantRoominf: RoomInfAll{},
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRoominf, err := ApiRoomProfile(tt.args.client, tt.args.room_id)

			startedat := time.Unix(gotRoominf.Event.StartedAt,0).Format("2006-01-02 15:04:05")
			log.Printf(" StartedAt = %s", startedat) 
			endedat := time.Unix(gotRoominf.Event.EndedAt, 0).Format("2006-01-02 15:04:05")
			log.Printf(" birthday = %s", endedat)
			// birthday := time.Unix(gotRoominf.Birthday, 0).Format("01-02")
			log.Printf(" birthday = %s", time.Unix(gotRoominf.Birthday, 0).Format("(2006-)01-02 (15:04:05)"))

			if (err != nil) != tt.wantErr {
				t.Errorf("ApiRoomProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRoominf, tt.wantRoominf) {
				t.Errorf("ApiRoomProfile() = %v, want %v", gotRoominf, tt.wantRoominf)
			}
		})
	}
}
