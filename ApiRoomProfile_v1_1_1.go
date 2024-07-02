package srapi
import (
	"time"
	"log"
	"bytes"
	"fmt"

	"encoding/json"
	"net/http"
	"net/url"

)
type RoomInf struct {
	Name      string //     ルーム名のリスト
	Longname  string
	Shortname string
	Account   string //     アカウントのリスト、アカウントは配信のURLの最後の部分の英数字です。
	ID        string //     IDのリスト、IDはプロフィールのURLの最後の部分で5～6桁の数字です。
	Userno    int
	//      APIで取得できるデータ(1)
	Genre      string
	Rank       string
	Irank      int
	Nrank      int
	Prank      int
	Followers  int
	Sfollowers string
	Fans       int
	Fans_lst   int
	Level      int
	Slevel     string
	//      APIで取得できるデータ(2)
	Order        int
	Point        int //     イベント終了後12時間〜36時間はイベントページから取得できることもある
	Spoint       string
	Istarget     string
	Graph        string
	Iscntrbpoint string
	Color        string
	Colorvalue   string
	//	Colorinflist ColorInfList
	Formid       string
	Eventid      string
	Status       string
	Statuscolor  string
}

func ApiRoomProfileV111(
	client *http.Client,
	room_id string,
	) (
	roominf	RoomInf,
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
        values := url.Values{}			// url.Valuesオブジェクト生成
        values.Add("room_id", room_id)		// key-valueを追加

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

	//	JSONをデコードする。
	//	次の記事を参考にさせていただいております。
	//		Go言語でJSONに泣かないためのコーディングパターン
	//		https://qiita.com/msh5/items/dc524e38073ed8e3831b

	var result interface{}
	decoder := json.NewDecoder(buf)
	if err = decoder.Decode(&result); err != nil {
		log.Printf("decoder.Decode(&result) err: %v", err)
		log.Printf(" room?id= %s", room_id)
		log.Printf("bufstr: %s", bufstr)
		err = fmt.Errorf("decoder.Decode(&result) err: %v", err)
		return
	}

	//	フォロワー数
	value, _ := result.(map[string]interface{})["follower_num"].(float64)
	roominf.Followers = int(value)

	tnow := time.Now()
	//	roominf.Fans = GetAciveFanByAPI(room_id, tnow.Format("200601"))
	yy := tnow.Year()
	mm := tnow.Month() - 1
	if mm < 0 {
		yy -= 1
		mm = 12
	}
	//	fans_lst = GetAciveFanByAPI(room_id, fmt.Sprintf("%04d%02d", yy, mm))

	roominf.Genre, _ = result.(map[string]interface{})["genre_name"].(string)

	rank, _ := result.(map[string]interface{})["league_label"].(string)
	ranks, _ := result.(map[string]interface{})["show_rank_subdivided"].(string)
	roominf.Rank = rank + " | " + ranks

	value, _ = result.(map[string]interface{})["next_score"].(float64)
	roominf.Nrank = int(value)
	value, _ = result.(map[string]interface{})["prev_score"].(float64)
	roominf.Prank = int(value)

	value, _ = result.(map[string]interface{})["room_level"].(float64)
	roominf.Level = int(value)

	roominf.Name, _ = result.(map[string]interface{})["room_name"].(string)

	roominf.Account, _ = result.(map[string]interface{})["room_url_key"].(string)

	//	配信開始時刻の取得
	//	value, _ = result.(map[string]interface{})["current_live_started_at"].(float64)
	//	startedat = time.Unix(int64(value), 0).Truncate(time.Second)
	//	log.Printf("current_live_stared_at %f %v\n", value, startedat)

	return

}
