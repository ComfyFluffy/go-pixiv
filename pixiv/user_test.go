package pixiv

import "testing"

func TestUser(t *testing.T) {
	id := 10489689
	api, err := getTestAPI()
	if err != nil {
		t.Fatal(err)
	}
	_, err = api.User.Detail(id)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(r)
	_, err = api.User.Illusts(id, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = api.User.BookmarkedIllusts(id, RPublic, nil)
	if err == nil {
		t.Fatal("should fail")
	}
	t.Log(err)

	_, err = api.User.BookmarkedIllusts(id, RPublic, nil)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(r2)
}
