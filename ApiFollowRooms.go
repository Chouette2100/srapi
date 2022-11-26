/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php


Ver. 0.1.0

*/
package srapi

import (
	"bytes"
	"fmt"

	"encoding/json"
	"net/http"
	"net/url"
)

type RoomFollowed struct {
	Is_official      int    // 0:Free, 1:Official
	Image_l          string //  画像
	Room_description string //  ルーム概要
	Room_url_key     string //  配信ページのURLの最後のフィールド
	Next_live        string //	"未定"または"11/25 18:00"の形式
	Image_m          string //  未確認
	Has_next_live    int    //  null:配信予定なし, 1:配信予定あり
	Is_online        int    //  0:  配信中ではない,
	Room_id          string //  RoomIDを文字列にしたもの
	Room_name        string //  ルーム名
	Open_fan_club    int    //  null:ファンルームなし、1:ファンルームあり
}

type FollowRooms struct {
	Next_page     int //  次ページのページ番号、null: 次ページなし
	Previous_page int //  前ページのページ番号、null: 前ページなし
	Total_Entries int //  フォローしている配信者の数
	First_page    int //  最初のページのページ番号（常に1 ？）
	Current_page  int //  現在のページのページ番号
	Rooms         []RoomFollowed
}

//	フォローしているルームの一覧を指定した範囲で取得する。
func ApiFollowRooms(
	client *http.Client, //	HTTPクライアント
	page int,
	count int,
) (
	followrooms *FollowRooms,
	err error,
) {

	turl := "https://www.showroom-live.com/api/follow/rooms"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{}                        // url.Valuesオブジェクト生成
	values.Add("page", fmt.Sprintf("%d", page))   // key-valueを追加
	values.Add("count", fmt.Sprintf("%d", count)) // key-valueを追加

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

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bufstr := buf.String()

	followrooms = new(FollowRooms)
	if err = json.NewDecoder(buf).Decode(followrooms); err != nil {
		err = fmt.Errorf("%w(resp.Body: %s)", err, bufstr)
		err = fmt.Errorf("json.NewDecoder(resp.Body).Decode(followrooms): %w", err)
		return
	}
	return
}

//	フォローしているすべてのルームの一覧を取得する。
func GetFollowRoomsByApi(
	client *http.Client, //	HTTPクライアント
) (
	roomsfollowed []RoomFollowed,
	err error,
) {

	roomsfollowed = make([]RoomFollowed, 0)

	page := 1
	count := 200
	var followrooms *FollowRooms

	for {
		//	count数ずつ分割して読み込んでいくわけですが、countを大きな値にして一挙にすべて読み込んだ方がいいでしょう。
		//	（このAPIではそういうことはないと思われますが）分割して読み込んだときデータの欠落や重複が発生することがあるAPIもあるようなので。
		//	ただ、count数に上限があるのかは確かめていません。数十ルームしかフォローしていないもので...

		followrooms, err = ApiFollowRooms(client, page, count)
		if err != nil {
			err = fmt.Errorf("ApiFollowRooms(): %w", err)
			return
		}

		roomsfollowed = append(roomsfollowed, followrooms.Rooms...)

		if followrooms.Next_page > 0 {
			//	次のページがある場合はよみこみを続ける。
			page = followrooms.Next_page
		} else {
			return
		}
	}

}
