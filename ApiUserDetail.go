/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php


Ver. 0.0.0

*/
package srapi

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type UserDetail struct {
	User_id	int	//	ユーザーID＝リスナーの識別子
	Name 	string	//	ユーザー名＝リスナー名
}

//	（ログインしている）リスナーの情報を取得する
func ApiUserDetail(
	client *http.Client, //	HTTP client
) (
	userdetail *UserDetail, //	配信中ルームのジャンル別一覧
	status int, //	0: 正常終了
) {

	status = 0

	turl := "https://www.showroom-live.com/api/user/detail"
	u, err := url.Parse(turl)
	if err != nil {
		log.Printf("url.Parse() returned error %s\n", err.Error())
		status = -1
		return
	}
	resp, err := client.Get(u.String())
	if err != nil {
		log.Printf("client.Get() returned error %s\n", err.Error())
		status = -2
		return
	}
	defer resp.Body.Close()

	userdetail = new(UserDetail)
	if err := json.NewDecoder(resp.Body).Decode(userdetail); err != nil {
		log.Printf("json.NewDecoder() returned error %s\n", err.Error())
		status = -3
		return
	}

	return
}
