package srapi_test

import (
	"log"
	"net/http"

	"testing"

	"github.com/Chouette2100/srapi/v2"
	"github.com/Chouette2100/srcom"
)

func TestApiLiveGiftlist(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		client  *http.Client
		roomid  int
		want    *srapi.LiveGiftlist
		wantErr bool
	}{
		{
			"TestApiLiveGiftlist",
			http.DefaultClient,
			552318, // Example room ID
			&srapi.LiveGiftlist{},
			false,
		},
		// TODO: Add test cases.
	}
	logfile, err := srcom.CreateLogfile3("ApiLiveGiftlist", srapi.Version)
	if err != nil {
		panic("cannnot open logfile: " + err.Error())
	}
	defer logfile.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := srapi.ApiLiveGiftlist(tt.client, tt.roomid)
			for _, gift := range got.Normal {
				log.Printf("Gift ID=%8d,  Point=%8d,  Name=%s,\n", gift.GiftID, gift.Point, gift.GiftName)
			}
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ApiLiveGiftlist() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ApiLiveGiftlist() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("ApiLiveGiftlist() = %v, want %v", got, tt.want)
			}
		})
	}
}
