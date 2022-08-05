/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php


Ver. 0.0.0
Ver. 0.1.0 リスナー情報 UserDetail にエラーに関するメンバーを追加する（作成済みのプログラムに影響はない）
Ver. 1.0.0 戻り値の staus を error 変更する

*/
package srapi

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

//	リスナー情報
type UserDetail struct {
	User_id int    //	ユーザーID＝リスナーの識別子、ログインしていない場合は 0 となる。
	Name    string //	ユーザー名＝リスナー名

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
		log.Printf("url.Parse() returned error %s\n", err.Error())
		return nil, err
	}
	resp, err := client.Get(u.String())
	if err != nil {
		log.Printf("client.Get() returned error %s\n", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	userdetail = new(UserDetail)
	if err := json.NewDecoder(resp.Body).Decode(userdetail); err != nil {
		log.Printf("json.NewDecoder() returned error %s\n", err.Error())
		return nil, err
	}

	//	log.Printf("userdetail: %+v\n", userdetail)

	return userdetail, nil
}
