package pixiv

import "testing"

func TestUser(t *testing.T) {
	id := 23459386
	api, err := getTestAPI()
	if err != nil {
		t.Fatal(err)
	}

	_, err = api.User.Detail(id)
	if err != nil {
		t.Fatal(err)
	}

	_, err = api.User.Illusts(id, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = api.User.Novels(id, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = api.User.BookmarkedIllusts(id, RPublic, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = api.User.BookmarkedNovels(id, RPublic, nil)
	if err != nil {
		t.Fatal(err)
	}

	r, err := api.User.Following(id, nil)
	if err != nil {
		t.Fatal(err)
	}
	r2, err := r.NextUserPreviews()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r2)
}
