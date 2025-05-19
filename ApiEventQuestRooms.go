package srapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "net/url"
	// "strconv"
	// "strings"
)

type EventQuestRooms struct {
	EventQuestLevelRanges []EventQuestLevelRanges `json:"event_quest_level_ranges"`
	IsPointVisiblePeriod  bool                    `json:"is_point_visible_period"`
	TotalEntries          int                     `json:"total_entries"`
}
type Rooms struct {
	QuestLevel      int    `json:"quest_level"`
	RoomID          int    `json:"room_id"`
	RoomImage       string `json:"room_image"`
	RoomName        string `json:"room_name"`
	RoomURLKey      string `json:"room_url_key"`
	RoomDescription string `json:"room_description"`
	IsOnLive        bool   `json:"is_on_live"`
	IsFollowing     bool   `json:"is_following"`
	IsOfficial      bool   `json:"is_official"`
	Point           int    `json:"point"`
}
type EventQuestLevelRanges struct {
	LevelFrom int     `json:"level_from"`
	LevelTo   int     `json:"level_to"`
	Rooms     []Rooms `json:"rooms"`
}

func ApiEventQuestRooms(
	client *http.Client,
	eventid string,
) (eqr *EventQuestRooms, err error) {
	eqr = new(EventQuestRooms)
	eqr.EventQuestLevelRanges = make([]EventQuestLevelRanges, 0)

	// URLの生成
	u := fmt.Sprintf("https://www.showroom-live.com/api/event/%s/quest_rooms", eventid)
	//fmt.Println(u)

	// GETリクエストを作成
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	// リクエストを送信
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get event quest rooms: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(eqr)
	if err != nil {
		return nil, err
	}

	return eqr, nil
}
func GetEventQuestRooms(
	client *http.Client,
	eventid string,
	ib int,
	ie int,
) (eqr *EventQuestRooms, err error) {

	eqr = new(EventQuestRooms)

	eqr.EventQuestLevelRanges = make([]EventQuestLevelRanges, 0)
	eqr.EventQuestLevelRanges = append(eqr.EventQuestLevelRanges, EventQuestLevelRanges{
		LevelFrom: 0,
		LevelTo:   0,
		Rooms:     make([]Rooms, 0),
	})
	eqr.EventQuestLevelRanges[0].Rooms = make([]Rooms, 0)
	eqr.EventQuestLevelRanges[0].Rooms = append(eqr.EventQuestLevelRanges[0].Rooms, Rooms{
		QuestLevel:      0,
		RoomID:          0,
		RoomImage:       "",
		RoomName:        "",
		RoomURLKey:      "",
		RoomDescription: "",
		IsOnLive:        false,
		IsFollowing:     false,
		IsOfficial:      false,
		Point:           0,
	})

	eqr, err = ApiEventQuestRooms(client, eventid)

	for i, eqlr := range eqr.EventQuestLevelRanges {
		if i != 0 && len(eqlr.Rooms) < ie {
			eqr.EventQuestLevelRanges[0].Rooms = append(eqr.EventQuestLevelRanges[0].Rooms, eqlr.Rooms...)
		}
	}
	if len(eqr.EventQuestLevelRanges[0].Rooms) < ie {
		ie = len(eqr.EventQuestLevelRanges[0].Rooms)
	}
	eqr.EventQuestLevelRanges[0].Rooms = eqr.EventQuestLevelRanges[0].Rooms[ib-1 : ie-ib+1]

	return eqr, nil
}
