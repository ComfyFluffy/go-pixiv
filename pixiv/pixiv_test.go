package pixiv

import (
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestOAuth(t *testing.T) {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	proxy, ok := os.LookupEnv("PIXIV_PROXY")
	if ok {
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			t.Fatal("PIXIV_PROXY cannot be resolved: ", proxy)
		}
		transport.Proxy = http.ProxyURL(proxyURL)
	}
	api := NewWithClient(&http.Client{
		Transport: transport,
		Timeout:   timeOut,
	})
	token, ok := os.LookupEnv("PIXIV_REFRESH_TOKEN")
	if !ok {
		t.Fatal("PIXIV_REFRESH_TOKEN not set")
	}
	err := api.AuthWithToken(token)
	if err != nil {
		t.Fatal(err)
	}
}
