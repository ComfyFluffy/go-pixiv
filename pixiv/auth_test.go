package pixiv

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestAuth(t *testing.T) {
	api := New()
	token, ok := os.LookupEnv("PIXIV_REFRESH_TOKEN")
	if !ok {
		t.Fatal("env:PIXIV_REFRESH_TOKEN not set")
	}
	api.SetRefreshToken(token)
	if proxy, ok := os.LookupEnv("PIXIV_PROXY"); ok {
		api.SetProxy(proxy)
	}
	u := &RespUserDetail{}
	api.Sling.Get("https://app-api.pixiv.net/v1/user/detail?user_id=10489689").ReceiveSuccess(u)
	t.Log(u)
	req, err := api.Sling.Get("https://app-api.pixiv.net/v2/illust/comments?illust_id=70161616").Request()
	if err != nil {
		t.Fatal(err)
	}
	// t.Logf("%+v", req)
	resp, err := api.Client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
}
