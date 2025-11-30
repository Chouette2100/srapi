package srapi_test

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/Chouette2100/srapi/v2"
)

func TestApiLiveOnlives3(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		client  *http.Client
		want    *srapi.LiveOnlive
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"TestApiLiveOnlives2",
			http.DefaultClient,
			&srapi.LiveOnlive{},
			false,
		},
	}

	logfile, err := srapi.CreateLogfile("ApiOnlive2", time.Now().Format("150405"), srapi.Version)
	if err != nil {
		panic("cannnot open logfile: " + err.Error())
	}
	defer logfile.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := srapi.ApiLiveOnlives3(tt.client)
			log.Printf("Bcsvr_post: %d", got.BcsvrPort)
			log.Printf("Bcsvr_Host: %s", got.BcsvrHost)
			for _, onlive := range got.Onlives {
				log.Printf("Genre_id: %d", onlive.GenreID)
				log.Printf("Genre_name: %s", onlive.GenreName)
				for _, live := range onlive.Lives {
					log.Printf(" RoomID: %d, MainName: %s, Bcsvrkey: %s", live.RoomID, live.MainName, live.BcsvrKey)
				}
			}
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ApiLiveOnlives2() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ApiLiveOnlives2() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("ApiLiveOnlives2() = %v, want %v", got, tt.want)
			}
		})
	}
}
