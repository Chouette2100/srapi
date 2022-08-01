/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php


Ver. 0.0.0

*/
package srapi

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
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

//      sort.Sort()のための関数
func (r Lives) Len() int {
        return len(r)
}

func (r Lives) Swap(i, j int) {
        r[i], r[j] = r[j], r[i]
}

func (r Lives) Choose(from, to int) (s Lives) {
        s = r[from:to]
        return
}

func (r Lives) Less(i, j int) bool {
        return r[i].Started_at > r[j].Started_at
}




type RoomOnlives struct {
	//	配信中ルーム一覧
	//	Onlives []OnLive //	ジャンルごとの配信中ルーム一覧
	Onlives []struct {
		Genre_id   int    //	0: 人気、 200: フリー、 100(?)〜199(?): 公式、 700〜: 特定属性
		Genre_name string //	ジャンル名
		Lives      Lives //	配信中ルーム一覧
	}

	Bcsvr_post int
	Bcsvr_host string
}

//	配信中のルームの一覧を取得する
func ApiLiveOnlives(
	client *http.Client, //	HTTP client
) (
	roomonlives *RoomOnlives, //	配信中ルームのジャンル別一覧
	status int, //	0: 正常終了
) {

	status = 0

	turl := "https://www.showroom-live.com/api/live/onlives"
	u, err := url.Parse(turl)
	if err != nil {
		log.Printf("url.Parse() returned error %s\n", err.Error())
		status = -1
		return
	}
	resp, err := client.Get(u.String())
	if err != nil {
		log.Printf("client.Get() returned error %s\n", err.Error())
		status = -2
		return
	}
	defer resp.Body.Close()

	roomonlives = new(RoomOnlives) //	ここで作られたRoomOnlives型の領域は参照可能な限り（関数外でも）存在します。
	if err := json.NewDecoder(resp.Body).Decode(&roomonlives); err != nil {
		log.Printf("json.NewDecoder() returned error %s\n", err.Error())
		status = -3
		return
	}

	return
}

//	指定したカテゴリー（"Free", "Official", "All"）のルーム一覧を作る。
//	"All"のときでもGenre_idが0や700以上は含まないので重複はない。
func ExtrRoomLiveByCtg(
	roomonlives *RoomOnlives, //	配信中ルームのジャンル別一覧
	tgt string, //	カテゴリ
) (
	roomlive Lives, //	配信中ルーム情報
	status int, //	0: 正常終了
) {
	status = 0
	roomlive = make([]Live, 0)
	for _, onlives := range roomonlives.Onlives {
		switch {
		case (tgt == "Free" || tgt == "All") && onlives.Genre_id == 200:
			fallthrough
		case (tgt == "Official" || tgt == "All") && (onlives.Genre_id >= 100 && onlives.Genre_id < 200):
			//	log.Printf("%d %s\n", onlives.Genre_id, onlives.Genre_name)
			roomlive = append(roomlive, onlives.Lives...)
		default:
		}
	}
	return
}

//	指定したジャンルのルーム一覧を作る。
func ExtrRoomLiveByGnr(
	roomonlives *RoomOnlives, //	配信中ルームのジャンル別一覧
	gnr map[string]bool, //	抽出したいジャンル、mapにジャンルIDがありTrueであれば抽出する。
) (
	roomlive Lives, //	配信中ルーム情報
	status int, //	0: 正常終了
) {
	status = 0
	roomlive = make([]Live, 0)
	for _, onlive := range roomonlives.Onlives {
		if ok, val := gnr[onlive.Genre_name]; ok && val {
			roomlive = append(roomlive, onlive.Lives...)
		}
	}
	return
}
