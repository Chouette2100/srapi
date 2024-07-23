package srapi

import (
	//	"time"
	"bytes"
	"fmt"
	//	"log"

	"encoding/json"
	"net/http"
	"net/url"
)

type RoomInfAll struct {
	PrevLeagueID int `json:"prev_league_id"`
	ImageList    []struct {
		MOriginal string `json:"m_original"`
		Ts        int    `json:"ts"`
		ID        int    `json:"id"`
		M         string `json:"m"`
	} `json:"image_list"`
	BannerList []struct {
		URL   string `json:"url"`
		Image string `json:"image"`
	} `json:"banner_list"`
	IsTalkOnline   bool        `json:"is_talk_online"`
	AwardList      interface{} `json:"award_list"` //	<== 要確認
	PushSendStatus struct {    //	<== 要確認
	} `json:"push_send_status"`
	PerformerName      string `json:"performer_name"`
	FollowerNum        int    `json:"follower_num"` //	フォロワー数
	LiveContinuousDays int    `json:"live_continuous_days"`
	NextLeagueID       int    `json:"next_league_id"`
	LiveID             int    `json:"live_id"`
	LeagueID           int    `json:"league_id"`
	IsOfficial         bool   `json:"is_official"`
	IsFollow           bool   `json:"is_follow"`
	VoiceList          []struct {
		MOriginal interface{} `json:"m_original"`
		Ts        int         `json:"ts"`
		ID        int         `json:"id"`
		M         string      `json:"m"`
	} `json:"voice_list"`
	ShowRankSubdivided string `json:"show_rank_subdivided"` //	現在のランク 例 SS-5
	Event              struct {
		EndedAt   int64    `json:"ended_at"`   //	unix time
		EventID   int    `json:"event_id"`   //	これが本来のイベント識別子？
		StartedAt int64    `json:"started_at"` //	unix time
		URL       string `json:"url"`        //	最後のフィールドが event_url_key (プログラム中ではこれを eventid としていることがある。
		Name      string `json:"name"`
		Label     struct {
			Color string `json:"color"`
			Text  string `json:"text"`
		} `json:"label"`
		Image string `json:"image"`
	} `json:"event"`
	IsBirthday  bool   `json:"is_birthday"`
	Description string `json:"description"`
	LiveTags    []struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"live_tags"`
	GenreID         int `json:"genre_id"` //	101: Music,
	DisplayedMedals []struct {
		MedalID     int    `json:"medal_id"`
		EventName   string `json:"event_name"`
		MedalName   string `json:"medal_name"`
		Displayed   int    `json:"displayed"`
		ImageURL    string `json:"image_url"`
		Description string `json:"description"`
	} `json:"displayed_medals"`
	PrevScore            int    `json:"prev_score"` //	下位のランクまでのポイント
	YoutubeID            string `json:"youtube_id"`
	VisitCount           int    `json:"visit_count"`
	RecommendCommentList []struct {
		CreatedAt int    `json:"created_at"`
		Comment   string `json:"comment"`
		User      struct {
			Name  string `json:"name"`
			Image string `json:"image"`
		} `json:"user"`
	} `json:"recommend_comment_list"`
	CurrentLiveStartedAt   int64    `json:"current_live_started_at"` //	unix time
	NextShowRankSubdivided string `json:"next_show_rank_subdivided"`
	ShareTextLive          string `json:"share_text_live"`
	SnsList                []struct {
		Icon string `json:"icon"`
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"sns_list"`
	RecommendCommentsURL    string `json:"recommend_comments_url"`
	ShareURL                string `json:"share_url"`
	RoomURLKey              string `json:"room_url_key"`	// 配信画面のURLの最後のフィールド
	LeagueLabel             string `json:"league_label"` //	例 SS
	IsLiveTagCampaignOpened bool   `json:"is_live_tag_campaign_opened"`
	Avatar                  struct {
		Description string   `json:"description"`
		List        []string `json:"list"`
	} `json:"avatar"`
	ShareURLLive            string `json:"share_url_live"`
	PrevShowRankSubdivided  string `json:"prev_show_rank_subdivided"` //	例 SS-4
	IsTalkOpened            bool   `json:"is_talk_opened"`
	ImageSquare             string `json:"image_square"`
	RecommendCommentPostURL string `json:"recommend_comment_post_url"`
	GenreName               string `json:"genre_name"` //	ジャンル名
	RoomName                string `json:"room_name"`  //	ルーム名　 = main_name ?
	Birthday                int64    `json:"birthday"`
	RoomLevel               int    `json:"room_level"` //	ルームレベル
	PartyLiveStatus         int    `json:"party_live_status"`
	Party                   struct {
	} `json:"party"`
	EcConfig struct {
		SalesAvailable int           `json:"sales_available"`
		IsExternalEc   int           `json:"is_external_ec"`
		Links          []interface{} `json:"links"`
	} `json:"ec_config"`
	Image                      string `json:"image"`
	RecommendCommentOpenStatus int    `json:"recommend_comment_open_status"`
	MainName                   string `json:"main_name"` //	ルーム名　 = room_name ?
	ViewNum                    int    `json:"view_num"`
	HasMoreRecommendComment    bool   `json:"has_more_recommend_comment"`
	IsPartyEnabled             bool   `json:"is_party_enabled"`
	PremiumRoomType            int    `json:"premium_room_type"`
	NextScore                  int    `json:"next_score"` //	上位のランクまでのポイント
	IsOnlive                   bool   `json:"is_onlive"`
	RoomID                     int    `json:"room_id"` //	ルーム識別子
	//	エラーが発生したときのレスポンス
	Errors                     []struct {
		ErrorUserMsg string `json:"error_user_msg"`
		Message      string `json:"message"`
		Code         int    `json:"code"`
	} `json:"errors"`
}

func ApiRoomProfile(
	client *http.Client,
	room_id string,
) (
	proominfall *RoomInfAll,
	err error,
) {

	//	https://qiita.com/takeru7584/items/f4ba4c31551204279ed2

	//	url := "https://www.showroom-live.com/api/room/profile?room_id=" + room_id
	turl := "https://www.showroom-live.com/api/room/profile"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{}         // url.Valuesオブジェクト生成
	values.Add("room_id", room_id) // key-valueを追加

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
	//	bufstr := buf.String()

	//	JSONをデコードする。
	//	次の記事を参考にさせていただいております。
	//		Go言語でJSONに泣かないためのコーディングパターン
	//		https://qiita.com/msh5/items/dc524e38073ed8e3831b

	proominfall = new(RoomInfAll)

	decoder := json.NewDecoder(buf)
	if err = decoder.Decode(proominfall); err != nil {
		//	log.Printf("decoder.Decode(proominfall) err: %v", err)
		//	log.Printf(" room?id= %s", room_id)
		//	log.Printf("bufstr: %s", bufstr)
		err = fmt.Errorf("decoder.Decode(proominfall)(room_id=%s) err: %v", room_id, err)
		return
	}

	return

}
