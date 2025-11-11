// Copyright © 2025 chouette.21.00@gmail.com
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

// イベントの参加ルーム情報
type EventRooms struct {
	CurrentPage          int       `json:"current_page"`
	FirstPage            int       `json:"first_page"`
	IsPointVisiblePeriod bool      `json:"is_point_visible_period"`
	LastPage             int       `json:"last_page"`
	NextPage             int       `json:"next_page"`
	Rooms                []RoomsEv `json:"rooms"`
	TotalEntries         int       `json:"total_entries"`
	Errors               []Errors  `json:"errors"`
}
type RoomsEv struct {
	IsFollowing     bool   `json:"is_following"`
	IsOfficial      bool   `json:"is_official"`
	IsOnLive        bool   `json:"is_on_live"`
	Point           int    `json:"point"`
	RoomDescription string `json:"room_description"`
	RoomID          int    `json:"room_id"`
	RoomImage       string `json:"room_image"`
	RoomName        string `json:"room_name"`
	RoomURLKey      string `json:"room_url_key"`
}

// イベントに参加しているルームのルーム情報を分割して取得する(順位は無関係)
//
//	イベント開催前から使えるが、順位は取得できないのでイベント開催後の使用が限定される
func ApiEventRooms(
	client *http.Client, //	HTTPクライアント
	eventUrlKey string, //	イベントID
	page int, //	ページ番号(1から始まる、デフォルトは1、１ページ＝30ルーム)
) (
	prooms *EventRooms,
	err error,
) {

	turl := fmt.Sprintf("https://www.showroom-live.com/api/event/%s/rooms", eventUrlKey)
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成

	// クエリを組み立て
	values.Add("page", fmt.Sprintf("%d", page)) // key-valueを追加

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

	prooms = &EventRooms{}

	if err = json.NewDecoder(buf).Decode(prooms); err != nil {
		err = fmt.Errorf("%w(buf: %s)", err, bufstr)
		err = fmt.Errorf("json.NewDecoder(buf).Decode(pranking): %w", err)
		return
	}
	//	fmt.Printf("pranking: %+v", pranking)
	return
}

// イベントで順位が ib から ie までのルームの一覧を取得する。
func GetEventRoomsByApi(
	client *http.Client, //	HTTPクライアント
	eventUrlKey string, //	イベントID
	ib int, // 出現順が ib から ie までのルームの一覧を取得する。
	ie int,
) (
	prooms *EventRooms,
	err error,
) {
	if ib < 1 || ie < ib {
		err = fmt.Errorf("invalid range: ib=%d, ie=%d", ib, ie)
		return
	}

	page := (ib-1)/30 + 1 // 開始順位に対応するページ番号
	prooms = &EventRooms{}

	for {
		// ApiEventRoomsを呼び出してページごとのルーム情報を取得
		result, err := ApiEventRooms(client, eventUrlKey, page)
		if err != nil {
			return nil, fmt.Errorf("ApiEventRooms failed: %w", err)
		}
		if len(result.Errors) > 0 {
			return nil, fmt.Errorf("ApiEventRooms returned errors: %+v", result.Errors)
		}

		if (*prooms).TotalEntries == 0 {
			// 最初の結果をセット
			*prooms = *result
			prooms.Rooms = nil // ルーム情報は後で再構築
			prooms.TotalEntries = result.TotalEntries
		}

		// 指定範囲のランキングを抽出
		for _, room := range result.Rooms {
			prooms.Rooms = append(prooms.Rooms, room)
			// 範囲の終端に達したら終了
			if len(prooms.Rooms) >= ie-ib+1 {
				break
			}
		}

		// すべてのページを処理したら終了
		if len(prooms.Rooms) >= ie-ib+1 || result.NextPage == 0 {
			break
		}

		page++ // 次のページへ
	}

	return
}
