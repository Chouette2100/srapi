/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

Ver. 0.1.0 ApiLiveCurrentUser.goをsrapi.goから分離する。ApiLiveCurrentUser()のRoomIDをstring型に変更する。

*/
package srapi

import (
	"strings"
	"time"

	"log"

	"bytes"
	//	"os"

	"encoding/json"
	"net/http"
	"net/url"
)

//	ダミーのUser-Agent
var useragent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36"

//	csrf_token ... 接続の識別子
type CsrfToken struct {
	Csrf_token string
	Data1      string
}

//	csrf_tokenを取得する。
func ApiCsrftoken(client *http.Client) (csrftoken string) {

	turl := "https://www.showroom-live.com/api/csrf_token"
	u, err := url.Parse(turl)
	if err != nil {
		log.Printf("url.Parse() returned error %s\n", err.Error())
		return
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成
	//	values.Add([クエリのキー], [値]) // key-valueを追加

	// Request を生成
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Printf("http.NewRequst() returned error %s\n", err.Error())
		return
	}

	// 組み立てたクエリを生クエリ文字列に変換して設定
	req.URL.RawQuery = values.Encode()

	// User-Agentを設定
	req.Header.Add("User-Agent", useragent)

	// Doメソッドでリクエストを投げる
	// http.Response型のポインタ（とerror）が返ってくる
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("client.Do() returned error %s\n", err.Error())
		return
	}

	// 関数を抜ける際に接続を切断し、リソースを解放するため必ずresponse.Bodyをcloseする
	defer resp.Body.Close()

	var ct CsrfToken
	if err := json.NewDecoder(resp.Body).Decode(&ct); err != nil {
		log.Printf("json.NewDecoder() returned error %s\n", err.Error())
		return
	}
	csrftoken = ct.Csrf_token

	return

}

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

//	ファンレベルの進捗状況の詳細
type ActiveFanNextLevel struct {
	Level      int //	現在のファンレベル
	Next_level struct {
		Level      int //	次のファンレベル
		Conditions []struct {
			Condition_details []struct {
				Goal  int    //	次のレベルになるのに必要な要素（視聴時間、ギフト、コメントの数）
				Unit  string //	単位（分、pt、G、コメント）
				Value int    //	（レベルアップに使った分を除く）現在の実績（手持ちの要素数）(視聴時間、ギフト、コメントの数)
				Label string //	次のレベルになるのに必要な要素の詳細（視聴時間、無料ギフト、有料ギフト、コメント数）
			}
			Label string //	次のレベルになるのに必要な要素（視聴時間、有料ギフト or 無料ギフト、コメント数）
		}
	}
	Title_id int //	不明
}

//	ファンレベルの進捗状況を得るためのAPI（/api/active_fan/next_level）を実行する。
func ApiActivefanNextlevel(
	client *http.Client, //	HTTPクライアント
	userid string, //	リスナーの識別子、ログインAPIの実行結果のJSONにあるUser_id（他にも取得方法はある）
	roomid string, //	配信ルームの識別子（プロフィールやファンルームのURLの最後にある6桁程度の数）
) (
	afnl ActiveFanNextLevel, //	ファンレベルの進捗状況の詳細
	status int, //	実行結果	0: 正常
) {

	status = 0

	turl := "https://www.showroom-live.com/api/active_fan/next_level"
	u, err := url.Parse(turl)
	if err != nil {
		log.Printf("ApiActivefanNextlevel() url.Parse() returned error %s\n", err.Error())
		status = 1
		return
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成
	values.Add("user_id", userid)
	values.Add("room_id", roomid)
	values.Add("ym", time.Now().Format("200601"))

	//	log.Printf("values=%+v\n", values)

	// Request を生成
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Printf("ApiActivefanNextlevel() http.NewRequst() returned error %s\n", err.Error())
		status = 2
		return
	}

	// 組み立てたクエリを生クエリ文字列に変換して設定
	req.URL.RawQuery = values.Encode()

	// User-Agentを設定
	req.Header.Add("User-Agent", useragent)

	// Doメソッドでリクエストを投げる
	// http.Response型のポインタ（とerror）が返ってくる
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("ApiActivefanNextlevel() client.Do() returned error %s\n", err.Error())
		status = 3
		return
	}

	// 関数を抜ける際に接続を切断し、リソースを解放するため必ずresponse.Bodyをcloseする
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&afnl); err != nil {
		log.Printf("json.NewDecoder() returned error %s\n", err.Error())
		status = 4
		return
	}

	return

}
