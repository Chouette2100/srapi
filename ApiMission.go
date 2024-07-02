//	Copyright © 2024 chouette.21.00@gmail.com
//	Released under the MIT license
//	https://opensource.org/licenses/mit-license.php

package srapi

import (
	//	"time"
	"bytes"
	"fmt"
	"log"

	"encoding/json"
	"net/http"
	"net/url"
)

type Mission struct {
	GenreList []struct {
		Priority      int    `json:"priority"`
		CurrentPeriod string `json:"current_period"`
		Night         struct {
			SingleMission []struct {
				Reward       string `json:"reward"`
				RemainReward int    `json:"remain_reward"`
				TargetValue  int    `json:"target_value"`
				MissionID    int    `json:"mission_id"`
				MaxLevel     int    `json:"max_level"`
				CurrentLevel int    `json:"current_level"`
				RewardType   string `json:"reward_type"`
				RewardID     int    `json:"reward_id"`
				RewardNum    int    `json:"reward_num"`
				CurrentValue int    `json:"current_value"`
				RewardURL    string `json:"reward_url"`
				Title        string `json:"title"`
				RewardName   string `json:"reward_name"`
			} `json:"single_mission"`
			CompositeMission struct {
				Reward       string `json:"reward"`
				RemainReward int    `json:"remain_reward"`
				TargetValue  int    `json:"target_value"`
				MissionID    int    `json:"mission_id"`
				MaxLevel     int    `json:"max_level"`
				CurrentLevel int    `json:"current_level"`
				RewardType   string `json:"reward_type"`
				RewardID     int    `json:"reward_id"`
				RewardNum    int    `json:"reward_num"`
				CurrentValue int    `json:"current_value"`
				RewardURL    string `json:"reward_url"`
				Title        string `json:"title"`
				RewardName   string `json:"reward_name"`
			} `json:"composite_mission"`
			EndAt             int `json:"end_at"`
			ContinuousMission []struct {
				Reward       string `json:"reward"`
				RemainReward int    `json:"remain_reward"`
				TargetValue  int    `json:"target_value"`
				MissionID    int    `json:"mission_id"`
				MaxLevel     int    `json:"max_level"`
				CurrentLevel int    `json:"current_level"`	//	視聴が終わったルームの数
				RewardType   string `json:"reward_type"`
				RewardID     int    `json:"reward_id"`
				RewardNum    int    `json:"reward_num"`
				CurrentValue int    `json:"current_value"`
				RewardURL    string `json:"reward_url"`
				Title        string `json:"title"`
				RewardName   string `json:"reward_name"`	//	"配信を視聴しよう（20/20）"
			} `json:"continuous_mission"`
			StartAt int `json:"start_at"`
		} `json:"night"`
		Day struct {
			SingleMission []struct {
				Reward       string `json:"reward"`
				RemainReward int    `json:"remain_reward"`
				TargetValue  int    `json:"target_value"`
				MissionID    int    `json:"mission_id"`
				MaxLevel     int    `json:"max_level"`
				CurrentLevel int    `json:"current_level"`
				RewardType   string `json:"reward_type"`
				RewardID     int    `json:"reward_id"`
				RewardNum    int    `json:"reward_num"`
				CurrentValue int    `json:"current_value"`
				RewardURL    string `json:"reward_url"`
				Title        string `json:"title"`
				RewardName   string `json:"reward_name"`
			} `json:"single_mission"`
			CompositeMission struct {
				Reward       string `json:"reward"`
				RemainReward int    `json:"remain_reward"`
				TargetValue  int    `json:"target_value"`
				MissionID    int    `json:"mission_id"`
				MaxLevel     int    `json:"max_level"`
				CurrentLevel int    `json:"current_level"`
				RewardType   string `json:"reward_type"`
				RewardID     int    `json:"reward_id"`
				RewardNum    int    `json:"reward_num"`
				CurrentValue int    `json:"current_value"`
				RewardURL    string `json:"reward_url"`
				Title        string `json:"title"`
				RewardName   string `json:"reward_name"`
			} `json:"composite_mission"`
			EndAt             int `json:"end_at"`
			ContinuousMission []struct {
				Reward       string `json:"reward"`
				RemainReward int    `json:"remain_reward"`
				TargetValue  int    `json:"target_value"`
				MissionID    int    `json:"mission_id"`
				MaxLevel     int    `json:"max_level"`
				CurrentLevel int    `json:"current_level"`
				RewardType   string `json:"reward_type"`
				RewardID     int    `json:"reward_id"`
				RewardNum    int    `json:"reward_num"`
				CurrentValue int    `json:"current_value"`
				RewardURL    string `json:"reward_url"`
				Title        string `json:"title"`
				RewardName   string `json:"reward_name"`
			} `json:"continuous_mission"`
			StartAt int `json:"start_at"`
		} `json:"day"`
		Name           string `json:"name"`
		MissionGenreID int    `json:"mission_genre_id"`
		Genre          string `json:"genre"`
	} `json:"genre_list"`
}

//	ミッション・デイリー（昼/夜）の進捗状況を取得する。
//	ログインして実行すること
func ApiMission(
	client *http.Client,
	room_id string,
) (
	pmission *Mission,
	err error,
) {

	//	https://qiita.com/takeru7584/items/f4ba4c31551204279ed2

	//	APIのコーリンgシーケンス
	//	url := "https://www.showroom-live.com/api/mission?room_id=" + room_id

	//	APIのURL
	turl := "https://www.showroom-live.com/api/mission"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立てる
	values := url.Values{}         // url.Valuesオブジェクト生成
	values.Add("room_id", room_id) // key-valueを追加

	// Request を生成する
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		err = fmt.Errorf("http.NewRequest(): %w", err)
		return
	}

	// 組み立てたクエリを生クエリ文字列に変換して設定する
	req.URL.RawQuery = values.Encode()

	// User-Agentを設定する
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

	//	resp.Bodyを使い回せるようにする。
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bufstr := buf.String()

	//	戻り値の格納先を作る
	pmission = new(Mission)

	//	jsonを構造体に格納する。	
	decoder := json.NewDecoder(buf)
	if err = decoder.Decode(pmission); err != nil {
		log.Printf("decoder.Decode(pmission) err: %v", err)
		log.Printf(" room_id= %s", room_id)
		log.Printf("bufstr: %s", bufstr)
		err = fmt.Errorf("decoder.Decode(pmission) err: %v", err)
		return
	}

	return

}
