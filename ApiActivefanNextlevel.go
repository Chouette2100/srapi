/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

Ver. 0.0.0 srapi.goから分離する。

*/
package srapi

import (
	"log"
	"time"

	"net/http"
	"net/url"

	"encoding/json"
)

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
