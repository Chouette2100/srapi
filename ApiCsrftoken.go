/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

Ver. 0.0.0 srapi.goから分離する。
Ver. 1.0.0 戻り値 status を err に変更する。
Ver. 1.1.1	ダミーのユーザーエージェントを FireFox/124.0に変更する。

*/
package srapi

import (
	"fmt"

	"encoding/json"
	"net/http"
	"net/url"
)

//	ダミーのUser-Agent
/*
	var useragent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36"
*/
var useragent = "Mozilla/5.0 (X11; Linux x86_64; rv:124.0) Gecko/20100101 Firefox/124.0"

//	csrf_token ... 接続の識別子
type CsrfToken struct {
	Csrf_token string
	Data1      string
}

//	csrf_tokenを取得する。
func ApiCsrftoken(
	client *http.Client, //	HTTPクライアント
) (
	csrftoken string, //	csrf_token
	err error, //	エラー
) {

	turl := "https://www.showroom-live.com/api/csrf_token"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse: %w", err)
		return csrftoken, err
	}

	// クエリを組み立て
	values := url.Values{} // url.Valuesオブジェクト生成
	//	values.Add([クエリのキー], [値]) // key-valueを追加

	// Request を生成
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		err = fmt.Errorf("http.NewRequest: %w", err)
		return csrftoken, err
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
		err = fmt.Errorf("client.Do: %w", err)
		return csrftoken, err
	}

	// 関数を抜ける際に接続を切断し、リソースを解放するため必ずresponse.Bodyをcloseする
	defer resp.Body.Close()

	var ct CsrfToken
	if err = json.NewDecoder(resp.Body).Decode(&ct); err != nil {
		err = fmt.Errorf("json.NewDecoder: %w", err)
		return csrftoken, err
	}
	csrftoken = ct.Csrf_token

	return csrftoken, nil

}
