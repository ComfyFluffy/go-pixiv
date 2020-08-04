package pixiv

import (
	"errors"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

var testAPI *AppAPI

func setProxy(tr *http.Transport, uri string) error {
	if uri == "none" {
		return nil
	}
	pr, err := url.Parse(uri)
	if err != nil {
		return err
	}

	switch strings.ToLower(pr.Scheme) {
	case "http":
		hp := http.ProxyURL(pr)
		tr.Proxy = hp
	default:
		return errors.New("set proxy: unsupported protocol")
	}
	return nil
}

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
	proxy, ok := os.LookupEnv("PIXIV_PROXY")
	if ok {
		setProxy(api.Client.Transport.(*http.Transport), proxy)
	}
	api.SetRefreshToken(token)
	testAPI = api
	return api
}
