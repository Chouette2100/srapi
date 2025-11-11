package srapi

import (
	"net/http"
	"testing"
	// "github.com/Chouette2100/srapi/v2"
)

func TestApiLiveOnlives(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		client  *http.Client
		want    *RoomOnlives
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"TestApiLiveOnlives",
			http.DefaultClient,
			&RoomOnlives{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := ApiLiveOnlives(tt.client)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ApiLiveOnlives() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ApiLiveOnlives() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("ApiLiveOnlives() = %v, want %v", got, tt.want)
			}
		})
	}
}
