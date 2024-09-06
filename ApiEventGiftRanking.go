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

type EventGiftRanking struct {
	RankingType int `json:"ranking_type"` //	1: ？
	RankingList []struct {
		OrderNo int `json:"order_no"` //	順位（1〜）
		Score   int `json:"score"`    //	獲得ポイント
		Room    struct {
			IsOfficial      int    `json:"is_official"`      //	公式？
			RoomURLKey      string `json:"room_url_key"`     //	配信画面のURLから"https://www.showroom-live.com/r/"を除いたもの
			RoomDescription string `json:"room_description"` //	ルーム説明
			ImageS          string `json:"image_s"`          //	サムネール
			RoomName        string `json:"room_name"`        //	ルーム名
			RoomID          string `json:"room_id"`          //	ルームID（整数 〜6桁）
			IsOnline        int    `json:"is_online"`        //	配信中？
		} `json:"room"`
		CreatedAt          int    `json:"created_at"`            //	2024-09-06 09:25 に gift_id=1497のとき 1725582065(2024-09-06 09:21:05)
		UpdatedAt          int    `json:"updated_at"`            //	〃
		EventGiftRankingID int    `json:"event_gift_ranking_id"` //	gift_id
		Rank               int    `json:"rank"`                  //	順位（= OrderNo ？）
		RoomID             string `json:"room_id"`               //	ルームID（ = Room.RoomID ？ ）
	} `json:"ranking_list"`
	GiftData []struct {
		GiftID      int    `json:"gift_id"` //	ギフトID
		Name        string `json:"name"`    //	ギフト名
		IsAnimation int    `json:"is_animation"`
		Path        string `json:"path"` //	ギフト画像
	} `json:"gift_data"`
	Errors []struct { //	例えば gift_id が（整数でなく）アルファベットの場合
		ErrorUserMsg string `json:"error_user_msg"`
		Message      string `json:"message"`
		Code         int    `json:"code"`
	} `json:"errors"`
}

// イベントギフトランキングを取得する
//	(ギフトランキング、ユーザーギフトランキングとは別のものです)
//	https://www.showroom.com/api/event_gift_ranking/nnnnnnn		nnnnnn: gift_id
//	取得に期限はないようで EventRankingID=1 にたいしてもそれらしい結果が得られている。
func ApiEventGiftRanking(
	client *http.Client, //	HTTPクライアント
	gift_id int, //	ギフトID
) (
	pranking *EventGiftRanking,
	err error,
) {
	turl := fmt.Sprintf("https://www.showroom-live.com/api/event_gift_ranking/%d", gift_id)
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成

	// クエリを組み立て
	//	values.Add("room_id", fmt.Sprintf("%d", roomid)) // key-valueを追加

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

	pranking = &EventGiftRanking{}

	if err = json.NewDecoder(buf).Decode(pranking); err != nil {
		err = fmt.Errorf("%w(buf: %s)", err, bufstr)
		err = fmt.Errorf("json.NewDecoder(buf).Decode(pranking): %w", err)
		return
	}
	//	fmt.Printf("pranking: %+v", pranking)
	return
}
