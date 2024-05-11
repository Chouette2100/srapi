package srapi

import (
	"bytes"
	"fmt"
	"log"
	//	"time"

	"encoding/json"
	"net/http"
	"net/url"
)

type ActivefanRoom struct {
	LevelImageBaseURL string `json:"level_image_base_url"`
	TotalUserCount    int    `json:"total_user_count"`
	ImageVersion      int    `json:"image_version"`
	FanPower          int    `json:"fan_power"`
	IsExcluded        bool   `json:"is_excluded"`
	TitleImageBaseURL string `json:"title_image_base_url"`
	FanName           string `json:"fan_name"`
}

func ApiActivefanRoom(
	client *http.Client,
	room_id string,
	ym string,
) (
	pafr *ActivefanRoom,
	err error,
) {

	turl := "https://www.showroom-live.com/api/active_fan/room"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return
	}

	// クエリを組み立て
	values := url.Values{}         // url.Valuesオブジェクト生成
	values.Add("room_id", room_id) // key-valueを追加
	values.Add("ym", ym)           // key-valueを追加

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

	pafr = new(ActivefanRoom)

	decoder := json.NewDecoder(buf)
	if err = decoder.Decode(pafr); err != nil {
		log.Printf("decoder.Decode(proominfall) err: %v", err)
		log.Printf(" room_id= %s, ym=%s", room_id, ym)
		log.Printf("bufstr: %s", bufstr)
		err = fmt.Errorf("decoder.Decode(pafr) err: %v", err)
		return
	}

	return

}
