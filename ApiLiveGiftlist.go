// Copyright © 2025 chouette.21.00@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package srapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type LiveGiftlist struct {
	Enquete []Enquete `json:"enquete,omitempty"`
	Normal  []Normal  `json:"normal,omitempty"`
}
type Enquete struct {
	Icon              int    `json:"icon,omitempty"`
	IsHidden          bool   `json:"is_hidden,omitempty"`
	OrderNo           int    `json:"order_no,omitempty"`
	GiftType          int    `json:"gift_type,omitempty"`
	Image             string `json:"image,omitempty"`
	GiftID            int    `json:"gift_id,omitempty"`
	Image2            string `json:"image2,omitempty"`
	Free              bool   `json:"free,omitempty"`
	Point             int    `json:"point,omitempty"`
	IsDeleteFromStage bool   `json:"is_delete_from_stage,omitempty"`
	GiftName          string `json:"gift_name,omitempty"`
	Scale             int    `json:"scale,omitempty"`
	Label             int    `json:"label,omitempty"`
	DialogID          int    `json:"dialog_id,omitempty"`
}
type Normal struct {
	Icon              int    `json:"icon,omitempty"`
	IsHidden          bool   `json:"is_hidden,omitempty"`
	OrderNo           int    `json:"order_no,omitempty"`
	GiftType          int    `json:"gift_type,omitempty"`
	Image             string `json:"image,omitempty"`
	GiftID            int    `json:"gift_id,omitempty"`
	Image2            string `json:"image2,omitempty"`
	Free              bool   `json:"free,omitempty"`
	Point             int    `json:"point,omitempty"`
	IsDeleteFromStage bool   `json:"is_delete_from_stage,omitempty"`
	GiftName          string `json:"gift_name,omitempty"`
	Scale             int    `json:"scale,omitempty"`
	Label             int    `json:"label,omitempty"`
	DialogID          int    `json:"dialog_id,omitempty"`
}

// 配信者ルームで使用可能なギフトの一覧を取得する
func ApiLiveGiftlist(client *http.Client, roomid int) (lgl *LiveGiftlist, err error) {

	turl := "https://www.showroom-live.com/api/live/gift_list"
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

	lgl = new(LiveGiftlist) //	ここで作られたLiveGiftlist型の領域は参照可能な限り（関数外でも）存在します。
	if err = json.NewDecoder(resp.Body).Decode(lgl); err != nil {
		err = fmt.Errorf("json.Decoder: %w", err)
		return nil, err
	}

	return lgl, nil
}
