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

// イベントの情報（イベントの詳細、参加ルームの一覧と順位、参加ルームの詳細情報）
type EventRanking struct {
	Ranking                []Ranking                `json:"ranking"`
	TotalEntries           int                      `json:"total_entries"`
	CurrentPage            int                      `json:"current_page"`
	FirstPage              int                      `json:"first_page"`
	LastPage               int                      `json:"last_page"`
	PreviousPage           int                      `json:"previous_page"`
	NextPage               int                      `json:"next_page"`
	IsEventBlock           bool                     `json:"is_event_block"`
	EventBlockLabel        string                   `json:"event_block_label"`
	EventBlockList         []EventBlockList         `json:"event_block_list"`
	EntryEventBlockID      int                      `json:"entry_event_block_id"`
	IsEventGiftRanking     bool                     `json:"is_event_gift_ranking"`
	EventGiftRankingList   []EventGiftRankingList   `json:"event_gift_ranking_list"`
	IsEventPeriodRanking   bool                     `json:"is_event_period_ranking"`
	EventPeriodRankingList []EventPeriodRankingList `json:"event_period_ranking_list"`
	IsPointVisiblePeriod   bool                     `json:"is_point_visible_period"`
	Errors                 []Errors                 `json:"errors"`
}
type Ranking struct {
	Rank            int    `json:"rank"`
	RoomID          int    `json:"room_id"`
	RoomImage       string `json:"room_image"`
	RoomName        string `json:"room_name"`
	RoomURLKey      string `json:"room_url_key"`
	RoomDescription string `json:"room_description"`
	IsOnLive        bool   `json:"is_on_live"`
	IsFollowing     bool   `json:"is_following"`
	IsOfficial      bool   `json:"is_official"`
	Point           int    `json:"point"`
}

type BlockList struct {
	BlockID int    `json:"block_id"`
	Label   string `json:"label"`
}
type EventBlockList struct {
	ShowRankLabel string      `json:"show_rank_label"`
	BlockList     []BlockList `json:"block_list"`
}

type EventGiftRankingList struct {
	EventGiftRankingID       int      `json:"event_gift_ranking_id"`
	Images                   []string `json:"images"`
	IsEventGiftRankingRanked bool     `json:"is_event_gift_ranking_ranked"`
	ExistsBlockRanking       bool     `json:"exists_block_ranking"`
}
type EventPeriodRankingList struct {
	PeriodRankingID int    `json:"period_ranking_id"`
	Label           string `json:"label"`
	StartAt         int    `json:"start_at"`
	EndAt           int    `json:"end_at"`
}

type Errors struct {
	ErrorUserMsg string `json:"error_user_msg"`
	Message      string `json:"message"`
	Code         int    `json:"code"`
}

// イベントに参加しているルームのルーム情報、順位（、獲得ポイント）を上位から30ルームずつ取得する
//
//	イベント開催前は len(ranking) は0である。
//	イベント開催中はPointには0が設定されている。
//	FIXME: イベント最終結果発表後1日間は上位30ルームの最終結果を取得することができる。
//	FIXME: イベント最終結果発表後1日をすぎると得られる獲得ポイントには0がセットされている。
//	このAPIIはイベントに対するものでありブロックイベントのブロックに対しては使用できない。
func ApiEventRanking(
	client *http.Client, //	HTTPクライアント
	eventUrlKey string, //	イベントID
	page int, //	ページ番号(1から始まる、デフォルトは1、１ページ＝30ルーム)
) (
	pranking *EventRanking,
	err error,
) {

	turl := fmt.Sprintf("https://www.showroom-live.com/api/event/%s/ranking", eventUrlKey)
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

	pranking = &EventRanking{}

	if err = json.NewDecoder(buf).Decode(pranking); err != nil {
		err = fmt.Errorf("%w(buf: %s)", err, bufstr)
		err = fmt.Errorf("json.NewDecoder(buf).Decode(pranking): %w", err)
		return
	}
	//	fmt.Printf("pranking: %+v", pranking)
	return
}

// イベントで順位が ib から ie までのルームの一覧を取得する。
func GetEventRankingByApi(
	client *http.Client, //	HTTPクライアント
	eventUrlKey string, //	イベントID
	ib int, // 順位が ib から ie までのルームの一覧を取得する。
	ie int,
) (
	pranking *EventRanking,
	err error,
) {
	if ib < 1 || ie < ib {
		err = fmt.Errorf("invalid range: ib=%d, ie=%d", ib, ie)
		return
	}

	page := (ib-1)/30 + 1 // 開始順位に対応するページ番号
	pranking = &EventRanking{}

	for {
		// ApiEventRankingを呼び出してページごとのランキングを取得
		result, err := ApiEventRanking(client, eventUrlKey, page)
		if err != nil {
			return nil, fmt.Errorf("ApiEventRanking failed: %w", err)
		}

		if (*pranking).TotalEntries == 0 {
			// 最初の結果をセット
			*pranking = *result
			pranking.Ranking = nil // ランキングは後で再構築
		}

		// 指定範囲のランキングを抽出
		for _, rank := range result.Ranking {
			if rank.Rank >= ib && rank.Rank <= ie {
				pranking.Ranking = append(pranking.Ranking, rank)
			}
		}

		// 範囲の終端に達したら終了
		if len(pranking.Ranking) >= ie-ib+1 || result.NextPage == 0 {
			break
		}

		page++ // 次のページへ
	}

	return
}
