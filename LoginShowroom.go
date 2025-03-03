/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php
*/
package srapi

import (
	"fmt"
	"log"

	"net/http"
)

/*

Ver.0.0.0
Ver.1.0.0 LoginShowroom()の戻り値 status を err に変更する。
Ver.-.-.- exsrapi.go から分離する。

*/

//	Showroomのサービスにログインし、ユーザIDを取得する。
func LoginShowroom(
	client *http.Client,
	acct string,
	pswd string,
) (
	userid string,
	err error,
) {
	//	SHOWROOMにログインした状態にあるか？
	ud, err := ApiUserDetail(client)
	if err != nil {
		log.Printf(" err = %+v\n", err)
		return "0", err
	}
	//	log.Printf("----------------------------------------------------\n")
	//	log.Printf("%+v\n", ud)
	//	log.Printf("----------------------------------------------------\n")

	if ud.User_id == 0 {
		//	ログインしていない

		//	csrftokenを取得する
		csrftoken, err := ApiCsrftoken(client)
		if err != nil {
			err = fmt.Errorf("ApiCsrftoken: %w", err)
			return "0", err
		}

		//	SHOWROOMのサービスにログインする。
		var ul UserLogin
		ul, err = ApiUserLogin(client, csrftoken, acct, pswd)
		if err != nil {
			err = fmt.Errorf("ApiUserLogin: %w", err)
			return "0", err
		} else {
			log.Printf("login status. Ok = %d User_id=%s\n", ul.Ok, ul.User_id)
		}
		userid = ul.User_id

	} else {
		//      ログインしている
		userid = fmt.Sprintf("%d", ud.User_id)
	}
	return userid, nil

}
