/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

*/

package srapi

import (
	"fmt"
	"net/http"
	"net/url"

	"encoding/json"
)

/*

Ver. 0.1.0

*/

type Event struct {
	Content_region_permission int
	Order_no                  int
	Event_id                  int //	イベント識別子
	Tag_list                  []struct {
		Tag_name string
	}
	Event_description     string
	Is_event_Block        bool
	Is_entry_scope_inner  bool	//	対象者限定か否か（？）
	Start_soon            bool
	League_ids            []int	//	val.= 9:XX, 10:S, 20:A, 30:B, 40:C （？）
	End_soon              bool
	Type_name             string
	Is_official           bool
	Event_name            string //	イベント名
	Image_l               string
	Event_type            string
	Image_s               string
	Genre_id              int
	Default_locale        string
	Ended_at              int64 //	イベント終了日時(UnixTime)
	Updated_at            int64
	Start_comming         interface{}
	Event_block_label     interface{}
	Event_image_file_type string
	Required_level_max    int	//	参加可能レベル上限
	Required_level        int	//	参加可能レベル下限
	Image_m               string
	Offer_started_at      int64
	No_amateur            bool
	End_remain            interface{}
	Just_offer_started    interface{}
	Created_at            int64
	Started_at            int64  //	イベント開始日時(UnixTime)
	Event_url_key         string //	イベントURL(最後のフィールド)
	Is_box_event          bool	//	親イベントか否か（？）
	Parent_event_id       int
	Show_ranking          bool
	Is_watch              bool
	Is_public             bool
	Public_status         int
	Is_closed             bool
	Is_pickup             bool
	Reward_rank           int
	Offer_ended_at        int64 //	申込み終了日時(UnixTime)
}

type EventSearch struct {
	Next_page     int
	Previous_page int
	Last_page     int //	ページ数、この数だけ取得を繰り返す。
	First_page    int
	Current_page  int
	Total_count   int //	イベント数、イベント情報取得後数が一致しているかチェックする。
	Event_list    []Event
}

//	複数ページにわかれたイベント情報を取得して一つのスライスにする。
func MakeEventListByApi(client *http.Client) (esl []Event, err error) {

	esl = make([]Event, 0, 200)

	for i := 1; ; i++ {
		es, err := ApiEventSearch(client, i)
		if err != nil {
			err = fmt.Errorf("ApiEventSearch error(): %w", err)
			return nil, err
		}
		esl = append(esl, es.Event_list...)

		if es.Current_page == es.Last_page {
			break
		}
	}

	return
}

//	配信中のルームの一覧を取得する
func ApiEventSearch(
	client *http.Client, //	HTTP client
	page int,
) (
	es *EventSearch, //	配信中ルームのジャンル別一覧
	err error, //	エラー
) {

	turl := "https://www.showroom-live.com/api/event/search"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse: %w", err)
		return es, err
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成
	values.Add("page", fmt.Sprintf("%d", page))

	//	log.Printf("values=%+v\n", values)

	// Request を生成
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		err = fmt.Errorf("http.NewRequest: %w", err)
		return es, err
	}

	// 組み立てたクエリを生クエリ文字列に変換して設定
	req.URL.RawQuery = values.Encode()

	// User-Agentを設定
	req.Header.Add("User-Agent", useragent)

	// Doメソッドでリクエストを投げる
	// http.Response型のポインタ（とerror）が返ってくる
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("client.Do: %w", err)
		return es, err
	}

	// 関数を抜ける際に接続を切断し、リソースを解放するため必ずresponse.Bodyをcloseする
	defer resp.Body.Close()

	es = &EventSearch{}
	es.Event_list = make([]Event, 0, 200)

	if err = json.NewDecoder(resp.Body).Decode(es); err != nil {
		err = fmt.Errorf("json.NewDecoder: %w", err)
		return es, err
	}

	return es, nil

}
