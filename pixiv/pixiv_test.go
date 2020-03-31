package pixiv

import (
	"errors"
	"os"
)

var testAPI *AppAPI

func getTestAPI() (*AppAPI, error) {
	if testAPI != nil {
		return testAPI, nil
	}
	api := New()
	token, ok := os.LookupEnv("PIXIV_REFRESH_TOKEN")
	if !ok {
		return nil, errors.New("env:PIXIV_REFRESH_TOKEN not set")
	}
	api.SetRefreshToken(token)
	if proxy, ok := os.LookupEnv("PIXIV_PROXY"); ok {
		api.SetProxy(proxy)
	}
	return api, nil
}
