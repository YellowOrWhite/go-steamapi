// Package steamapi provides an interface to the
// Steam Web API methods.
package steamapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

const steamUrl = "http://api.steampowered.com/%v/%v/v%v/"

// Api represents a Steam Web API handler
type Api struct {
	apiKey      string
	apiKeyMutex *sync.Mutex
}

// Create new Steam Web API handler
func NewSteamApi(apiKey string) (*Api, error) {
	api := &Api{apiKey: apiKey, apiKeyMutex: &sync.Mutex{}}
	return api, nil
}

func (api *Api) ChangeApiKey(apiKey string) error {
	api.apiKeyMutex.Lock()
	api.apiKey = apiKey
	api.apiKeyMutex.Unlock()
	return nil
}

// A SteamMethod represents a Steam Web API method.
type SteamMethod string

// Creates a new SteamMethod.
func NewSteamMethod(interf, method string, version int) SteamMethod {
	m := fmt.Sprintf(steamUrl, interf, method, strconv.Itoa(version))
	return SteamMethod(m)
}

// Makes a request to the Steam Web API with the given
// url values and stores the result in v.
//
// Returns an error if the return status code was not 200.
func (s SteamMethod) Request(data url.Values, v interface{}) error {
	url := string(s)
	if data != nil {
		url += "?" + data.Encode()
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Status code %d returned by %v", resp.StatusCode, string(s)))
	}

	d := json.NewDecoder(resp.Body)
	return d.Decode(&v)
}
