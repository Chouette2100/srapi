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
00AD01	ApiCdnGiftRanking(), ApiCdnUserGiftRanking()で構造体にErrorsを追加する
00AD02	ApiCdnGiftRanking(), ApiCdnUserGiftRanking()の構造体を展開する。
00AE00	ApiCdnGiftRankingContribution()をあらたに作成する
00AE01	ApiCdnGiftRankingContribution()でtype名GrcRankingListをGrcRankingに変更する
00AF00	ApiCdnSeasonAwardRanking()をあらたに作成する
01AA00  ApiRoomEventAndSupport()の構造体を変更し、GetPointByAPI()の引数にBlockIDを追加する
01AH00  GetPointByApi() で レベルイベントの場合は順位に レベル - 10000 を設定する
01AH01	ApiRoomEventAndSupport()の構造体でanyをinterface{}に変更する
01AH02  GetEventBlockRanking()でルーム数が現実のルーム数と異なる場合の処理を追加する
01AJ00  v2版 2.1.0、exsapiとの循環参照を解消する v2.1.1
01AK00  ApiEventAbstraction()を追加する
01AL00  ApiEventRanking(), ApiCurrentUser(), ApiGenrerankingRoomGenre()を追加する
01AM00  ApiEventQuestRooms()を追加する
01AN00  GetEventQuestRooms()をGetEventQuestRoomsByApi()とする。GetEventQuestRoomsByApi()でルーム数のチェックを結合後に行うようにする。
01AN01  GetEventQuestRoomsByApi()でApiEventQuestRooms()の戻り値をチェックし、エラーがあればnilを返すようにする。
200305  ApiRoomEventAndSupport()のJSONを取得するまでを JsonRoomEventAndSupport() として独立させる
200306  init()関数を追加しuseragentを設定する。CrawlContirbutionRanking(), ApiEventRooms()を追加する。

[要確認]
https://www.showroom-live.com/api/event/kvs2510/on_going_events
https://www.showroom-live.com/api/live/onlive_num
https://www.showroom-live.com/api/event/kvs2510/description
https://www.showroom-live.com/api/event/kvs2510/abstraction
https://www.showroom-live.com/api/event/kvs2510/can_apply
https://www.showroom-live.com/api/event/besthits2025/rooms
*/
const Version = "200306"

// ダミーのUser-Agent
// var useragent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36"
// const useragent = "Mozilla/5.0 (X11; Linux x86_64; rv:124.0) Gecko/20100101 Firefox/124.0"
var useragent string

// var mailAddress = "chouette2100@gmail.com"
var mailAddress = "+https://chouette2100.com/disp-bbs"
