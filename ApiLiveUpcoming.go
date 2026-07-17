// Copyright © 2026 chouette.21.00@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php

package srapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type LiveUpcoming struct {
	Upcomings []LiveUpcomingRoom `json:"upcomings"`
}

type LiveUpcomingRoom struct {
	IsFollow          bool   `json:"is_follow"`
	RoomURLKey        string `json:"room_url_key"`
	NextLiveStartAt   int    `json:"next_live_start_at"`
	OfficialLv        int    `json:"official_lv"`
	HasFanRoom        bool   `json:"has_fan_room"`
	Description       string `json:"description"`
	LiverBadge        int    `json:"liver_badge"`
	Image             string `json:"image"`
	MainName          string `json:"main_name"`
	NewDaysCountLabel string `json:"new_days_count_label"`
	LiverThemeTitle   string `json:"liver_theme_title"`
	ImageSquare       string `json:"image_square,omitempty"`
	CellType          int    `json:"cell_type"`
	IsFanRoomOnline   bool   `json:"is_fan_room_online"`
	RoomID            int    `json:"room_id"`
}

// 指定したジャンルの配信予定ルーム一覧を取得する。
func ApiLiveUpcoming(
	client *http.Client, // HTTPクライアント
	genreID int, // ジャンルID
) (
	liveupcoming *LiveUpcoming,
	err error,
) {

	turl := "https://www.showroom-live.com/api/live/upcoming"
	u, err := url.Parse(turl)
	if err != nil {
		err = fmt.Errorf("url.Parse(): %w", err)
		return nil, err
	}

	values := url.Values{}
	values.Add("genre_id", fmt.Sprintf("%d", genreID))

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		err = fmt.Errorf("http.NewRequest(): %w", err)
		return nil, err
	}

	req.URL.RawQuery = values.Encode()
	req.Header.Add("User-Agent", useragent)
	req.Header.Add("Accept-Language", "ja-JP")

	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("client.Do(): %w", err)
		return nil, err
	}
	defer resp.Body.Close()

	liveupcoming = new(LiveUpcoming)
	if err = json.NewDecoder(resp.Body).Decode(liveupcoming); err != nil {
		err = fmt.Errorf("json.NewDecoder(resp.Body).Decode(liveupcoming): %w", err)
		return nil, err
	}

	return liveupcoming, nil
}
