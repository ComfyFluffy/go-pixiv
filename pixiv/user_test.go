package pixiv

import "testing"

func TestUser(t *testing.T) {
	id := 23459386
	api, err := getTestAPI()
	if err != nil {
		t.Fatal(err)
	}

	_, err = api.User.Detail(id, nil)
	if err != nil {
		t.Fatal(err)
	}

	ri, err := api.User.Illusts(id, nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = ri.NextIllusts()
	if err != nil {
		t.Fatal(err)
	}

	rn, err := api.User.Novels(id, nil)
	if err != nil {
		t.Fatal(err)
	}
	for rn.NextURL != "" {
		rn, err = rn.NextNovels()
		if err != nil {
			t.Fatal(err)
		}
	}

	_, err = api.User.BookmarkedIllusts(id, RPublic, nil)
	if err != nil {
		t.Fatal(err)
	}

	rbn, err := api.User.BookmarkedNovels(id, RPublic, nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = rbn.NextNovels()
	if err != nil {
		t.Fatal(err)
	}

	rf, err := api.User.Following(id, nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = rf.NextFollowing()
	if err != nil {
		t.Fatal(err)
	}
}
