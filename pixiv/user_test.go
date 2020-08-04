package pixiv

import (
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	id := 23459386
	api := getTestAPI(t)

	_, err := api.User.Detail(id, nil)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(51 * time.Second)

	ri, err := api.User.Illusts(id, nil)
	_, err = ri.NextIllusts()
	if err != nil {
		t.Fatal(err)
	}

	rn, err := api.User.Novels(id)
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
	_, err = api.User.Recommended(nil)
	_, err = api.User.IllustBookmarkTags(RPublic)
	_, err = api.User.NovelBookmarkTags(RPublic)
	rf, err := api.User.Followings(id, nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = rf.NextFollowing()
	if err != nil {
		t.Fatal(err)
	}
}
