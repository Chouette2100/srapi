package srapi_test

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/Chouette2100/srapi/v2"
	"github.com/Chouette2100/srcom"
)

func TestApiLiveUpcoming(t *testing.T) {
	tests := []struct {
		name    string
		client  *http.Client
		genreID int
		wantErr bool
	}{
		{
			name:    "genre_105",
			client:  http.DefaultClient,
			genreID: 105,
			wantErr: false,
		},
	}

	logfile, err := srcom.CreateLogfile3("ApiLiveUpcoming", srapi.Version)
	if err != nil {
		panic("cannnot open logfile: " + err.Error())
	}
	defer logfile.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := srapi.ApiLiveUpcoming(tt.client, tt.genreID)
			log.Printf("TestApiLiveUpcoming: genre_id=%d", tt.genreID)

			if got != nil {
				log.Printf("upcomings: %d", len(got.Upcomings))
				for i, upcoming := range got.Upcomings {
					startedAt := time.Unix(int64(upcoming.NextLiveStartAt), 0).Format("2006-01-02 15:04:05")
					log.Printf(
						"upcoming[%d]: RoomID=%d, RoomURLKey=%s, MainName=%s, NextLiveStartAt=%s, IsFollow=%t, HasFanRoom=%t, IsFanRoomOnline=%t, OfficialLv=%d, LiverBadge=%d, CellType=%d",
						i,
						upcoming.RoomID,
						upcoming.RoomURLKey,
						upcoming.MainName,
						startedAt,
						upcoming.IsFollow,
						upcoming.HasFanRoom,
						upcoming.IsFanRoomOnline,
						upcoming.OfficialLv,
						upcoming.LiverBadge,
						upcoming.CellType,
					)
					// log.Printf("description=%s", upcoming.Description)
					// log.Printf("theme=%s, image=%s, image_square=%s, new_days_count_label=%s", upcoming.LiverThemeTitle, upcoming.Image, upcoming.ImageSquare, upcoming.NewDaysCountLabel)
				}
			}

			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ApiLiveUpcoming() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ApiLiveUpcoming() succeeded unexpectedly")
			}
		})
	}
}