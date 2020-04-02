package pixiv

import "testing"

func TestComment(t *testing.T) {
	api := getTestAPI(t)
	_, err := api.Comment.RepliesNovel(12384655)
	_, err = api.Comment.RepliesIllust(98822844)
	if err != nil {
		t.Fatal(err)
	}
}
