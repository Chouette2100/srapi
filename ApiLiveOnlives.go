/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php


Ver. 0.0.0
Ver. 1.0.0 戻り値 status を err に変更する。
Ver. 2.0.0 ExtrRoomLiveByCtg()とExtrRoomLiveByGnr()をRoomOnlivesのメソッドとする。
Ver. 2.0.1 fmt.Errorf()の %v を %w に変更する。 
Ver. 3.0.0 ExtrRoomLiveByCtg(), ExtrRoomLiveyypByGnr()をそれぞれExtrByCtg()、ExtrByCtg()に変更する。
Ver. 4.0.0 ExtrByCtg()、ExtrByCtg()の引数を Roomonlives から *RoomOnLives に変更する。
Ver. 4.1.0 上位でsort.Sort()をsort.Slice()に変更したため、sort.Sort()のためのメソッドを削除する。
Ver. 4.1.0 ExtrByCtg()の引数を Roomonlives から *RoomOnLives に変更する（修正もれ）

*/
package srapi

import (
	"fmt"

	"net/http"
	"net/url"

	"encoding/json"
)

type Live struct {
	//	配信中ルーム情報ルーム
	Room_id      int    //	ルームID 配信者を識別する
	Room_url_key string //	配信のURLの最後のフィールド
	Live_id      int    //	Live ID 配信を識別する
	Started_at   int64  //	配信開始時刻（Unix Time）
	View_num     int    //	視聴者数
	Main_name    string //	ルーム名
	Genre_id     int    //	0: 人気、 200: フリー、 100(?)〜199(?): 公式、 700〜: 特定属性
}
type Lives []Live

type RoomOnlives struct {
	//	配信中ルーム一覧
	//	Onlives []OnLive //	ジャンルごとの配信中ルーム一覧
	Onlives []struct {
		Genre_id   int    //	0: 人気、 200: フリー、 100(?)〜199(?): 公式、 700〜: 特定属性
		Genre_name string //	ジャンル名
		Lives      Lives  //	配信中ルーム一覧
	}

	Bcsvr_post int
	Bcsvr_host string
}
//	指定したカテゴリー（"Free", "Official", "All"）のルーム一覧を作る。
//	"All"のときでもGenre_idが0や700以上は含まないので重複はない。
func (r *RoomOnlives)ExtrByCtg(
	tgt string, //	カテゴリ
) (
	roomlive *Lives, //	配信中ルーム情報
	err error, //	エラー
) {
	roomlive = new(Lives)
	for _, onlives := range r.Onlives {
		switch {
		case (tgt == "Free" || tgt == "All") && onlives.Genre_id == 200:
			fallthrough
		case (tgt == "Official" || tgt == "All") && (onlives.Genre_id >= 100 && onlives.Genre_id < 200):
			//	log.Printf("%d %s\n", onlives.Genre_id, onlives.Genre_name)
			*roomlive = append(*roomlive, onlives.Lives...)
		default:
		}
	}
	return	roomlive, nil
}

//	指定したジャンルのルーム一覧を作る。
func (r *RoomOnlives)ExtrByGnr(
	gnr map[string]bool, //	抽出したいジャンル、mapにジャンルIDがありTrueであれば抽出する。
) (
	roomlive *Lives, //	配信中ルーム情報
	err error, //	エラー
) {
	roomlive = new(Lives)
	for _, onlive := range r.Onlives {
		if ok, val := gnr[onlive.Genre_name]; ok && val {
			*roomlive = append(*roomlive, onlive.Lives...)
		}
	}
	return roomlive, nil
}

//	配信中のルームの一覧を取得する
func ApiLiveOnlives(
	client *http.Client, //	HTTP client
) (
	roomonlives *RoomOnlives, //	配信中ルームのジャンル別一覧
	err 	error, //	エラー
) {

	turl := "https://www.showroom-live.com/api/live/onlives"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse: %w", err)
		return nil, err
	}

	// Request を生成
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		err = fmt.Errorf("http.NewRequest(): %w", err)
		return
	}

	// User-Agentを設定
	req.Header.Add("User-Agent", useragent)
	req.Header.Add("Accept-Language", "ja-JP")

	// Doメソッドでリクエストを投げる
	// http.Response型のポインタ（とerror）が返ってくる
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("client.Do(): %w", err)
		return
	}

	defer resp.Body.Close()

	roomonlives = new(RoomOnlives) //	ここで作られたRoomOnlives型の領域は参照可能な限り（関数外でも）存在します。
	if err = json.NewDecoder(resp.Body).Decode(&roomonlives); err != nil {
		err = fmt.Errorf("json.Decoder: %w", err)
		return nil, err
	}

	return roomonlives, nil
}

