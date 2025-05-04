// Copyright © 2024 chouette.21.00@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package srapi

import (
	"bytes"
	"fmt"

	"encoding/json"
	"net/http"
	"net/url"
)

type CurrentUser struct {
	UserFanLevel         int    `json:"user_fan_level"`
	OwnRoomID            int    `json:"own_room_id"`
	EventOrganizerID     int    `json:"event_organizer_id"`
	HasUserUnreadNotice  int    `json:"has_user_unread_notice"`
	IsEvnetOrgPending    bool   `json:"is_evnet_org_pending"`
	HasEventNotice       bool   `json:"has_event_notice"`
	IsLogin              bool   `json:"is_login"`
	OrganizerName        string `json:"organizer_name"`
	OrganizerID          int    `json:"organizer_id"`
	OwnRoomURLKey        string `json:"own_room_url_key"`
	UserGold             int    `json:"user_gold"`
	OrganizerAccountID   string `json:"organizer_account_id"`
	UserExpiringGold     int    `json:"user_expiring_gold"`
	HasPaymentOrganaizer bool   `json:"has_payment_organaizer"`
	UserName             string `json:"user_name"`
	AccountID            string `json:"account_id"`
	SmsAuth              bool   `json:"sms_auth"`
	ContributionPoint    int    `json:"contribution_point"`
	HasUnreadUserNotice  int    `json:"has_unread_user_notice"`
	ImageURL             string `json:"image_url"`
	UserID               int    `json:"user_id"`
	IsOrganizer          bool   `json:"is_organizer"`
}

// ログインしているユーザの情報を取得する
func ApiCurrentUser(
	client *http.Client, //	HTTPクライアント
) (
	pcu *CurrentUser,
	err error,
) {
	turl := "https://www.showroom-live.com/api/current_user"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成

	//	クエリを組み立て
	// values.Add("limit", fmt.Sprintf("%d", limit)) // key-valueを追加

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
	// 関数を抜ける際に接続を切断し、リソースを解放するため必ずresponse.Bodyをcloseする
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bufstr := buf.String()

	//	fmt.Printf("bufstr: %s", bufstr)

	pcu = &CurrentUser{}

	if err = json.NewDecoder(buf).Decode(pcu); err != nil {
		err = fmt.Errorf("%w(buf: %s)", err, bufstr)
		err = fmt.Errorf("json.NewDecoder(buf).Decode(pranking): %w", err)
		return
	}
	// fmt.Printf("CurrentUser: %+v", pcu)
	return
}
