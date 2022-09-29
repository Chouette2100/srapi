/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

Ver. 0.1.0

*/
package srapi

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"

	"encoding/json"
)

type Room struct {
	Room_id      int
	Room_url_key string
	Room_name    string
	Point        int //	これは /event/room_list の戻り値には含まれない。以下同様
	Gap          int
	Rank         int
	Islive       bool
	Isofficial   bool
	Startedat    int64
	Nextlive     int64
}

type RoomListInf struct {
	RoomList []Room
}

//	イベント参加中のルーム情報の一覧を取得する。
func GetRoominfFromEventByApi(
	client *http.Client,
	eventid int,
	ib int,
	ie int,
) (
	roomlistinf *RoomListInf,
	err error,
) {

	roomlistinf = new(RoomListInf)

	next_page := 0

	roomlistinf.RoomList = make([]Room, 0)

	//	ルーム一覧にルームの重複がないかチェックするためのマップ。
	usernomap := make(map[int]bool, 500)

	it := 0
	for ip := 1; ip > 0; {

		//	イベント参加ルーム一覧のデータ（htmmの一部をぬきだした形になっている）を取得する。
		//	データを分割して取得するタイプのAPIを使うときはこのような処理を入れておいた方が安全です。
		eventroomlist, err := ApiEventRoomList(client, eventid, ip)
		if err != nil {
			err = fmt.Errorf("ApiEventRoomList(): %w", err)
			return nil, err
		}

		//	eventroomlist.Html はhtmlのbodyの中身なのでhtmlの形式にする。
		valuehtml := "<html>" + "<body>" + eventroomlist.Html + "</body>" + "</html>"

		//	html のクロールを行う。
		err = CrawlEventRoomList(valuehtml, &roomlistinf.RoomList, &usernomap, ib, ie, &it)
		if err != nil {
			err = fmt.Errorf("CrawlEventRoomList(): %w", err)
			return nil, err
		}
		noroom := len(roomlistinf.RoomList)
		log.Printf(" ip=%d next_page=%d len=%d\n", ip, next_page, noroom)
		if noroom >= ie-ib+1 {
			return roomlistinf, nil
		}
		ip = eventroomlist.Next_page
	}

	return
}

func CrawlEventRoomList(valuehtml string,
	roomlist *[]Room,
	usernomap *map[int]bool,
	ib int,
	ie int,
	it *int,
) (
	err error,
) {

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(valuehtml))
	if err != nil {
		err = fmt.Errorf("goquery.NewDocumentFromReader(): %w", err)
		return
	}

	//	抽出したルームすべてに対して処理を繰り返す(が、イベント開始後の場合の処理はルーム数がbreg、eregの範囲に限定）
	//	イベント開始前のときはすべて取得し、ソートしたあてで範囲を限定する）
	doc.Find("body > li").EachWithBreak(func(i int, s *goquery.Selection) bool {
		//	log.Printf("i=%d\n", i)

		var room Room

		room.Room_name = s.Find(".listcardinfo-main-text").Text()

		//	----------------------------------------------
		//	pointの取得は イベント最終日の翌日に /event/{eventname} の結果をクロールするときだけ意味があります。
		//	イベント最終日翌日であってもポイントを取得できるのはランキングイベントで上位の30ルームについてだけです。
		spoint1 := strings.Split(s.Find(".listcardinfo-sub-single-right-text").Text(), ": ")

		var point int64
		if spoint1[0] != "" {
			spoint2 := strings.Split(spoint1[1], "pt")
			fmt.Sscanf(spoint2[0], "%d", &point)

		} else {
			point = -1
		}
		room.Point = int(point)
		//	----------------------------------------------

		selection_c := s.Find(".listcardinfo-menu")

		account, exist := selection_c.Find(".room-url").Attr("href")
		if !exist {
			err = errors.New(`selection_c.Find(".room-url").Attr("href")' retruned false`)
			return false
		}

		account_a := strings.Split(account, "/")
		room.Room_url_key = account_a[len(account_a)-1]

		sri, exist := selection_c.Find(".js-follow-btn").Attr("data-room-id")
		if !exist {
			err = errors.New(`selection_c.Find(".js-follow-btn").Attr("data-room-id") returned false`)
			return false
		}
		room.Room_id, _ = strconv.Atoi(sri)

		_, ok := (*usernomap)[room.Room_id]
		if !ok && room.Room_id != 0 {
			//	ルームに重複がない
			(*usernomap)[room.Room_id] = true
			*it++
			if *it >= ib {
				*roomlist = append(*roomlist, room)
				if *it == ie {
					//	必要なルーム数のデータが得られた。
					return false
				}
			}
		}

		return true

	})

	log.Printf(" CrawlEventInfAndRoomList() len(*roominfolist)=%d\n", len(*roomlist))

	return

}

type EventRoomList struct {
	Next_page int
	Html      string
}

//	配信状況を確認し、room_url_key から　room_id を取得する。
func ApiEventRoomList(
	client *http.Client, //	HTTPクライアント
	eventid int,
	page int, //	ルームのURLの最後のフィールド
) (
	eventroomlist *EventRoomList,
	err error,
) {

	turl := "https://www.showroom-live.com/event/room_list"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{}                             // url.Valuesオブジェクト生成
	values.Add("event_id", fmt.Sprintf("%d", eventid)) // key-valueを追加
	values.Add("p", fmt.Sprintf("%d", page))           // key-valueを追加

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

	eventroomlist = new(EventRoomList)
	if err = json.NewDecoder(resp.Body).Decode(eventroomlist); err != nil {
		err = fmt.Errorf("json.NewDecoder(resp.Body).Decode(&roomstatus): %w", err)
		return
	}
	return
}
