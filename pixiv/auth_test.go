package pixiv

import (
	"io/ioutil"
	"testing"
)

func TestAuth(t *testing.T) {

	// u := &RespUserDetail{}
	// api.Sling.Get("https://app-api.pixiv.net/v1/user/detail?user_id=10489689").ReceiveSuccess(u)
	// t.Log(u)
	api, err := getTestAPI()
	if err != nil {
		t.Fatal(err)
	}

	req, err := api.NewRequest("GET", "https://app-api.pixiv.net/v2/illust/comments?illust_id=70161616", nil)
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
