package srapi

import (
	// "errors"
	"fmt"
	"log"
	// "sort"
	"strconv"
	"strings"

	"net/http"
	//	"net/url"

	"github.com/PuerkitoBio/goquery"
	//	"encoding/json"
)

type CrEventRank struct {
	Order    int
	Rank     int    //      貢献順位
	Listner  string //      リスナー名
	Lastname string //      前配信枠でのリスナー名

	LsnID       int //      リスナーのユーザID（Ver.3.0A00より前のバージョンではAPIで取得できなかったため0がセットされている）
	T_LsnID     int //      Ver.3.0A00より前のバージョンで用いたリスナー識別のための（仮の）ユーザーID（>イベントごとに異なる）
	Point       int //      貢献ポイント
	Incremental int //      貢献ポイントの増分（＝配信枠別貢献ポイント）
	Status      int
}

// 構造体のスライス
type CrEventRanking []CrEventRank

// sort.Sort()のための関数三つ
func (e CrEventRanking) Len() int {
	return len(e)
}

func (e CrEventRanking) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// 降順に並べる
func (e CrEventRanking) Less(i, j int) bool {
	//      return e[i].point < e[j].point
	return e[i].Point > e[j].Point
}

type CrContributionRanking struct {
	TotalScore   int
	BonusPoints  int
	EventRanking CrEventRanking
}

// イベント貢献ランキングをクロールして取得します（2025-10-20の時点ではこの方法は使えなさそう、JavaScriptが必要？、chromedpを試してみる）
func CrawlContirbutionRanking(EventName string, roomno int) (
	crCR *CrContributionRanking,
	err error,
) {

	crCR = &CrContributionRanking{}

	roomid := strconv.Itoa(roomno)

	//	貢献ランキングのページを開き、データ取得の準備をします。
	//	_url := "https://www.showroom-live.com/event/contribution/" + EventName + "?room_id=" + roomid
	ename := EventName
	ename_a := strings.Split(EventName, "?")
	if len(ename_a) == 2 {
		ename = ename_a[0]
	}
	_url := "https://www.showroom-live.com/event/contribution/" + ename + "?room_id=" + roomid

	resp, error := http.Get(_url)
	if error != nil {
		log.Printf("GetEventInfAndRoomList() http.Get() err=%s\n", error.Error())
		err = fmt.Errorf("failed to get event info: %w", error)
		return
	}
	defer resp.Body.Close()

	var doc *goquery.Document
	doc, error = goquery.NewDocumentFromReader(resp.Body)
	if error != nil {
		log.Printf("GetEventInfAndRoomList() goquery.NewDocumentFromReader() err=<%s>.\n", error.Error())
		err = fmt.Errorf("failed to parse event page: %w", error)
		return
	}

	/*
		u := url.URL{}
		u.Scheme = doc.Url.Scheme
		u.Host = doc.Url.Host
	*/

	/* test code
	bns1 := doc.Find("section.p-b4:nth-child(3) > table:nth-child(2) > tbody:nth-child(2) > tr:nth-child(2) > td:nth-child(3)").Text()
	bns2 := doc.Find("section.p-b4:nth-child(3) > table:nth-child(2) > tbody:nth-child(2) > tr:nth-child(3) > td:nth-child(3)").Text()

	fmt.Printf("BNS1=%s\n", bns1)
	fmt.Printf("BNS2=%s\n", bns2)
	*/

	// doc.Find("section.p-b4:nth-child(3) > table:nth-child(2) > tbody:nth-child(2) > tr").Each(func(i int, s *goquery.Selection) {
	// 	if i != 0 {
	// 		bns := s.Find("td:nth-child(3)").Text()

	// 	}
	// })

	//	各リスナーの情報を取得します。
	//	var selector_ranking, selector_listner, selector_point, ranking, listner, point string
	var ranking, listner, point string
	var iranking, ipoint int
	var eventrank CrEventRank

	crCR.TotalScore = 0
	crCR.BonusPoints = 0

	//	eventranking = make([]EventRank)

	doc.Find(".table-type-01:nth-child(2) > tbody > tr").Each(func(i int, s *goquery.Selection) {
		if i != 0 {

			//	データを一つ取得するたびに(戻り値となる)リスナー数をカウントアップします。
			//	NoListner++

			//	以下セレクターはブラウザの開発ツールを使って確認したものです。

			//	順位を取得し、文字列から数値に変換します。
			//	selector_ranking = fmt.Sprintf("table.table-type-01:nth-child(2) > tbody:nth-child(2) > tr:nth-child(%d) > td:nth-child(%d)", NoListner+2, 1)
			ranking = s.Find("td:nth-child(1)").Text()

			/*
				//	データがなくなったらbreakします。このときのNoListnerは通常100、場合によってはそれ以下です。
				if ranking == "" {
					break
				}
			*/

			iranking, _ = strconv.Atoi(ranking)

			//	リスナー名を取得します。
			//	selector_listner = fmt.Sprintf("table.table-type-01:nth-child(2) > tbody:nth-child(2) > tr:nth-child(%d) > td:nth-child(%d)", NoListner+2, 2)
			listner = s.Find("td:nth-child(2)").Text()

			//	貢献ポイントを取得し、文字列から"pt"の部分を除いた上で数値に変換します。
			//	selector_point = fmt.Sprintf("table.table-type-01:nth-child(2) > tbody:nth-child(2) > tr:nth-child(%d) > td:nth-child(%d)", NoListner+2, 3)
			point = s.Find("td:nth-child(3)").Text()
			point = strings.Replace(point, "pt", "", -1)
			ipoint, _ = strconv.Atoi(point)
			crCR.TotalScore += ipoint
			if iranking == 0 {
				crCR.BonusPoints += ipoint
			}

			//	戻り値となるスライスに取得したデータを追加します。
			eventrank.Rank = iranking
			eventrank.Point = ipoint
			eventrank.Listner = listner
			eventrank.Order = i
			crCR.EventRanking = append(crCR.EventRanking, eventrank)
		}
	})

	return
}
