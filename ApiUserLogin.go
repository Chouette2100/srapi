/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

Ver. 0.0.0 srapi.goから分離する。
Ver. 1.0.0 戻り値の staus を error 変更する

*/
package srapi

import (
	"bytes"
	"fmt"
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
	err error, //	nil: 正常終了
) {

	// POSTメソッド
	turl := "https://www.showroom-live.com/user/login"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parsse: %w", err)
		return userlogin, err
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
		err = fmt.Errorf("http.NewRequest: %w", err)
		return userlogin, err
	}

	// Content-Typeを設定
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// User-Agentを設定
	req.Header.Add("User-Agent", useragent)
	req.Header.Add("Accept-Language", "ja-JP")

	// Doメソッドでリクエストを投げる
	// http.Response型のポインタ（とerror）が返ってくる
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("client.Do: %w", err)
		return userlogin, err
	}

	// 接続を切断し、リソースを開放する。
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bufstr := buf.String()

	if err := json.NewDecoder(buf).Decode(&userlogin); err != nil {
		err = fmt.Errorf("%w(resp.Body: %s)", err, bufstr)
		err = fmt.Errorf("json.NewDecoder: %w", err)
		
		return userlogin, err
	}
	//	log.Printf("resp.Body=%s\n", bufstr)
	return userlogin, nil

}
