/*
!
Copyright © 2024 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php
*/
package srapi

import (
	"log"
	"io"
	"os"

	"net/http"
	"reflect"
	"testing"
)

func TestGetGenreRankingByApi(t *testing.T) {

	logfile, err := CreateLogfile("ApiGenre_ranking")
	if err != nil {
		panic("cannnot open logfile: " + err.Error())
	}
	defer logfile.Close()
	//	log.SetOutput(logfile)
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))


	client, cookiejar, err := CreateNewClient("")
	if err != nil {
		log.Printf("CeateNewClient(): %s", err.Error())
		return //	エラーがあれば、ここで終了
	}
	defer cookiejar.Save()

	type args struct {
		client  *http.Client
		genreid int
		period  string
		iscurrent bool
		pages   int
	}
	tests := []struct {
		name      string
		args      args
		wantRlist *[]GenreRanking
		wantErr   bool
	}{
		// TODO: Add test cases.
			/*
		{
			name: "daily",
			args: args{
				client:  client,
				genreid: 0,
				period:  "daily",
				iscurrent: true,
				pages:   2,
			},
			wantRlist: nil,
			wantErr:   false,
		},
			*/
		{
			name: "weekly",
			args: args{
				client:  client,
				genreid: 0,
				period:  "weekly",
				iscurrent: false,
				pages:   1,
			},
			wantRlist: nil,
			wantErr:   false,
		},
		/*
		{
			name: "monthly",
			args: args{
				client:  client,
				genreid: 0,
				period:  "monthly",
				iscurrent: true,
				pages:   1,
			},
			wantRlist: nil,
			wantErr:   false,
		},
		{
			name: "annually",
			args: args{
				client:  client,
				genreid: 0,
				period:  "annually",
				iscurrent: true,
				pages:   1,
			},
			wantRlist: nil,
			wantErr:   false,
		},
		{
			name: "all_time",
			args: args{
				client:  client,
				genreid: 0,
				period:  "all_time",
				iscurrent: true,
				pages:   1,
			},
			wantRlist: nil,
			wantErr:   false,
		},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRlist, err := GetGenreRankingByApi(tt.args.client, tt.args.genreid, tt.args.period, tt.args.iscurrent, tt.args.pages)
			log.Printf("tt.args %+v\n", tt.args)
			for _, v := range *gotRlist {
				log.Printf("%3d %9d %8d\n", v.OrderNo, v.Room.RoomID, v.Point)
			}	
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGenreRankingByApi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRlist, tt.wantRlist) {
				t.Errorf("GetGenreRankingByApi() = %v, want %v", gotRlist, tt.wantRlist)
			}
		})
	}
}
