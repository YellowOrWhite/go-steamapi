package steamapi

import (
	"net/url"
	"strconv"
	"strings"
)

type playerBansJson struct {
	Players []PlayerBan
}

type PlayerBan struct {
	Steamid         uint64 `json:"SteamId,string"`
	CommunityBanned bool
	VACBanned       bool
	EconomyBan      string
}

var getPlayerBans = NewSteamMethod("ISteamUser", "GetPlayerBans", 1)

func (api *Api) GetPlayerBans(ids []uint64) ([]PlayerBan, error) {
	strIds := make([]string, len(ids))
	for _, id := range ids {
		strIds = append(strIds, strconv.FormatUint(id, 10))
	}

	data := url.Values{}
	data.Add("key", api.GetApiKey())
	data.Add("steamids", strings.Join(strIds, ","))

	var resp playerBansJson
	err := getPlayerBans.Request(data, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Players, nil
}
