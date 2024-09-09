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

type CdnGiftRanking struct {
	TotalScore  int           `json:"total_score"`
	RankingList []GrRanking	`json:"ranking_list"`
		Errors []struct { //	例えば gift_id が（整数でなく）アルファベットの場合
		ErrorUserMsg string `json:"error_user_msg"`
		Message      string `json:"message"`
		Code         int    `json:"code"`
	} `json:"errors"`
}
type GrRoom struct {
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
type GrRanking struct {
	RoomID  int  `json:"room_id"`
	Rank    int  `json:"rank"`
	Score   int  `json:"score"`
	OrderNo int  `json:"order_no"`
	Room    GrRoom `json:"room"`
}
// ギフトランキングを取得する
//	(イベントギフトランキングとは別のものです)
//	https://public-api.showroom-cdn.com/gift_ranking/mmm?limit=nnn		mmm: ジャンルID　nnn;　最大取得件数
//
//	|コード|名称|補足|
//	|---|---|---|
//	|486|人気ライバーランキング||
//	|490|新人スタートダッシュ||
//	|494|アイドル||
//	|495|俳優||
//	|496|アナウンサー||
//	|497|グローバル||
//	|498|声優||
//	|499|芸人||
//	|500|タレント||
//	|501|ライバー||
//	|502|モデル||
//	|503|バーチャル||
//	|504|アーティスト||
func ApiCdnGiftRanking(
	client *http.Client, //	HTTPクライアント
	genre_id int, //	ジャンルID
	limit int, //	最大取得件数
) (
	pranking *CdnGiftRanking,
	err error,
) {
	turl := fmt.Sprintf("https://public-api.showroom-cdn.com/gift_ranking/%d", genre_id)
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成

	//	クエリを組み立て
	values.Add("limit", fmt.Sprintf("%d", limit)) // key-valueを追加

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

	pranking = &CdnGiftRanking{}

	if err = json.NewDecoder(buf).Decode(pranking); err != nil {
		err = fmt.Errorf("%w(buf: %s)", err, bufstr)
		err = fmt.Errorf("json.NewDecoder(buf).Decode(pranking): %w", err)
		return
	}
	//	fmt.Printf("pranking: %+v", pranking)
	return
}


type CdnUserGiftRanking struct {
	RankingList []UgrRanking `json:"ranking_list"`
	Errors []struct {
		ErrorUserMsg string `json:"error_user_msg"`
		Message      string `json:"message"`
		Code         int    `json:"code"`
	} `json:"errors"`
}
type UgrUser struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}
type UgrRanking struct {
	UserID  int  `json:"user_id"`
	Rank    int  `json:"rank"`
	Score   int  `json:"score"`
	OrderNo int  `json:"order_no"`
	User    UgrUser `json:"user"`
}


// ユーザーギフトランキングを取得する
//
//	https://public-api.showroom-cdn.com/user_gift_ranking/mmm?limit=nnn
func ApiCdnUserGiftRanking(
	client *http.Client, //	HTTPクライアント
	genre_id int, //	206
	limit int, //	最大取得件数
) (
	pranking *CdnUserGiftRanking,
	err error,
) {
	turl := fmt.Sprintf("https://public-api.showroom-cdn.com/user_gift_ranking/%d", genre_id)
	//	turl := "https://public-api.showroom-cdn.com/user_gift_ranking/AAA"	//	エラーのテスト用
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成

	//	クエリを組み立て
	values.Add("limit", fmt.Sprintf("%d", limit)) // key-valueを追加

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

	pranking = &CdnUserGiftRanking{}

	if err = json.NewDecoder(buf).Decode(pranking); err != nil {
		err = fmt.Errorf("%w(buf: %s)", err, bufstr)
		err = fmt.Errorf("json.NewDecoder(buf).Decode(pranking): %w", err)
		return
	}
	//	fmt.Printf("pranking: %+v", pranking)
	return
}
