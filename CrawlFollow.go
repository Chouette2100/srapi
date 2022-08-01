/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

Ver. 0.0.0

*/
package srapi

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
)

type RoomFollowing struct {
	Room_id      string //	ルームID	配信者を識別する
	Room_url_key string //	配信のURLの最後のフィールド
	Main_name    string //	ルーム名
	Next_live    string //	次の配信時刻
}

func CrwlFollow(
	client *http.Client,
	maxnoroom 	int,
) (
	rooms *[]RoomFollowing,
	status int,
) {

	var doc *goquery.Document

	status = 0

	turl := "https://www.showroom-live.com/follow"
	u, err := url.Parse(turl)
	if err != nil {
		log.Printf("url.Parse() returned error %s\n", err.Error())
		return
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成
	//	values.Add([クエリのキー], [値]) // key-valueを追加

	// Request を生成
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Printf("http.NewRequst() returned error %s\n", err.Error())
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
		log.Printf("client.Do() returned error %s\n", err.Error())
		return
	}
	defer resp.Body.Close()

	rooms = &[]RoomFollowing{}

	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("GetRoomsFollowing() goquery.NewDocumentFromReader() err=<%s>.\n", err.Error())
		status = 1
		return
	}
	defer resp.Body.Close()

	//	抽出したルームすべてに対して処理を繰り返す
	doc.Find(".listcardinfo").EachWithBreak(func(i int, s *goquery.Selection) bool {

		var room RoomFollowing

		room.Main_name = s.Find(".listcardinfo .listcardinfo-main-text").Text()
		room.Room_url_key, _ = s.Find(".listcardinfo a").Attr("href")
		room.Room_id, _ = s.Find(".listcardinfo a").Attr("data-room-id")
		room.Next_live = s.Find(".listcardinfo .is-nextlive").Text()

		//	log.Printf("%+v\n", room)

		*rooms = append(*rooms, room)

		i++
		return i < maxnoroom 

	})

	return
}
