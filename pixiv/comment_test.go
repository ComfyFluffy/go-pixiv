package pixiv

import "testing"

func TestComment(t *testing.T) {
	api := getTestAPI(t)
	_, err := api.Comment.RepliesNovel(12384655)
	_, err = api.Comment.RepliesIllust(98822844)
	r, err := api.Comment.AddToIllust(69228362, "Hi")
	if err != nil {
		t.Fatal(err)
	}
	err = api.Comment.DeleteFromIllust(r.Comment.ID)

	r2, err := api.Comment.AddToNovel(12632158, "Hi")
	if err != nil {
		t.Fatal(err)
	}
	err = api.Comment.DeleteFromNovel(r2.Comment.ID)

	if err != nil {
		t.Fatal(err)
	}
}
