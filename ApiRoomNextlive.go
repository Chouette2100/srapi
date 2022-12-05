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

type RoomNextlive struct {
	Epoch int64  //	次回配信開始時刻（UnixTime）
	Text  string //	次回配信開始時刻（ "1/2 15:04" ）
}

//	次回配信時刻を取得する
func ApiRoomNextlive(
	client *http.Client, //	HTTPクライアント
	room_id int, //	ルーム識別子
) (
	roomnextlive *RoomNextlive,
	err error,
) {

	turl := "https://www.showroom-live.com/api/room/next_live"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{}         // url.Valuesオブジェクト生成
	values.Add("room_id", fmt.Sprintf("%d",room_id)) // key-valueを追加

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

	if err = json.NewDecoder(resp.Body).Decode(&roomnextlive); err != nil {
		err = fmt.Errorf("json.NewDecoder(resp.Body).Decode(&roomnextlive): %w", err)
		return
	}
	return
}
