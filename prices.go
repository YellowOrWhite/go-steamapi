package steamapi

import (
	"errors"
	"net/url"
	"strconv"
)

type storeJson struct {
	Result struct {
		Success bool
		Assets  []Asset
	}
}

// An item in the store.
type Asset struct {
	Prices   map[string]int
	Defindex int `json:"name,string"`
	Date     string
	Tags     []string
	TagIds   []int64
}

func (i *Asset) HasTag(tag string) bool {
	for _, t := range i.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

var getAssetPrices = NewSteamMethod("ISteamEconomy", "GetAssetPrices", 1)

func (api *Api) GetAssetPrices(appid uint64, language, currency string) ([]Asset, error) {
	vals := url.Values{}
	vals.Add("key", api.apiKey)
	vals.Add("appid", strconv.FormatUint(appid, 10))
	vals.Add("language", language)
	vals.Add("currency", currency)

	var resp storeJson
	err := getAssetPrices.Request(vals, &resp)
	if err != nil {
		return nil, err
	}
	if !resp.Result.Success {
		return nil, errors.New("API call 'GetAssetPrices' did not succeed!")
	}
	return resp.Result.Assets, nil
}
