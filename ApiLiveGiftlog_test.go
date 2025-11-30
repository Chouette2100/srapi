package srapi_test

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/Chouette2100/srapi/v2"
)

func TestApiLiveGiftlog(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		client  *http.Client
		roomid  int
		want    *srapi.GiftLoglist
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"TestApiLiveGiftlog",
			http.DefaultClient,
			75721, // Example room ID
			&srapi.GiftLoglist{},
			false,
		},
	}

	logfile, err := srapi.CreateLogfile("ApiLiveGiftlog", time.Now().Format("150405"), srapi.Version)
	if err != nil {
		panic("cannnot open logfile: " + err.Error())
	}
	defer logfile.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := srapi.ApiLiveGiftlog(tt.client, tt.roomid)
			for _, giftlog := range got.GiftLog {
				unixtime := int64(giftlog.CreatedAt)
				t := time.Unix(unixtime, 0).Format("15:04:05")
				log.Printf("GiftLog Time=%s, GiftID=%8d, Num=%3d, GiftID=%d,\n",
					t, giftlog.GiftID, giftlog.Num, giftlog.GiftID)
			}
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ApiLiveGiftlog() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ApiLiveGiftlog() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("ApiLiveGiftlog() = %v, want %v", got, tt.want)
			}
		})
	}
}
