// Copyright © 2024 chouette.21.00@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package srapi

import (
	"bytes"
	"fmt"

	"encoding/json"
	"net/http"
	"net/url"
)

type RoomGenreList struct {
	RoomGenre []RoomGenre `json:"room_genre"`
}
type RoomGenre struct {
	GenreID int    `json:"genre_id"`
	OrderNo int    `json:"order_no"`
	Name    string `json:"name"`
	IsEntry bool   `json:"is_entry"`
}

// ジャンルの一覧を取得する
func ApiGenrerankingRoomGenre(
	client *http.Client, //	HTTPクライアント
) (
	pgenre *RoomGenreList,
	err error,
) {
	turl := "https://www.showroom-live.com/api/genre_ranking/room_genre"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成

	//	クエリを組み立て
	// values.Add("limit", fmt.Sprintf("%d", limit)) // key-valueを追加

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

	//	fmt.Printf("bufstr: %s", bufstr)

	pgenre = &RoomGenreList{}

	if err = json.NewDecoder(buf).Decode(pgenre); err != nil {
		err = fmt.Errorf("%w(buf: %s)", err, bufstr)
		err = fmt.Errorf("json.NewDecoder(buf).Decode(pranking): %w", err)
		return
	}
	//	fmt.Printf("pgenre: %+v", pgenre)
	return
}
