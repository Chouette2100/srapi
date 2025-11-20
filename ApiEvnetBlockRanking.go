/*
!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

Ver. 0.1.0
*/
package srapi

import (
	//"errors"
	"fmt"
	//	"log"
	//"strconv"
	//"strings"

	"net/http"
	"net/url"

	"encoding/json"
)

type EventBlockRanking struct {
	TotalEntries     int                `json:"total_entries"`
	EntriesPerPage   int                `json:"entries_per_page"`
	CurrentPage      int                `json:"current_page"`
	BlockRankingList []BlockRankingList `json:"block_ranking_list"`
}
type BlockRankingList struct {
	//	boolとintのinterface{}型はデータ取得時にbool型として値を設定しておくこと
	//	該当する変数を使うときは一律に if br.Is_official.(bool) { ... } のようにすればよい
	IsOfficial      interface{} `json:"is_official"`
	RoomURLKey      string      `json:"room_url_key"`
	RoomDescription string      `json:"room_description"`
	ImageS          string      `json:"image_s"`
	IsOnline        interface{} `json:"is_online"`
	IsFav           bool        `json:"is_fav"`
	GenreID         int         `json:"genre_id"`
	Point           int         `json:"point"`
	RoomID          string      `json:"room_id"`
	Rank            int         `json:"rank"`
	RoomName        string      `json:"room_name"`
}

// ブロックランキングイベント参加中のルーム情報の一覧を取得する。
func GetEventBlockRanking(
	client *http.Client,
	eventid int,
	blockid int,
	ib int,
	ie int,
) (
	ebr *EventBlockRanking,
	err error,
) {

	ebr = new(EventBlockRanking)
	ebr.BlockRankingList = make([]BlockRankingList, 0)

	noroom := 0

	for page := 1; page > 0; {
		tebr := new(EventBlockRanking)
		//	tebr.Block_ranking_list = make([]Block_ranking, 0)

		//	イベント参加ルーム一覧のデータ（htmmの一部をぬきだした形になっている）を取得する。
		//	データを分割して取得するタイプのAPIを使うときはこのような処理を入れておいた方が安全です。
		tebr, err = ApiBlockEventRnaking(client, eventid, blockid, page)
		if err != nil {
			err = fmt.Errorf("ApiEventRoomList(): %w", err)
			return nil, err
		}
		if len(tebr.BlockRankingList) == 0 {
			// 理屈としては起こり得ないが、現実には起きたことがある。
			break
		}

		if page == 1 {
			ebr = tebr
		} else {
			ebr.BlockRankingList = append(
				ebr.BlockRankingList,
				tebr.BlockRankingList...,
			)
		}

		noroom = len(ebr.BlockRankingList)

		if noroom == ebr.TotalEntries || noroom >= ie {
			break
		}

		//	次のページへ
		page++
	}

	if ib <= noroom {
		if ie > noroom {
			ie = noroom
		}
		ebr.BlockRankingList = ebr.BlockRankingList[ib-1 : ie]
	} else {
		ebr.BlockRankingList = nil
		return
	}

	return
}

func ApiBlockEventRnaking(
	client *http.Client, //	HTTPクライアント
	event_id int,
	block_id int,
	page int,
) (
	ebr *EventBlockRanking,
	err error,
) {

	turl := "https://www.showroom-live.com/api/event/block_ranking"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{}                              // url.Valuesオブジェクト生成
	values.Add("event_id", fmt.Sprintf("%d", event_id)) // key-valueを追加
	values.Add("block_id", fmt.Sprintf("%d", block_id)) // key-valueを追加
	values.Add("page", fmt.Sprintf("%d", page))         // key-valueを追加

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

	if err = json.NewDecoder(resp.Body).Decode(&ebr); err != nil {
		err = fmt.Errorf("json.NewDecoder(resp.Body).Decode(&ebr): %w", err)
		return
	}

	//	現状では下記でもOK
	//	if block_id == 0 {
	//		for i, br := range ebr.Block_ranking_list {
	//			ebr.Block_ranking_list[i].Is_official = br.Is_online != 0
	//			ebr.Block_ranking_list[i].Is_online = br.Is_online != 0
	//		}
	//	}
	//	今後のSHOWROOMの仕様変更の可能性を考え以下のようにしておく
	//	最悪(?)のケースを考えるとこれでも不十分。
	for i, br := range ebr.BlockRankingList {
		if vi, ok := br.IsOnline.(float64); ok {
			ebr.BlockRankingList[i].IsOnline = vi != 0
		} else {
			break
		}
		if vi, ok := br.IsOfficial.(float64); ok {
			ebr.BlockRankingList[i].IsOfficial = vi != 0
		}
	}

	return
}
