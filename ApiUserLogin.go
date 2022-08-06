/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

Ver. 0.0.0 srapi.goから分離する。

*/
package srapi

import (
	"bytes"
	"log"
	"strings"

	"net/http"
	"net/url"

	"encoding/json"
)

//	ログインのAPIの結果
type UserLogin struct {
	Ok            int    //	1: ログインできた 0:ログインできなかった
	Is_room_owner int    //	未確認、おそらく配信者登録をしている(1)か否(0)かだと思う。
	User_id       string //	ログインできたときはリスナーの識別子（数字からなる文字列）
	Room_id       string //	未確認、おそらく配信ルームの識別子（数字からなる文字列 ？）
	Account_id    string //	ログインできたときはログインアカウント
}

//	SHOWROOMのサービスにログインする
func ApiUserLogin(
	client *http.Client, //	HTTPクライアント
	csrftoken string, //	csrftoken
	accountid string, //	ログインアカウント
	password string, //	ログインパスワード
) (
	userlogin UserLogin, //	ログイン結果
	status int, //	終了ステータス 0: 正常
) {

	status = 0 //	終了ステータス

	// POSTメソッド
	turl := "https://www.showroom-live.com/user/login"
	u, err := url.Parse(turl)
	if err != nil {
		log.Printf("url.Parse() returned error %s\n", err.Error())
		return
	}

	// url.Values{}でPOSTで送信する入れ物を準備
	values := url.Values{}

	// Add()でPOSTで送信するデータを作成
	values.Add("csrf_token", csrftoken)
	values.Add("account_id", accountid)
	values.Add("password", password)
	// 特殊文字や日本語をエンコード
	//	fmt.Println(values.Encode())

	// Request を生成
	req, err := http.NewRequest("POST", u.String(), strings.NewReader(values.Encode()))
	if err != nil {
		log.Printf("http.NewRequst() returned error %s\n", err.Error())
		return
	}

	// Content-Typeを設定
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// User-Agentを設定
	req.Header.Add("User-Agent", useragent)

	// Doメソッドでリクエストを投げる
	// http.Response型のポインタ（とerror）が返ってくる
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("client.Do() returned error %s\n", err.Error())
		return
	}

	// 接続を切断し、リソースを開放する。
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bufstr := buf.String()

	if err := json.NewDecoder(buf).Decode(&userlogin); err != nil {
		log.Printf("json.NewDecoder() returned error %s\n", err.Error())
		log.Printf("resp.Body=%s\n", bufstr)
		status = -2
		return
	}
	//	log.Printf("resp.Body=%s\n", bufstr)
	return

}
