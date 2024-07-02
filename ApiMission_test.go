//	Copyright © 2024 chouette.21.00@gmail.com
//	Released under the MIT license
//	https://opensource.org/licenses/mit-license.php

package srapi

import (
	"log"

	"net/http"
	"reflect"
	"testing"

	"github.com/chouette2100/exsrapi"
)

func TestApiMission(t *testing.T) {

	type SRConfig struct {
		SRacct string //	SHOWROOMのアカウント名
		SRpswd string //	SHOWROOMのパスワード
	}

	//	設定ファイルを読み込む。
	var srconfig SRConfig
	err := exsrapi.LoadConfig("SRConfig.yml", &srconfig)
	if err != nil {
		log.Printf("LoadConfig: %s\n", err.Error())
		return
	}

	//	cookiejarがセットされたHTTPクライアントを作る
	client, jar, err := exsrapi.CreateNewClient(srconfig.SRacct)
	if err != nil {
		log.Printf("CreateNewClient() returned error %s\n", err.Error())
		return
	}
	//	すべての処理が終了したらcookiejarを保存する。
	defer jar.Save()

	//	SHOWROOMのサービスにログインし、ユーザIDを取得する。
	_, err = exsrapi.LoginShowroom(client, srconfig.SRacct, srconfig.SRpswd)
	if err != nil {
		log.Printf("exsrapi.LoginShowroom: %s\n", err.Error())
		return
	}

	type args struct {
		client  *http.Client
		room_id string
	}
	tests := []struct {
		name         string
		args         args
		wantPmission *Mission
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "mission1",
			args: args{
				client:  client,
				room_id: "87911",
			},
			wantPmission: &Mission{},
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPmission, err := ApiMission(tt.args.client, tt.args.room_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiMission() error = %v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPmission, tt.wantPmission) {
				t.Errorf("ApiMission() = %+v, want %+v", gotPmission, tt.wantPmission)
			}
		})
	}
}
