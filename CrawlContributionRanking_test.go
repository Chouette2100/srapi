package srapi

import (
	"testing"
	// "github.com/Chouette2100/srapi/v2"
)

func TestCrawlContirbutionRanking(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		EventName string
		roomno    int
		want      int
		want2     CrContributionRanking
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name:      "Test Case 1",
			EventName: "tpidol_shizuoka_sf",
			roomno:    545286,
			want:      100,
			want2:     CrContributionRanking{TotalScore: 100, BonusPoints: 0, EventRanking: CrEventRanking{{Rank: 1, Point: 100, Listner: "User1"}}},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := CrawlContirbutionRanking(tt.EventName, tt.roomno)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CrawlContirbutionRanking() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CrawlContirbutionRanking() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("CrawlContirbutionRanking() = %v, want %v", got, tt.want)
			}
		})
	}
}
