package srapi
/*
00AA00	v1.3.0に対し
		ApiRoomProfile()をApiRoomProfileV111()とする。
		ApiRoomProfileAll()をApiRoomRoomProfile()とする。
		ApiMission()を追加する。
		srapi.goを追加する(Versionおよびuseragent)
00AA01	ApiEventRanking()のイベント終了後の動作を調べコメントに記す。
00AA02	EventRankingのメンバーにErrorsを追加する。ApiRoomProfile()でDecodeのエラー出力を簡略化する。
*/
const Version = "00AA02"

//	ダミーのUser-Agent
//	var useragent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36"
const useragent = "Mozilla/5.0 (X11; Linux x86_64; rv:124.0) Gecko/20100101 Firefox/124.0"

