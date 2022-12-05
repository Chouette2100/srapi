/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php


Ver. 0.0.0
Ver. 0.1.0 リスナー情報 UserDetail にエラーに関するメンバーを追加する（作成済みのプログラムに影響はない）
Ver. 1.0.0 戻り値の staus を error 変更する
Ver. 1.1.0 err に fmt.Errorf("package.Function: %w", err) を設定する。

*/
package srapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

//	リスナー情報
type UserDetail struct {
	User_id         int    //	ユーザーID＝リスナーの識別子、ログインしていない場合は 0 となる。
	Name            string //	ユーザー名＝リスナー名
	Age             int
	Regidence       int
	Description     string
	Image           string
	Birthday        string
	Is_ng_nick_name bool //	？

	//	エラー情報
	//	ログインしている場合は空配列となる模様。
	//	ログインしていない場合は
	//		Error_user_msg:	"Error occured"
	//		Message:		"Not Found"
	//		Code:			1002
	Errors []struct {
		Error_user_msg string //	エラーメッセージ
		Message        string //	エラー内容
		Code           int    //	エラーコード
	}
}

//	（ログインしている）リスナーの情報を取得する
func ApiUserDetail(
	client *http.Client, //	HTTP client
) (
	userdetail *UserDetail, //	リスナー情報
	err error, //	nil: 正常終了
) {

	turl := "https://www.showroom-live.com/api/user/detail"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parsse: %w", err)
		return nil, err
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

	userdetail = new(UserDetail)
	if err := json.NewDecoder(resp.Body).Decode(userdetail); err != nil {
		err = fmt.Errorf("json.NewDecoder: %w", err)
		return nil, err
	}

	//	log.Printf("userdetail: %+v\n", userdetail)

	return userdetail, nil
}
