/*
!
Copyright © 2024 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php
*/
package srapi

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"time"

	"encoding/json"

	"net/http"
	"net/url"
)

type GenreRanking struct {
	BeforeRank int `json:"before_rank"`
	OrderNo    int `json:"order_no"`
	Point      int `json:"point"`
	Room       struct {
		RoomURLKey  string `json:"room_url_key"`
		ImageSquare string `json:"image_square"`
		RoomName    string `json:"room_name"`
		RoomID      int    `json:"room_id"`
		Image       string `json:"image"`
	} `json:"room"`
	Rank int `json:"rank"`
}


type GenreRankingArray struct {
	NextPage     int `json:"next_page"`
	GenreRanking []GenreRanking `json:"genre_ranking"`
	TotalCount int `json:"total_count"`
}

/*
	func ApiGenre_ranking() ランキング一覧を取得する。

	APIの使用例
	https://www.showroom-live.com/api/genre_ranking/{genreid}/{period}?page=1&count=20

genreid:
0       全て
101     アーティスト
102     アイドル
103     タレント
104     声優
105     芸人
107     バーチャル
108     モデル
109     俳優
110     アナウンサー
111     グローバル
200     ライバー

period:
all_time	総合
annually/20240101
monthly/20240601
weekly/20240603  ... 月曜日
daily/20240605

*/
func ApiGenre_ranking(
	client *http.Client,
	genreid int,
	period string,
	iscurrent bool,
	page int,
	count int,
) (
	pgr *GenreRankingArray,
	err error,
) {

	turl := "https://www.showroom-live.com/api/genre_ranking/" + strconv.Itoa(genreid) + "/"

	tnow := time.Now()
	_, mm, dd := tnow.Date()
	switch period {
	case "daily":
		if !iscurrent {
			//	一日前のデータを取得する
			tnow = tnow.AddDate(0, 0, -1)
		}
		period += "/" + tnow.Format("20060102")
	case "weekly":
		//	週の最初を月曜日とし、今日（tnow）を含む週の最初の日を求める
		wd := int(tnow.Weekday())
		if wd == 0 {
			wd = 7
		}
		tnow = tnow.AddDate(0, 0, -wd+1)
		if !iscurrent {
		//	先週のデータを取得する場合
			tnow = tnow.AddDate(0, 0, -7)
		}
		period += "/" + tnow.Format("20060102")
	case "monthly":
		//	月の最初を1日とし、今日（tnow）を含む月の最初の日を求める
		tnow = tnow.AddDate(0, 0, -dd+1)
		if !iscurrent {
		//	先月のデータを取得する場合（2024年7月から利用可能になると思われう）
			tnow = tnow.AddDate(0, -1, 0)
		}	
		period += "/" + tnow.Format("20060102")
	case "annually":
		//	年の最初を1日とし、今日（tnow）を含む年の最初の日を求める
		tnow = tnow.AddDate(0, -int(mm)+1, 0)
		if !iscurrent {
		//	前年のデータを取得する場合	(2025年から利用可能になると思われる)
			tnow = tnow.AddDate(-1, 0, 0)
		}
		period += "/" + tnow.Format("2006") + "0101"
	case "all_time":
	default:
		err = fmt.Errorf("period is not defined <%s>", period)
		return
	}

	turl += period

	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{}         // url.Valuesオブジェクト生成
	values.Add("page", strconv.Itoa(page)) // key-valueを追加
	values.Add("count", strconv.Itoa(count)) // key-valueを追加

	// Request を生成
	var req *http.Request
	req, err = http.NewRequest("GET", u.String(), nil)
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
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		err = fmt.Errorf("client.Do(): %w", err)
		return
	}
	// 関数を抜ける際に接続を切断し、リソースを解放するため必ずresponse.Bodyをcloseする
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bufstr := buf.String()

	//	JSONをデコードする。
	//	次の記事を参考にさせていただいております。
	//		Go言語でJSONに泣かないためのコーディングパターン
	//		https://qiita.com/msh5/items/dc524e38073ed8e3831b

	pgr = new(GenreRankingArray)
	decoder := json.NewDecoder(buf)
	if err = decoder.Decode(pgr); err != nil {
		log.Printf("decoder.Decode(&result) err: %v", err)
		log.Printf("genreid= %d", genreid)
		log.Printf("period= %s", period)
		log.Printf("page= %d", page)
		log.Printf("count= %d", count)
		log.Printf("bufstr: %s", bufstr)
		err = fmt.Errorf("decoder.Decode(&result) err: %v", err)
		return
	}

	return
}

func GetGenreRankingByApi(
	client *http.Client,
	genreid int,
	period string,
	iscurrent bool,
	pages int,
) (
	rlist *[]GenreRanking,
	err error,
) {

	rlist = &[]GenreRanking{}

	for  p := 1; p <= pages; p++ {
		var pgr *GenreRankingArray
		pgr, err = ApiGenre_ranking(client, genreid, period, iscurrent, p, 20)
		if err != nil {
			err = fmt.Errorf("ApiGenre_ranking(): %w", err)
			return
		}

		*rlist = append(*rlist, pgr.GenreRanking...)
	}

	return
}
