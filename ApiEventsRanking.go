/*
!
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

type Eventranking struct {
	Ranking []struct {
		Point int `json:"point"`
		Room  struct {
			Name        string `json:"name"`
			ImageSquare string `json:"image_square"`
			RoomID      int    `json:"room_id"`
			Image       string `json:"image"`
		} `json:"room"`
		Rank int `json:"rank"`
	} `json:"ranking"`
	TargetRoomRanking struct {
		Gap   int `json:"gap"`
		Point int `json:"point"`
		Room  struct {
			Name        string `json:"name"`
			ImageSquare string `json:"image_square"`
			RoomID      int    `json:"room_id"`
			Image       string `json:"image"`
		} `json:"room"`
		Rank     int `json:"rank"`
		LowerGap int `json:"lower_gap"`
	} `json:"target_room_ranking"`
	Event struct {
		EndedAt     int    `json:"ended_at"`
		EventName   string `json:"event_name"`
		StartedAt   int    `json:"started_at"`
		EventType   string `json:"event_type"`
		EventURL    string `json:"event_url"`
		Image       string `json:"image"`
		ShowRanking int    `json:"show_ranking"`
	} `json:"event"`
}

// フォローしているルームの一覧を指定した範囲で取得する。
func ApiEventsRanking(
	client *http.Client, //	HTTPクライアント
	ieventid int,
	roomid int,
	blockid int,
) (
	pranking *Eventranking,
	err error,
) {

	turl := fmt.Sprintf("https://www.showroom-live.com/api/events/%d/ranking", ieventid)
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成

	// クエリを組み立て
	values.Add("room_id", fmt.Sprintf("%d", roomid)) // key-valueを追加

	if blockid > 0 {
		//	ブロックイベントの場合
		values.Add("event_block_id", fmt.Sprintf("%d", blockid)) // key-valueを追加
	}

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

	pranking = &Eventranking{}

	if err = json.NewDecoder(buf).Decode(pranking); err != nil {
		err = fmt.Errorf("%w(buf: %s)", err, bufstr)
		err = fmt.Errorf("json.NewDecoder(buf).Decode(pranking): %w", err)
		return
	}
	//	fmt.Printf("pranking: %+v", pranking)
	return
}
