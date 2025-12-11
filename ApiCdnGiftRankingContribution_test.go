// Copyright © 2024 chouette.21.00@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package srapi_test

import (
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/Chouette2100/srapi/v2"
	"github.com/Chouette2100/srcom"
)

func TestApiCdnGiftRankingContribution(t *testing.T) {
	type args struct {
		client   *http.Client
		genre_id int
		url_key  string
	}

	logfile, err := srcom.CreateLogfile3("TestApiCdnGiftRankingContribution", srapi.Version)
	if err != nil {
		panic("cannnot open logfile: " + err.Error())
	}
	defer logfile.Close()
	//	log.SetOutput(logfile)
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))

	client, cookiejar, err := srapi.CreateNewClient("")
	if err != nil {
		log.Printf("CeateNewClient(): %s", err.Error())
		return //	エラーがあれば、ここで終了
	}
	defer cookiejar.Save()

	tests := []struct {
		name         string
		args         args
		wantPranking *srapi.GiftRankingContribution
		wantErr      bool
	}{
		{
			name: "492 kogachan",
			args: args{
				client:   client,
				genre_id: 492,
				url_key:  "kogachan",
			},
			wantPranking: nil,
			wantErr:      false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPranking, err := srapi.ApiCdnGiftRankingContribution(tt.args.client, tt.args.genre_id, tt.args.url_key)
			if len(gotPranking.RankingList) > 0 {
				for _, v := range gotPranking.RankingList {
					log.Printf("%4d %6d %10d %s\n", v.OrderNo, v.Score, v.UserID, v.User.Name)
				}
			} else {
				log.Printf("no ranking\n")

			}
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiCdnGiftRankingContribution() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPranking, tt.wantPranking) {
				//	t.Errorf("ApiCdnGiftRankingContribution() = %v, want %v", gotPranking, tt.wantPranking)
				t.Errorf("ApiCdnGiftRankingContribution() want %v", tt.wantPranking)
			}
		})
	}
}
