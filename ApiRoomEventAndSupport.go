/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

Ver. 0.0.0
Ver. 0.1.0 レベルイベントのRankとGapを−１とする。

*/
package srapi

import (
	"bytes"
	"fmt"
	//	"log"
	"strings"

	"net/http"
	"net/url"

	"encoding/json"

)

//	イベントの順位と獲得ポイント
//	JSONには他にもいろんなフィールドがあります。
//	ぜったい必要そうなものとフィールドの作りが面倒そうなところだけ構造体にしてあります。
type RoomEvnetAndSupport struct {
	Event struct {
		Event_id   int      //	イベント識別子
		Event_url  string   //	イベントのURLの末尾のフィールド
		Event_name string   //	イベント名
		Ranking   struct { //	ランキングイベントのポイントデータ
			Rank  int //	順位
			Point int //	獲得ポイント
			Gap   int //	上位とのポイント差
		}
		Quest struct { //	レベルイベントのポイントデータ
			Support struct {
				Goal_point    int //	目標ポイント
				Current_point int //	現在のポイント
				Support_users []struct {
					Point int //	ポイント
					Order int //	順位
					Name string //	ユーザー名
					User_id int //	ユーザー識別子
				}
			}
			Quest_list []struct {
				Number_of_items int //	アイテム数
				Color		   string//	色
				Goal_point	  int //	目標ポイント
				Rest_time	  int //	残り時間？
				Title	  string//	タイトル
				Quest_level	  int //	クエストレベル
				Is_aquired	  bool//	クエスト獲得フラグ？
			}
		}
	}
}

//	イベントでの獲得ポイントを取得する
func GetPointByApi(
	client *http.Client, //	HTTPクライアント
	roomid int, //	配信ルームの識別子（プロフィールやファンルームのURLの最後にある6桁程度の数）
) (
	point int, //	獲得ポイント
	rank int, //	順位
	gap int, //	上位とのポイント差
	eventid int, //	イベント識別子
	eventurl string, //	イベントのURLの末尾のフィールド
	eventname string, //	イベント名
	err error, //	エラー情報
) {
	
	res, e := ApiRoomEventAndSupport(client, fmt.Sprintf("%d", roomid))
	if e != nil {
		err = fmt.Errorf("ApiRoomEventAndSupport(): %w", e)
		return
	}

	url := strings.Split(res.Event.Event_url, "/")
	surl := url[len(url)-1]
	if res.Event.Ranking.Rank != 0 {
		//	ランキングイベントの場合はRankingから取得する
		return res.Event.Ranking.Point, res.Event.Ranking.Rank, res.Event.Ranking.Gap, res.Event.Event_id, surl, res.Event.Event_name, nil
	} else {
		//	レベルイベントの場合はQuestから取得する
		return res.Event.Quest.Support.Current_point, -1, -1, res.Event.Event_id, surl, res.Event.Event_name, nil

	}
}

//	イベントの順位と獲得ポイントを知るAPI（/api/room/event_and_suport）を実行する。
func ApiRoomEventAndSupport(
	client *http.Client, //	HTTPクライアント
	roomid string, //	配信ルームの識別子（プロフィールやファンルームのURLの最後にある6桁程度の数）
) (
	res *RoomEvnetAndSupport, //	ファンレベルの進捗状況の詳細
	err error, //	エラー情報
) {

	res = &RoomEvnetAndSupport{}

	turl := "https://www.showroom-live.com/api/room/event_and_support"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse: %w", err)
		return res, err
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成
	values.Add("room_id", roomid)

	//	log.Printf("values=%+v\n", values)

	// Request を生成
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		err = fmt.Errorf("http.NewRequest: %w", err)
		return res, err
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
		err = fmt.Errorf("client.Do: %w", err)
		return res, err
	}

	// 関数を抜ける際に接続を切断し、リソースを解放するため必ずresponse.Bodyをcloseする
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	//	bufstr := buf.String()
	//	log.Printf("bufstr=%s\n", bufstr)


	if err = json.NewDecoder(buf).Decode(res); err != nil {
		err = fmt.Errorf("json.NewDecoder: %w", err)
		return nil, err
	}

	//	log.Printf("res = %+v\n", res)

	return res, nil

}
