/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

Ver. 0.0.0
Ver. 0.0.1 s.Find()のセレクタの冗長部分を削除する。

*/
package srapi

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

//	フォローしているルームのデータの構造体（必要とするものだけ、残りは取得していません）
type RoomFollowing struct {
	Room_id      string //	ルームID	配信者を識別する
	Room_url_key string //	配信のURLの最後のフィールド
	Main_name    string //	ルーム名
	Next_live    string //	次の配信時刻
}

//	フォローしているルームの一覧を取得する。
func CrwlFollow(
	client *http.Client,
	maxnoroom int, //	取得するルーム数（99999とかしていれば全部取得）
) (
	rooms *[]RoomFollowing,
	err error,
) {

	var doc *goquery.Document

	turl := "https://www.showroom-live.com/follow"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse: %w", err)
		return nil, err
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成
	//	values.Add([クエリのキー], [値]) // key-valueを追加

	// Request を生成
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		err = fmt.Errorf("http.NewRequest: %w", err)
		return nil, err
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
		return nil, err
	}
	defer resp.Body.Close()

	rooms = &[]RoomFollowing{}

	//	APIと違って（JSONではなく）単にHTMLを返してくるので、goqueryでパースします。
	doc, err = goquery.NewDocumentFromReader(resp.Body) //	NewDocument()を使うのは現在非推奨になっています。
	if err != nil {
		err = fmt.Errorf("NewDocumentFromReade: %w", err)
		return nil, err
	}
	defer resp.Body.Close()

	//	フォローするルームmaxnoroom分に対して処理を繰り返す
	doc.Find(".listcardinfo").EachWithBreak(func(i int, s *goquery.Selection) bool {

		var room RoomFollowing

		//	ここからはほしいデータがある場所を見つけて、そこのCSSセレクタを指定してデータを取得します。
		//	CSSセレクタはGoogle ChromeのDeveloper Toolsで調べられます。
		//	ただこの方法で得られるCSSセレクタは冗長になりがちなので、htmlながめながら自分で書いた方がいいかも(好みに合わせて)
		//	自分が作ったセレクタで狙ったところを指定できているかもDevelopper Toolsで確認できます(検索窓はctrl-Fで開きます）
		//
		//	ちなみに下記の
		//		".listcardinfo .listcardinfo-main-text"
		//	については Google ChromeのDeveloper Toolsで Copy Selector をやると
		//		"#js-genre-section-all > ul > li:nth-child(1) > div > div > div.listcardinfo-info > h4"
		//	となります。こちらを使うのであれば、
		//
		//	doc.Find("#js-genre-section-all > ul > li").EachWithBreak(func(i int, s *goquery.Selection) bool {
		//		.....
		//		room.Main_name = s.Find("div > div > div.listcardinfo-info > h4").Text()
		//		.....
		//	}
		//
		//	というような書き方になります。
		//
		//	他にもセレクターを
		//			"#js-genre-section-all > ul > li:nth-child(" + fmt.Sprintf("%d", i) + ") > div > div > div.listcardinfo-info > h4"
		//	としてfor文でぐるぐる回すという方法もあります。
		//
		//	たぶんいちばんいい（適切な）方法というのがあると思うのですが、あんまり考えたことないです。すみません。
		//
		//	詳しいことは
		//		https://pkg.go.dev/github.com/PuerkitoBio/goquery
		//		https://developer.mozilla.org/ja/docs/Web/CSS/CSS_Selectors
		//	を参考にしてください。

		room.Main_name = s.Find(".listcardinfo-main-text").Text()
		room.Room_url_key, _ = s.Find("a").Attr("href")
		room.Room_id, _ = s.Find("a").Attr("data-room-id")
		room.Next_live = s.Find(".is-nextlive").Text()

		//	log.Printf("%+v\n", room)

		*rooms = append(*rooms, room)

		i++
		return i < maxnoroom //   EachWithBreak() は do while のようなものです。while 条件 に相当するのが return 条件 です。

	})

	return rooms, nil
}
