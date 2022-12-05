
//		** 作成中　**
//		** インターフェースを変更予定 **

/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

Ver. 0.0.0

*/

package srapi

import (
	"fmt"
	"strings"

	"encoding/json"
	"net/http"
	"net/url"
)

type LiveGiftingFree struct {
	Ua              int      // ?
	Aft             int      // ?
	Num             int      // ?
	Remaining_num   int      // 星、種の残り個数
	Gifting_bonus   int      // ?
	Ok              bool     //	true: 正常終了？
	Notify_level_up bool     //	true: レベルアップ通知？
	Gift_id         int      //	(Official) 1, 2, 1001, 1002, 1003, (Free) 1501, 1502, 1503, 1504. 1505
	Add_point       int      //	ボーナスを含んだ獲得ポイント
	Gift_name       string   //	ギフト名 （Offiicial） "星A", "星B", "星C", "星D", "星E",（Free） "種A", "種B", "種C", "種D", "種E"
	Bonus_rate      float64  //	ボーナス率（num >= 10であれば0.2）
	Fan_Level       struct { // ファンレベル（いわゆるリスナーレベル？）
		Contribution_point int //	累積貢献ポイント ？
	}
}

//	星を投げる/種を投げる。
func ApiLiveGiftingFree(
	client *http.Client, //	HTTPクライアント
	csrftoken string, //	接続の識別子
	liveid string, //	配信枠の識別子
	giftid string, //	(Official) "1", "2", "1001", "1002", "1003", (Free) "1501", "1502", "1503", "1504". "1505"
	num string, //	投げる星、種の個数
) (
	livegiftingfree *LiveGiftingFree,
	err error,
) {

	err = nil

	turl := "https://www.showroom-live.com/api/live/gifting_free"
	u, e := url.Parse(turl)
	if e != nil {
		err = fmt.Errorf("url.Parse(): %w", e)
		return
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成
	//	values.Add([クエリのキー], [値]) // key-valueを追加
	values.Add("gift_id", giftid)
	values.Add("live_id", liveid)
	values.Add("num", num)
	values.Add("csrf_token", csrftoken)

	// Request を生成
	req, e := http.NewRequest("POST", u.String(), strings.NewReader(values.Encode()))
	if e != nil {
		err = fmt.Errorf("http.NewRequest(): %w", e)
		return
	}

	// Content-Typeを設定
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// User-Agentを設定
	req.Header.Add("User-Agent", useragent)
	req.Header.Add("Accept-Language", "ja-JP")

	//	req.SetBasicAuth([USER], [PASS})

	//	log.Printf("****** req=%+v\n", req)

	// Doメソッドでリクエストを投げる
	// http.Response型のポインタ（とerror）が返ってくる
	resp, e := client.Do(req)
	if e != nil {
		err = fmt.Errorf("client.Do(): %w", e)
		return
	}

	// 関数を抜ける際に接続を切断し、リソースを解放するため必ずresponse.Bodyをcloseする
	defer resp.Body.Close()

	if e = json.NewDecoder(resp.Body).Decode(&livegiftingfree); e != nil {
		err = fmt.Errorf("json.NewDecoder(resp.Body).Decode(): %w", e)
		return
	}

	return

}
