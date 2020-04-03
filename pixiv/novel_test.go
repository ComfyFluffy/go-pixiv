package pixiv

import "testing"

func TestNovel(t *testing.T) {
	id := 12525505
	api := getTestAPI(t)
	api.Novel.DeleteBookmark(id)
	err := api.Novel.AddBookmark(id, RPublic, &AddBookmarkOptions{
		Tags: []string{"ショタ", "正太", "test"},
	})
	err = api.Novel.AddHistory([]int{id})
	_, err = api.Novel.Comments(id)
	_, err = api.Novel.Detail(id)
	_, err = api.Novel.Text(id)
	_, err = api.Novel.Recommended(nil)
	_, err = api.Novel.Ranking(nil)

	if err != nil {
		t.Fatal(err)
	}
}
