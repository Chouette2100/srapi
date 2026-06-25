package srapi_test

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/Chouette2100/srapi/v2"
	"github.com/Chouette2100/srcom"
)

func TestApiActivefanNextlevel(t *testing.T) {

	logfile, err := srcom.CreateLogfile3("ApiActivefanNextlevel", srapi.Version)
	if err != nil {
		panic("cannnot open logfile: " + err.Error())
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	client, cookiejar, err := srapi.CreateNewClient("")
	if err != nil {
		log.Printf("CeateNewClient(): %s", err.Error())
		return //	エラーがあれば、ここで終了
	}
	defer cookiejar.Save()

	//	csrftokenを取得する
	csrftoken, err := srapi.ApiCsrftoken(client)
	if err != nil {
		err = fmt.Errorf("ApiCsrftoken: %w", err)
		return
	}

	// 環境変数からアカウントとパスワードを取得する
	/*
	acct, pswd, err := GetAccountAndPassword()
	if err != nil {
		log.Printf("GetAccountAndPassword(): %s", err.Error())
		return //	エラーがあれば、ここで終了
	}
		*/
	acct := "Seppina1"
	pswd := "Dgn895=="

	//	SHOWROOMのサービスにログインする。
	var ul srapi.UserLogin
	ul, err = srapi.ApiUserLogin(client, csrftoken, acct, pswd)
	if err != nil {
		err = fmt.Errorf("ApiUserLogin: %w", err)
		return
	} else {
		log.Printf("login status. Ok = %d User_id=%s\n", ul.Ok, ul.User_id)
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		client  *http.Client
		userid  string
		roomid  string
		want    srapi.ActiveFanNextLevel
		wantErr bool
	}{
		{
			name:    "test1",
			client:  client,
			userid:  ul.User_id,
			roomid:  "570195",
			want:    srapi.ActiveFanNextLevel{},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	// for _, tt := range tests {
	// 100回繰り返す
	for range [100]int{} {
		tt := tests[0]
		t.Run(tt.name, func(t *testing.T) {
			// got, gotErr := srapi.ApiActivefanNextlevel(tt.client, tt.userid, tt.roomid)
			got, _ := srapi.ApiActivefanNextlevel(tt.client, tt.userid, tt.roomid)
			log.Printf("got= %+v\n", got)
			/*
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ApiActivefanNextlevel() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ApiActivefanNextlevel() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("ApiActivefanNextlevel() = %v, want %v", got, tt.want)
			}
				*/
		})
		time.Sleep(1 * time.Minute) // 1秒待機
	}
}

// 環境変数からアカウントとパスワードを取得する
func GetAccountAndPassword() (
	acct string,
	pswd string,
	err error,
) {
	SRACCT := os.Getenv("SRACCT")
	SRPSWD := os.Getenv("SRPSWD")

	if SRACCT == "" || SRPSWD == "" {
		err = fmt.Errorf("環境変数 SRACCT または SRPSWD が設定されていません。")
		return "", "", err
	}
	return SRACCT, SRPSWD, nil
}
