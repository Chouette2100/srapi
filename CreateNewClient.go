/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php
*/
package srapi

import (
	"fmt"
	//	"log"

	"net/http"

	"github.com/juju/persistent-cookiejar"
)

/*
HTTPクライアントを作り、cookiejarをセットする。

**使用上の注意**
	この関数は必ず次のような形で使ってください。
        //      cookiejarがセットされたHTTPクライアントを作る
        client, jar, err := exsrapi.CreateNewClient(sr_acct)
        if err != nil {
                log.Printf("CreateNewClient() returned error %s\n", err.Error())
                return
        }
        //      すべての処理が終了したらcookiejarを保存する。
        defer jar.Save()	//	忘れずに！


Ver.0.0.0
Ver.1.0.0 LoginShowroom()の戻り値 status を err に変更する。
Ver.-.-.- exsrapi.go から分離する。
Ver.1.0.1 errの印字をやめ、エラー内容をerrにセットして返すようにする。
Ver.1.0.2 ログ出力を削除する（Ver.1.0.1では、ログ出力が残っていた）
Ver.1.0.3 logのimportを除く。

*/

//	HTTPクライアントを作り、cookiejarをセットする。
func CreateNewClient(
	cookiename string,
) (
	client *http.Client,
	jar *cookiejar.Jar,
	err error,
) {
	//	Cookiejarを作る
	//	Filenameは、cookieを保存するファイル名
	jar, err = cookiejar.New(&cookiejar.Options{Filename: cookiename + "_cookies"})
	if err != nil {
		//	log.Printf("cookiejar.New() returned error %s\n", err.Error())
		err = fmt.Errorf("cookiejar.New(): %w", err)
		return
	}

	//	httpクライアントを作る
	client = &http.Client{}
	client.Jar = jar

	return
}
