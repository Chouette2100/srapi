/*!
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

type Block_ranking struct {
	Is_official      bool
	Room_url_key     string
	Room_description string
	Image_s          string
	Is_online        bool
	Is_fav           bool
	Genre_id         int
	Point            int
	Room_id          string
	Rank             int
	Room_name        string
}

type EventBlockRanking struct {
	Total_entries      int
	Entries_per_pages  int
	Current_page       int
	Block_ranking_list []Block_ranking
}

//	ブロックランキングイベント参加中のルーム情報の一覧を取得する。
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
	ebr.Block_ranking_list = make([]Block_ranking, 0)

	noroom := 0

	for page := 1; page > 0; {

		//	イベント参加ルーム一覧のデータ（htmmの一部をぬきだした形になっている）を取得する。
		//	データを分割して取得するタイプのAPIを使うときはこのような処理を入れておいた方が安全です。
		ebr, err = ApiBlockEventRnaking(client, eventid, blockid, page)
		if err != nil {
			err = fmt.Errorf("ApiEventRoomList(): %w", err)
			return nil, err
		}

		noroom = len(ebr.Block_ranking_list)

		if noroom == ebr.Total_entries || noroom >= ie {
			break
		}

		//	次のページへ
		page++
	}

	if ib <= noroom {
		if ie > noroom {
			ie = noroom
		}
		ebr.Block_ranking_list = ebr.Block_ranking_list[ib-1 : ie]
	} else {
		ebr.Block_ranking_list = nil
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
	return
}
