// Copyright © 2024 chouette.21.00@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package srapi

import (
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/Chouette2100/exsrapi"
)

func TestApiiCdnUserGiftRanking(t *testing.T) {
	type args struct {
		client   *http.Client
		genre_id int
		limit    int
	}

	logfile, err := exsrapi.CreateLogfile("TestApiCdnGiftRanking")
	if err != nil {
		panic("cannnot open logfile: " + err.Error())
	}
	defer logfile.Close()
	//	log.SetOutput(logfile)
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))

	client, cookiejar, err := exsrapi.CreateNewClient("")
	if err != nil {
		log.Printf("exsrapi.CeateNewClient(): %s", err.Error())
		return //	エラーがあれば、ここで終了
	}
	defer cookiejar.Save()

	tests := []struct {
		name         string
		args         args
		wantPranking *CdnUserGiftRanking
		wantErr      bool
	}{
		{
			name: "case500-タレント",
			args: args{
				client:   client,
				genre_id: 500,
				limit:    10,
			},
			wantPranking: nil,
			wantErr:      false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPranking, err := ApiCdnGiftRanking(tt.args.client, tt.args.genre_id, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiCdnGiftRanking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPranking, tt.wantPranking) {
				for _, v := range gotPranking.RankingList {
					log.Printf("%3d%10d %10d %s\n", v.Rank, v.Score, v.Room.ID, v.Room.Name)
				}
				//	t.Errorf("ApiiCdnUserGiftRanking() = %v, want %v", gotPranking, tt.wantPranking)
			}
		})
	}
}

func TestApiCdnUserGiftRanking(t *testing.T) {
	type args struct {
		client   *http.Client
		genre_id int
		limit    int
	}

	logfile, err := exsrapi.CreateLogfile("TestApiCdnGiftRanking")
	if err != nil {
		panic("cannnot open logfile: " + err.Error())
	}
	defer logfile.Close()
	//	log.SetOutput(logfile)
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))

	client, cookiejar, err := exsrapi.CreateNewClient("")
	if err != nil {
		log.Printf("exsrapi.CeateNewClient(): %s", err.Error())
		return //	エラーがあれば、ここで終了
	}
	defer cookiejar.Save()

	tests := []struct {
		name         string
		args         args
		wantPranking *CdnUserGiftRanking
		wantErr      bool
	}{
		{
			name: "case206-ユーザー",
			args: args{
				client:   client,
				genre_id: 206,
				limit:    10,
			},
			wantPranking: nil,
			wantErr:      false,
		},
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPranking, err := ApiCdnUserGiftRanking(tt.args.client, tt.args.genre_id, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiCdnUserGiftRanking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPranking, tt.wantPranking) {
				for _, v := range gotPranking.RankingList {
					log.Printf("%3d%10d %10d %s\n", v.Rank, v.Score, v.User.ID, v.User.Name)
				}
				//	t.Errorf("ApiCdnUserGiftRanking() = %v, want %v", gotPranking, tt.wantPranking)
			}
		})
	}
}
