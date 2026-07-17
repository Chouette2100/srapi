package srapi_test

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/Chouette2100/srapi/v2"
	"github.com/Chouette2100/srcom"
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

	logfile, err := srcom.CreateLogfile3("ApiOnlive3", srapi.Version)
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

func TestGetLiveOnlives3(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		client   *http.Client
		genreids []int
		want     []srapi.Lives2
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			"TestGetLiveOnlives3",
			http.DefaultClient,
			nil,
			[]srapi.Lives2{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := srapi.GetLiveOnlives3(tt.client, tt.genreids)
			for i, live := range got {
				log.Printf("got[%d]: RoomID=%d, GenreName=%s, StartedAt=%s, MainName=%s",
					i, live.RoomID, live.GenreName,
					time.Unix(live.StartedAt, 0).Format("2006-01-02 15:04:05"), live.MainName)
			}
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetLiveOnlives3() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetLiveOnlives3() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				// t.Errorf("GetLiveOnlives3() = %v, want %v", got, tt.want)
			}
		})
	}
}
