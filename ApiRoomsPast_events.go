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
	"log"
	"strconv"

	"encoding/json"
	"net/http"
	"net/url"
)

type RoomsPastevents struct {
	LastPage       int         `json:"last_page"`
	CurrentPage    int         `json:"current_page"`
	NextPage       int         `json:"next_page"`
	FavoriteEvents interface{} `json:"favorite_events"`
	TotalEntries   int         `json:"total_entries"`
	PreviousPage   int         `json:"previous_page"`
	Events         []struct {
		EventName   string `json:"event_name"`
		EventID     int    `json:"event_id"`
		EventType   string `json:"event_type"`
		Image       string `json:"image"`
		EndedAt     int    `json:"ended_at"`
		IsFavorite  int    `json:"is_favorite"`
		StartedAt   int    `json:"started_at"`
		ShowRanking int    `json:"show_ranking"`
	} `json:"events"`
	FirstPage int `json:"first_page"`
}

// 指定したルームのこれまで参加したイベントの一覧を取得する。
// ルームIDとページ番号を指定する。
// ページ番号は1から始まる。
func ApiRoomsPast_events(
	client *http.Client,
	roomid int,
	page int,
) (
	rpe *RoomsPastevents,
	err error,
) {
	//	https://www.showroom-live.com/api/rooms/nnnnnn/past_events?page=1
	turl := "https://www.showroom-live.com/api/rooms/" + strconv.Itoa(roomid) + "/past_events"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成

	// クエリを組み立て
	values.Add("page", strconv.Itoa(page)) // key-valueを追加

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

	rpe = &RoomsPastevents{}

	if err = json.NewDecoder(buf).Decode(rpe); err != nil {
		err = fmt.Errorf("%w(buf: %s)", err, bufstr)
		err = fmt.Errorf("json.NewDecoder(buf).Decode(rpe): %w", err)
		return
	}
	//	fmt.Printf("RoomsPastevents: %+v", *rpe)

	return
}

// 指定したルームが過去にエントリーしたすべてのイベントを取得する
func GetRoomsPasteventsByApi(
	client *http.Client,
	roomid int,
) (
	roomspastevents *RoomsPastevents,
	err error,
) {


	roomspastevents = &RoomsPastevents{}
	te := 0

	for page := 1; ; page++ {

		var rpe *RoomsPastevents
		rpe, err = ApiRoomsPast_events(client, roomid, page)
		if err != nil {
			err = fmt.Errorf("ApiRoomsPast_events(): %w", err)
			return
		}
		roomspastevents.Events = append(roomspastevents.Events, rpe.Events...)
		if page == 1 {
			te = rpe.TotalEntries
		}
		if rpe.LastPage == page {
			break
		}
	}

	//	fmt.Printf("RoomsPastevents: %+v", *roomspastevents)
	if len(roomspastevents.Events) != te {
		log.Printf(" len(roomspastevents.Events) = %d, te = %d\n", len(roomspastevents.Events), te)
		err = fmt.Errorf("len(roomspastevents.Events) != TotalEntries")
		return
	} else {
		roomspastevents.TotalEntries = te
	}

	return
}
