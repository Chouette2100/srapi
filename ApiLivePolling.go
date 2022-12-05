/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php


Ver. 0.1.0

*/
package srapi

import (
	"bytes"
	"fmt"

	"encoding/json"
	"net/http"
	"net/url"
)

type LivePolling struct {
	Is_login          bool //	ログインしているか？
	Show_login_dialog int
	Online_user_num   int //	視聴者数
	Active_fan        struct {
		Can_boostable bool
		User          struct {
			Before_level  int
			Titile_id     int
			Current_level int
		}
		Room struct {
			Total_user_count int    //
			Fan_name         string //	ファンの名称
		}
	}
	Live_watch_incentive struct {
		Ok         int    //	1:星・種をもらえた
		Is_amateur string // 	"1": フリー
	}
	Toast struct {
		Message string //	星・種集めのトーストの文面はここにある
	}
}

//	星・種をもらえたことを確認する。
//	配信状況を確認し、room_url_key から　room_id を取得する。
func ApiLivePolling(
	client *http.Client, //	HTTPクライアント
	room_id int, //	ルーム識別子
) (
	livepolling *LivePolling,
	err error,
) {

	turl := "https://www.showroom-live.com/api/live/polling"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{}                            // url.Valuesオブジェクト生成
	values.Add("room_id", fmt.Sprintf("%d", room_id)) // key-valueを追加

	// Request を生成
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		err = fmt.Errorf("http.NewRequest(): %w", err)
		return
	}

	// 組み立てたクエリを生クエリ文字列に変換して設定
	req.URL.RawQuery = values.Encode()

	// User-Agentを設定
	req.Header.Add("User-Agent", useragent)
	req.Header.Add("Accept-Language", "ja-JP")

	// Doメソッドでリクエストを投げる
	// http.Response型のポインタ（とerror）が返ってくる
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("client.Do(): %w", err)
		return
	}
	// 関数を抜ける際に接続を切断し、リソースを解放するため必ずresponse.Bodyをcloseする
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bufstr := buf.String()

	livepolling = new(LivePolling)
	if err = json.NewDecoder(buf).Decode(livepolling); err != nil {
		err = fmt.Errorf("%w(resp.Body: %s)", err, bufstr)
		err = fmt.Errorf("json.NewDecoder(resp.Body).Decode(livepolling): %w", err)
		return
	}
	return
}
