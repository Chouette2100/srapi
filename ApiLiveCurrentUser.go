/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php


Ver. 0.1.0 ApiLiveCurrentUser.goをsrapi.goから分離する。ApiLiveCurrentUser()のRoomIDをstring型に変更する。
Ver. 1.0.0 戻り値 status を err に変更する。

*/
package srapi

import (
	"fmt"

	"encoding/json"
	"net/http"
	"net/url"
)

//	リスナー情報
type LiveCurrentUser struct {
	Is_login           bool   //	ログインしているか? ログインしていればtrue
	User_id            int    //	リスナーの識別子
	Name               string //	リスナー名
	Fan_level          int	  //	ファンレベル、Maxが56のやつ
	Add_free_gift      int	  //	？
	Contribution_point int	  //	おそらく貢献ポイントの累計
	Gift_list          struct {	//	手持ちのギフトの数
		Normal []struct {
			Gift_id  int	//	ギフトの種別 1であれば赤星、...
			Free_num int	//	Gift_idのギフトの個数
		}
	}
}

//	リスナーに関する情報を取得する
func ApiLiveCurrentUser(
	client *http.Client, //	HTTP client
	roomid string, //	配信者ID
) (
	livecurrentuser LiveCurrentUser, //	配信者ルームにおけるリスナーの情報
	err error, //	エラー情報
) {

	turl := "https://www.showroom-live.com/api/live/current_user?room_id=" + roomid
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parsse: %w", err)
		return livecurrentuser, err
	}

	// Request を生成
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		err = fmt.Errorf("http.NewRequest(): %w", err)
		return
	}

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

	if err = json.NewDecoder(resp.Body).Decode(&livecurrentuser); err != nil {
		err = fmt.Errorf("json.NewDecoder: %w", err)
		return livecurrentuser, err
	}

	return livecurrentuser, nil
}
