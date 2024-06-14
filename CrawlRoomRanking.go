/*
!
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
	"sort"
	"strconv"
	"strings"

	"net/http"
	//	"net/url"

	"github.com/PuerkitoBio/goquery"
	//	"encoding/json"
)

type RoomRanking struct {
	GenreID int
	Period  string
	Order   int
	Room_id int
}

type Genre struct {
	Genre   string
	GenreID int
}

var GenreTbl []Genre = []Genre{
	{"総合", 999},
	{"フリー", 200},
	{"タレント・モデル", 103},
	{"ミュージック", 101},
	{"声優・アニメ", 104},
	{"お笑い・トーク", 105},
	{"バーチャル", 107},
	{"新人", 998},
}

/*
var PeriodTbl []string = []string{
	"all",
	"daily",
	"weekly",
	"monthly",
}
*/

var PrdMap map[string]int = map[string]int{
	"all":     0,
	"daily":   1,
	"weekly":  2,
	"monthly": 3,
}

/*
	CrawlRoomRanking()
	
*/
func CrawlRoomRanking(
	mode string,
) (
	rr *[]RoomRanking,
	err error,
) {

	webPage := ("https://www.showroom-live.com/ranking")
	resp, err := http.Get(webPage)
	if err != nil {
		err = fmt.Errorf("http.Get(): %w", err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		err = fmt.Errorf("goquery.NewDocument(): %w", err)
		return
	}

	rmrnk := []RoomRanking{}

	/*
		for i:= 0; i < 30; i++ {
			selector := doc.Find("#js-ranking-section-all-0 > ul > li:nth-child(" + strconv.Itoa(i+1) + ") .listcardinfo-menu > ul > li")
			ridurl, exist := selector.Find("a").Attr("href")
			if !exist {
				err = errors.New("selector.Find(\"a\").Attr(\"href\")")
				return
			}
			rid := strings.Split(ridurl, "=")[1]
			ridint := 0
			ridint, err = strconv.Atoi(rid)
			if err != nil {
				err = fmt.Errorf("strconv.Atoi(): %w", err)
				return
			}

			rmrnk = append(rmrnk, RoomRanking{
				GenreID: GenreTbl[0].GenreID,
				Period: "All",
				Order: i+1,
				Room_id: ridint,
			})
		}
	*/

	for k := 0; k < len(GenreTbl)-1; k++ { //	"新人"は除く
		//	for k := 4; k < len(GenreTbl)-1; k++ {
		log.Printf(" k=%d\n", k)

		//	for j := 0; j < len(PeriodTbl); j++ {
		j, exist := PrdMap[mode]
		if !exist {
			err = errors.New("mode <" + mode + "> is not found")
			return
		}

		log.Printf(" j=%d\n", j)

		//	#js-ranking-section-all-0 > ul > li:nth-child(1) > div > div.listcardinfo > div.listcardinfo-menu > ul > li:nth-child(1) > a
		//	#js-ranking-section-all-0 > ul > li:nth-child(1) .listcardinfo-menu > ul > li:nth-child(1) > a
		selector := doc.Find("#js-ranking-section-" + mode + "-" + strconv.Itoa(k) + " > ul > li")
		selector.Each(func(i int, s *goquery.Selection) { //	条件で途中でやめる場合は EachWithBreak を使い、return falseでブレイクする
			ridurl, exist := s.Find(" .listcardinfo-menu > ul > li:nth-child(1) > a").Attr("href")
			if !exist {
				err = errors.New("selector.Find(\"a\").Attr(\"href\")")
				return
			}
			rid := strings.Split(ridurl, "=")[1]
			ridint := 0
			ridint, err = strconv.Atoi(rid)
			if err != nil {
				err = fmt.Errorf("strconv.Atoi(): %w", err)
				return
			}

			rmrnk = append(rmrnk, RoomRanking{
				GenreID: GenreTbl[k].GenreID,
				Period:  mode,
				Order:   i + 1,
				Room_id: ridint,
			})
		})
		log.Printf("size of mrnk=%d\n", len(rmrnk))

	}

	sort.SliceStable(rmrnk, func(i, j int) bool { return rmrnk[i].Room_id < rmrnk[j].Room_id })
	log.Printf("size of mrnk=%d\n", len(rmrnk))

	rmrnksu := []RoomRanking{}
	lastid := 0
	for i := 0; i < len(rmrnk); i++ {
		//	重複したルームIDを除く
		if rmrnk[i].Room_id != lastid {
			rmrnksu = append(rmrnksu, rmrnk[i])
			lastid = rmrnk[i].Room_id
		}
	}
	log.Printf("size of mrnksu=%d\n", len(rmrnksu))

	rr = &rmrnksu

	return

}
