

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

//	星・種の一気投げの結果
type BulkGiftingFree struct {
	Ua              int         //	3: ?
	Aft             int         //	0: ?
	Gifting_bonus   interface{} //	null: ?　現時点では型が不明
	Ok              bool        //	true: 正常終了？
	Notify_level_up bool        //	false: ?
	Gifts           []struct {
		Gift_id       int     //	(Official) 1, 2, 1001, 1002, 1003, (Free) 1501, 1502, 1503, 1504. 1505
		Add_point     int     //	ボーナスを含んだ獲得ポイント
		num           int     //	投げた星、種の個数
		Remaining_num int     //	残りの星、種の個数（0になっているはず
		Gift_name     string  //	ギフト名 （Offiicial） "星A", "星B", "星C", "星D", "星E",（Free）"種A", "種B", "種C", "種D", "種E"
		Bonus_rate    float64 //	ボーナス率（num >= 10であれば0.2）
		Gift_type     int     //	2: ?
	}
	Fan_level struct {
		Next_level_point    int //	未確認
		Current_level_point int //	未確認
		Next_fan_level      int //	未確認
		Contribution_point  int //	累積貢献ポイント ？
		User_id             int //	リスナーの識別子
		Fan_level           int //	いわゆるリスナーレベル（現時点で最大値 56）
	}
	Level_up string //	"0": ?
}



//	星・種をまとめて投げる
func ApiLiveBulkGiftingFree(
	client *http.Client, //	HTTPクライアント
	csrftoken string, //	接続の識別子
	liveid string, //	配信枠の識別子
) (
	bulkgiftingfree *BulkGiftingFree,
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
	values.Add("live_id", liveid)
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

	if e = json.NewDecoder(resp.Body).Decode(&bulkgiftingfree); e != nil {
		err = fmt.Errorf("json.NewDecoder(resp.Body).Decode(): %w", e)
		return
	}

	return

}