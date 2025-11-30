// Copyright © 2025 chouette.21.00@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package srapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type GiftLoglist struct {
	GiftLog []GiftLog `json:"gift_log"`
}
type GiftLog struct {
	Aft       int    `json:"aft"`
	AvatarID  int    `json:"avatar_id"`
	AvatarURL string `json:"avatar_url"`
	CreatedAt int    `json:"created_at"`
	GiftID    int    `json:"gift_id"`
	Image     string `json:"image"`
	Image2    string `json:"image2"`
	Name      string `json:"name"`
	Num       int    `json:"num"`
	Ua        int    `json:"ua"`
	UserID    int    `json:"user_id"`
}

// 配信者ルームのギフトログを取得する
func ApiLiveGiftlog(client *http.Client, roomid int) (gll *GiftLoglist, err error) {

	turl := "https://www.showroom-live.com/api/live/gift_log"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse: %w", err)
		return nil, err
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成

	// クエリを組み立て
	values.Add("room_id", fmt.Sprintf("%d", roomid)) // key-valueを追加

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
	req.Header.Add("Accept-Language", "ja-JP")

	// Doメソッドでリクエストを投げる
	// http.Response型のポインタ（とerror）が返ってくる
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("client.Do(): %w", err)
		return
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	//	bufstr := buf.String()
	//	log.Printf("bufstr=%s\n", bufstr)

	gll = new(GiftLoglist) //	ここで作られたLiveGiftlog型の領域は参照可能な限り（関数外でも）存在します。
	if err = json.NewDecoder(buf).Decode(gll); err != nil {
		return nil, err
	}

	return gll, nil
}
