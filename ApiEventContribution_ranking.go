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

type Contribution_ranking struct {
	Ranking []struct {
		AvatarID  int    `json:"avatar_id"`
		Point     int    `json:"point"`
		AvatarURL string `json:"avatar_url"`
		Name      string `json:"name"`
		UserID    int    `json:"user_id"`
		Rank      int    `json:"rank"`
	} `json:"ranking"`
	Me      any `json:"me"`
	MyPoint int `json:"my_point"`
	Event   struct {
		EndedAt     int    `json:"ended_at"`
		EventName   string `json:"event_name"`
		EventType   string `json:"event_type"`
		StartedAt   int    `json:"started_at"`
		EventURL    string `json:"event_url"`
		ShowRanking int    `json:"show_ranking"`
		Image       string `json:"image"`
	} `json:"event"`
}

//	リスナー別の貢献ポイントを取得する
func ApiEventContribution_ranking(
	client *http.Client, //	HTTPクライアント
	ieventid int,
	roomid int,
) (
	pranking *Contribution_ranking,
	err error,
) {

	turl := "https://www.showroom-live.com/api/event/contribution_ranking"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成

	// クエリを組み立て
	values.Add("event_id", fmt.Sprintf("%d", ieventid)) // key-valueを追加
	values.Add("room_id", fmt.Sprintf("%d", roomid)) // key-valueを追加

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

	pranking = &Contribution_ranking{}

	if err = json.NewDecoder(buf).Decode(pranking); err != nil {
		err = fmt.Errorf("%w(buf: %s)", err, bufstr)
		err = fmt.Errorf("json.NewDecoder(buf).Decode(pranking): %w", err)
		return
	}
	//	fmt.Printf("pranking: %+v", pranking)
	return
}
