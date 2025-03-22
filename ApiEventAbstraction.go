// Copyright © 2025 chouette.21.00@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package srapi

import (
	"bytes"
	"fmt"
	"log"

	"encoding/json"

	"net/http"
	"net/url"
)

type EventAbstraction struct {
	EventID           int           `json:"event_id"`
	EventOrganizerID  int           `json:"event_organizer_id"`
	Title             string        `json:"title"`
	Abstraction       string        `json:"abstraction"`
	EventType         string        `json:"event_type"`
	HeadImageL        string        `json:"head_image_l"`
	HeadImageM        string        `json:"head_image_m"`
	OfferStartAt      int           `json:"offer_start_at"`
	OfferEndAt        int           `json:"offer_end_at"`
	EventStartAt      int           `json:"event_start_at"`
	EventEndAt        int           `json:"event_end_at"`
	RequiredLevelFrom int           `json:"required_level_from"`
	RequiredLevelTo   int           `json:"required_level_to"`
	QuestLevels       []QuestLevels `json:"quest_levels"`
	QuestLevelMax     int           `json:"quest_level_max"`
}
type AchievedRooms struct {
	Rooms []any `json:"rooms"`
}
type QuestLevels struct {
	Level             int           `json:"level"`
	Point             int           `json:"point"`
	Description       string        `json:"description"`
	TheNumOfItems     int           `json:"the_num_of_items"`
	TheNumOfRestItems int           `json:"the_num_of_rest_items"`
	Remarks           string        `json:"remarks"`
	AchievedRooms     AchievedRooms `json:"achieved_rooms"`
}

// イベントの概要を取得する
func ApiEventAbstraction(
	client *http.Client,
	eventid string,
) (
	ea *EventAbstraction,
	err error,
) {

	turl := "https://www.showroom-live.com/api/event/" + eventid + "/abstraction"

	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	// 例 values := url.Values{}         // url.Valuesオブジェクト生成
	//    values.Add("page", strconv.Itoa(page)) // key-valueを追加

	// Request を生成
	var req *http.Request
	req, err = http.NewRequest("GET", u.String(), nil)
	if err != nil {
		err = fmt.Errorf("http.NewRequest(): %w", err)
		return
	}

	// 組み立てたクエリを生クエリ文字列に変換して設定
	// req.URL.RawQuery = values.Encode()

	// User-Agentを設定
	req.Header.Add("User-Agent", useragent)
	req.Header.Add("Accept-Language", "ja-JP")

	// Doメソッドでリクエストを投げる
	// http.Response型のポインタ（とerror）が返ってくる
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		err = fmt.Errorf("client.Do(): %w", err)
		return
	}
	// 関数を抜ける際に接続を切断し、リソースを解放するため必ずresponse.Bodyをcloseする
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bufstr := buf.String()

	//	JSONをデコードする。
	//	次の記事を参考にさせていただいております。
	//		Go言語でJSONに泣かないためのコーディングパターン
	//		https://qiita.com/msh5/items/dc524e38073ed8e3831b

	ea = new(EventAbstraction)
	decoder := json.NewDecoder(buf)
	if err = decoder.Decode(ea); err != nil {
		log.Printf("decoder.Decode(&result) err: %s", err.Error())
		log.Printf("eventid= %s", eventid)
		log.Printf("bufstr= %s", bufstr)
		err = fmt.Errorf("decoder.Decode(&result) err: %w", err)
		return
	}

	return
}