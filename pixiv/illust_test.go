package pixiv

import "testing"

func TestBookmarkOps(t *testing.T) {
	id := 80486549
	api, err := getTestAPI()
	if err != nil {
		t.Fatal(err)
	}

	err = api.Illust.DeleteBookmark(id)
	if err != nil && err.(*ErrAppAPI).response.StatusCode != 404 {
		t.Fatal(err)
	}

	err = api.Illust.AddBookmark(id, RPublic, &AddBookmarkOptions{
		Tags: []string{"ショタ", "正太", "test"},
	})
	err = api.Illust.AddHistory([]int{id})
	_, err = api.Illust.Comments(id)
	_, err = api.Illust.Detail(id)
	_, err = api.Illust.NewFromAll(nil)
	_, err = api.Illust.NewFromFollowings(RPublic)
	_, err = api.Illust.NewFromMyPixiv()
	_, err = api.Illust.Related(id, nil)
	if err != nil {
		t.Fatal(err)
	}
}
