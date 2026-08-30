package srapi

import (
        "bytes"
        "log"
        "strings"

        "encoding/json"
        "net/http"
        "net/url"
)

type CommentStatus struct {
        Ua      interface{}             //      不明、カウントの場合は3、コメントの場合は"3"と設定されている。
        Ok      int                             //      1: 正常終了
        Aft     int                             //      不明、2が設定されている
        Comment string                  //      送信した、カウント・コメントがそのまま設定される。
        Room_id int                             //      コメント対象のルームのルーム識別子（6桁以下の整数）
}

//      コメントをポストする。
func ApiLivePostLiveComment(
        client *http.Client,    //      httpクライアント
        comment string,                 //      コメント内容
        csrftoken string,               //      指定したクライアント（の接続）に対するcsrf_token
        liveid string,                  //      配信の識別子（？）、/api/room/status?room_url_key=... 等で取得する。
) (
        cstatus CommentStatus,
) {

        turl := "https://www.showroom-live.com/api/live/post_live_comment"
        u, err := url.Parse(turl)
        if err != nil {
                log.Printf("ApiLivePostLiveComment() url.Parse() returned error %s\n", err.Error())
                return
        }

        // クエリを組み立て
        values := url.Values{} // url.Valuesオブジェクト生成
        //      values.Add([クエリのキー], [値]) // key-valueを追加
        values.Add("live_id", liveid)
        values.Add("comment", comment)
        values.Add("recommend_comment_id", "")
        values.Add("comment_type", "")
        values.Add("is_delay", "0")
        values.Add("csrf_token", csrftoken)

        // Request を生成
        req, err := http.NewRequest("POST", u.String(), strings.NewReader(values.Encode()))
        if err != nil {
                log.Printf("ApiLivePostLiveComment() http.NewRequst() returned error %s\n", err.Error())
                return
        }

        // Content-Typeを設定
        req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

        // User-Agentを設定
        req.Header.Add("User-Agent", useragent)

        //      req.SetBasicAuth([USER], [PASS})

        //      log.Printf("****** req=%+v\n", req)

        // Doメソッドでリクエストを投げる
        // http.Response型のポインタ（とerror）が返ってくる
        resp, err := client.Do(req)
        if err != nil {
                log.Printf("ApiLivePostLiveComment() client.Do() returned error %s\n", err.Error())
                return
        }

        // 関数を抜ける際に接続を切断し、リソースを解放するため必ずresponse.Bodyをcloseします。
        defer resp.Body.Close()

        //      JSONとして読み込めないときに備えてコピーを作っておきます。
        buf := new(bytes.Buffer)
        buf.ReadFrom(resp.Body)
        bufstr := buf.String()

        //      実行結果JSONを読み込みます。
        if err := json.NewDecoder(buf).Decode(&cstatus); err != nil {
                //      if err.Error() != "EOF" {
                        log.Printf("ApiLivePostLiveComment() json.NewDecoder() returned error %s\n", err.Error())
                        //      log.Printf("req=%+v\n", req)
                        log.Printf("buf=\n%s\n", bufstr)
                        return
                //      }
        }

        return

}