/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php


Ver. 0.1.0

*/
package srapi

import (
	"fmt"

	"encoding/json"

	"net/http"
	"net/url"
)

type RoomStatus struct {
	Started_at               int64 //	配信開始時刻（UnixTime）
	Is_live                  bool  //	配信中か？
	Live_status              int	// 0|2
	Is_enquete               bool //	アンケートが行われているか？
	Live_id                  int  //	配信の識別子
	Is_official              bool	//	公式か？
	Genre_id                 int	//	配信ジャンル（Freeだったら200みたいなの）
	Room_id                  int    //	ルームの識別子
	Room_name                string //	ルーム名
	Room_url_key             string	//	配信時URLの最後のフィールド
	Is_owner                 bool
	Is_fav                   bool
	Youtube_id               string	//	非配信時Youtube動画（ https://www.youtube.com/watch?v=4{Youtube_id}）
	Did_send_live_bad_report bool
	Can_comment              bool
	Background_image_url     string
	Video_type               int
	Broadcast_host           string
	Broadcast_port           int
	Broadcast_key            string
	Image_s                  string
	Nsta_owner               bool
	Live_type                int
	Live_user_key            string
	Share                    struct {
		Twitter struct {
			Text string
			Url  string
		}
	}
}

//	配信状況を確認し、room_url_key から　room_id を取得する。
func ApiRoomStatus(
	client *http.Client, //	HTTPクライアント
	room_url_key string, //	ルームのURLの最後のフィールド
) (
	roomstatus *RoomStatus,
	err error,
) {

	turl := "https://www.showroom-live.com/api/room/status"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{}                   // url.Valuesオブジェクト生成
	values.Add("room_url_key", room_url_key) // key-valueを追加

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

	if err = json.NewDecoder(resp.Body).Decode(&roomstatus); err != nil {
		err = fmt.Errorf("json.NewDecoder(resp.Body).Decode(&roomstatus): %w", err)
		return
	}
	return
}
