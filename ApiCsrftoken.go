/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

Ver. 0.0.0 srapi.goから分離する。

*/
package srapi

import (
	"log"

	"encoding/json"
	"net/http"
	"net/url"
)

//	ダミーのUser-Agent
var useragent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36"

//	csrf_token ... 接続の識別子
type CsrfToken struct {
	Csrf_token string
	Data1      string
}

//	csrf_tokenを取得する。
func ApiCsrftoken(client *http.Client) (csrftoken string) {

	turl := "https://www.showroom-live.com/api/csrf_token"
	u, err := url.Parse(turl)
	if err != nil {
		log.Printf("url.Parse() returned error %s\n", err.Error())
		return
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成
	//	values.Add([クエリのキー], [値]) // key-valueを追加

	// Request を生成
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Printf("http.NewRequst() returned error %s\n", err.Error())
		return
	}

	// 組み立てたクエリを生クエリ文字列に変換して設定
	req.URL.RawQuery = values.Encode()

	// User-Agentを設定
	req.Header.Add("User-Agent", useragent)

	// Doメソッドでリクエストを投げる
	// http.Response型のポインタ（とerror）が返ってくる
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("client.Do() returned error %s\n", err.Error())
		return
	}

	// 関数を抜ける際に接続を切断し、リソースを解放するため必ずresponse.Bodyをcloseする
	defer resp.Body.Close()

	var ct CsrfToken
	if err := json.NewDecoder(resp.Body).Decode(&ct); err != nil {
		log.Printf("json.NewDecoder() returned error %s\n", err.Error())
		return
	}
	csrftoken = ct.Csrf_token

	return

}
