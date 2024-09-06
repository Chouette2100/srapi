package srapi
/*
00AA00	v1.3.0に対し
		ApiRoomProfile()をApiRoomProfileV111()とする。
		ApiRoomProfileAll()をApiRoomRoomProfile()とする。
		ApiMission()を追加する。
		srapi.goを追加する(Versionおよびuseragent)
00AA01	ApiEventRanking()のイベント終了後の動作を調べコメントに記す。
00AA02	EventRankingのメンバーにErrorsを追加する。ApiRoomProfile()でDecodeのエラー出力を簡略化する。
00AA03	ApiRoomProfile()でエラー出力を行わないようにする。
00AA04	ApiEventsRanking() の event_block_id を event_block_division_id とする
00AB00	ApiEventBlockRanking()を複数ページに対応させる。
		ApiEventRoomList_test.goとApiEventBlockRanking.goを作成する。
		importのgithub.com/chouette2100をChouette2100に変更する。
00AC00	ApiEventGiftRanking()をあらたに作成する
00AD00	ApiCdnGiftRanking(), ApiCdnUserGiftRanking()をあらたに作成する

*/
const Version = "00AD00"

//	ダミーのUser-Agent
//	var useragent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36"
const useragent = "Mozilla/5.0 (X11; Linux x86_64; rv:124.0) Gecko/20100101 Firefox/124.0"

