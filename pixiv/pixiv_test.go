package pixiv

import (
	"errors"
	"os"
	"testing"
)

var testAPI *AppAPI

func getTestAPI(t *testing.T) *AppAPI {
	if testAPI != nil {
		return testAPI
	}
	api := New()
	token, ok := os.LookupEnv("PIXIV_REFRESH_TOKEN")
	if !ok {
		t.Fatal(errors.New("env:PIXIV_REFRESH_TOKEN not set"))
		return nil
	}
	api.SetRefreshToken(token)
	if proxy, ok := os.LookupEnv("PIXIV_PROXY"); ok {
		api.SetProxy(proxy)
	}
	return api
}
