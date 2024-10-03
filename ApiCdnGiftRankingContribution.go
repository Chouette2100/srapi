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

type GiftRankingContribution struct {
	Room        GrcRoom          `json:"room"`
	RankingList []GrcRanking     `json:"ranking_list"`
}
type GrcRoom struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	URLKey           string `json:"url_key"`
	ImageURL         string `json:"image_url"`
	Description      string `json:"description"`
	FollowerNum      int    `json:"follower_num"`
	IsLive           bool   `json:"is_live"`
	IsParty          bool   `json:"is_party"`
	NextLiveSchedule int    `json:"next_live_schedule"`
}
type GrcUser struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}
type GrcRanking struct {
	UserID  int     `json:"user_id"`
	Rank    int     `json:"rank"`
	Score   int     `json:"score"`
	OrderNo int     `json:"order_no"`
	User    GrcUser `json:"user"`
}

// ギフトランキングの獲得ギフトに対するリスナー貢献ランキングを取得する
//
//	https://public-api.showroom-cdn.com/gift_ranking/492/contribution/kogachan
func ApiCdnGiftRankingContribution(
	client *http.Client, //	HTTPクライアント
	genre_id int, //	ジャンルID
	url_key string, //	ルームうのurl_key
) (
	pranking *GiftRankingContribution,
	err error,
) {
	turl := fmt.Sprintf("https://public-api.showroom-cdn.com/gift_ranking/%d/contribution/%s", genre_id, url_key)
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成

	//	クエリを組み立て
	//	values.Add("limit", fmt.Sprintf("%d", limit)) // key-valueを追加

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

	pranking = &GiftRankingContribution{}

	if err = json.NewDecoder(buf).Decode(pranking); err != nil {
		err = fmt.Errorf("%w(buf: %s)", err, bufstr)
		err = fmt.Errorf("json.NewDecoder(buf).Decode(pranking): %w", err)
		return
	}
	//	fmt.Printf("pranking: %+v", pranking)
	return
}

