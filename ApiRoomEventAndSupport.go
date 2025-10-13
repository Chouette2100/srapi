/*
!
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
/*
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
*/
type RoomEvnetAndSupport struct {
	Support      interface{} `json:"support"`
	ResEvent     ResEvent    `json:"event"`
	RegularEvent interface{} `json:"regular_event"`
}
type EndAnimation struct {
	TriggerTime int `json:"trigger_time"`
	Value       int `json:"value"`
	Type        int `json:"type"`
}
type ContributionRanking struct {
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	Point     int    `json:"point"`
	Rank      int    `json:"rank"`
	AvatarID  int    `json:"avatar_id"`
	AvatarURL string `json:"avatar_url"`
}
type SupportUsers struct {
	UserID    int    `json:"user_id"`
	Point     int    `json:"point"`
	Order     int    `json:"order"`
	Name      string `json:"name"`
	AvatarID  int    `json:"avatar_id"`
	AvatarURL string `json:"avatar_url"`
}
type Support struct {
	SupportID      int            `json:"support_id"`
	GoalPoint      int            `json:"goal_point"`
	IsAchieved     bool           `json:"is_achieved"`
	CurrentPoint   int            `json:"current_point"`
	NextLevel      int            `json:"next_level"`
	Title          string         `json:"title"`
	TutorialURL    string         `json:"tutorial_url"`
	SupportUsers   []SupportUsers `json:"support_users"`
	SupportMyPoint int            `json:"support_my_point"`
}
type QuestList struct {
	SupportID     int    `json:"support_id"`
	QuestLevel    int    `json:"quest_level"`
	Title         string `json:"title"`
	GoalPoint     int    `json:"goal_point"`
	NumberOfItems int    `json:"number_of_items"`
	RestItems     int    `json:"rest_items"`
	Color         string `json:"color"`
	IsAcquired    bool   `json:"is_acquired"`
}
type Quest struct {
	QuestLevel          int         `json:"quest_level"`
	Support             Support     `json:"support"`
	QuestList           []QuestList `json:"quest_list"`
	Text                string      `json:"text"`
	ContributorListURL  string      `json:"contributor_list_url"`
	IsAllQuestCompleted int         `json:"is_all_quest_completed"`
}
type ResRanking struct {
	BeforeRank           int    `json:"before_rank"`
	MaxRank              int    `json:"max_rank"`
	Point                int    `json:"point"`
	Gap                  int    `json:"gap"`
	Text                 string `json:"text"`
	SequenceNum          int    `json:"sequence_num"`
	IsAnimation          int    `json:"is_animation"`
	NextRank             int    `json:"next_rank"`
	Rank                 int    `json:"rank"`
	LargeRank            int    `json:"large_rank"`
	LowerRank            int    `json:"lower_rank"`
	LowerGap             int    `json:"lower_gap"`
	EventBlockDivisionID int    `json:"event_block_division_id"`
}
type ResEvent struct {
	EventID                 int                   `json:"event_id"`
	EventName               string                `json:"event_name"`
	Image                   string                `json:"image"`
	EventURL                string                `json:"event_url"`
	EndAnimation            []EndAnimation        `json:"end_animation"`
	EventType               string                `json:"event_type"`
	StartedAt               int                   `json:"started_at"`
	EndedAt                 int                   `json:"ended_at"`
	RemainTime              int                   `json:"remain_time"`
	EventDescription        string                `json:"event_description"`
	AdditionalEventPoints   []interface{}         `json:"additional_event_points"`
	AdditionalEventPointSum int                   `json:"additional_event_point_sum"`
	TutorialURL             string                `json:"tutorial_url"`
	ContributionRanking     []ContributionRanking `json:"contribution_ranking"`
	Quest                   Quest                 `json:"quest"`
	ShowRanking             int                   `json:"show_ranking"`
	ResRanking              ResRanking            `json:"ranking"`
}

// イベントでの獲得ポイントを取得する
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
	blockid int, //		ブロックイベントのブロックID
	err error, //	エラー情報
) {

	res, e := ApiRoomEventAndSupport(client, fmt.Sprintf("%d", roomid))
	if e != nil {
		err = fmt.Errorf("ApiRoomEventAndSupport(): %w", e)
		return
	}

	url := strings.Split(res.ResEvent.EventURL, "/")
	surl := url[len(url)-1]
	if res.ResEvent.ResRanking.Rank != 0 {
		//	ランキングイベントの場合はRankingから取得する
		return res.ResEvent.ResRanking.Point, res.ResEvent.ResRanking.Rank,
			res.ResEvent.ResRanking.Gap, res.ResEvent.EventID, surl,
			res.ResEvent.EventName, res.ResEvent.ResRanking.EventBlockDivisionID, nil
	} else {
		//	レベルイベントの場合はQuestから取得する
		return res.ResEvent.Quest.Support.CurrentPoint, res.ResEvent.Quest.QuestLevel - 10000, -1, res.ResEvent.EventID,
			surl, res.ResEvent.EventName, res.ResEvent.ResRanking.EventBlockDivisionID, nil

	}
}

// イベントの順位と獲得ポイントを知るAPI（/api/room/event_and_suport）を実行する。
func ApiRoomEventAndSupport(
	client *http.Client, //	HTTPクライアント
	roomid string, //	配信ルームの識別子（プロフィールやファンルームのURLの最後にある6桁程度の数）
) (
	res *RoomEvnetAndSupport, //	ファンレベルの進捗状況の詳細
	err error, //	エラー情報
) {
	res = &RoomEvnetAndSupport{}
	var buf *bytes.Buffer

	buf, err = JsonRoomEventAndSupport(client, roomid)
	if err != nil {
		err = fmt.Errorf("JsonRoomEventAndSupport(): %w", err)
		return nil, err
	}

	//	log.Printf("buf=%s\n", buf.String())

	if err = json.NewDecoder(buf).Decode(res); err != nil {
		err = fmt.Errorf("json.NewDecoder: %w", err)
		return nil, err
	}

	//	log.Printf("res = %+v\n", res)

	return res, nil
}
func JsonRoomEventAndSupport(
	client *http.Client, //	HTTPクライアント
	roomid string, //	配信ルームの識別子（プロフィールやファンルームのURLの最後にある6桁程度の数）
) (
	buf *bytes.Buffer, //	JSON文字列を格納したバッファ
	err error, //	エラー情報
) {

	turl := "https://www.showroom-live.com/api/room/event_and_support"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse: %w", err)
		return nil, err
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成
	values.Add("room_id", roomid)

	//	log.Printf("values=%+v\n", values)

	// Request を生成
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		err = fmt.Errorf("http.NewRequest: %w", err)
		return nil, err
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
		return nil, err
	}

	// 関数を抜ける際に接続を切断し、リソースを解放するため必ずresponse.Bodyをcloseする
	defer resp.Body.Close()

	buf = new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	//	bufstr := buf.String()
	//	log.Printf("bufstr=%s\n", bufstr)

	return buf, nil

}
