package pixiv

import "testing"

func TestSearch(t *testing.T) {
	api := getTestAPI(t)
	_, err := api.Search.IllustTrendingTags(nil)
	_, err = api.Search.NovelTrendingTags(nil)
	_, err = api.Search.Illusts("ショタ", &SearchQuery{
		SearchTarget: STExactMatchTags,
		Sort:         SDateDesc,
	})
	_, err = api.Search.PopularIllustsPreview("ショタ", &SearchQuery{
		SearchTarget: STExactMatchTags,
	})
	_, err = api.Search.Novels("ショタ", &SearchQuery{
		SearchTarget: STExactMatchTags,
		Sort:         SDateDesc,
	})
	_, err = api.Search.PopularNovelsPreview("ショタ", &SearchQuery{
		SearchTarget: STExactMatchTags,
	})
	_, err = api.Search.TagsStartWith("シ")

	if err != nil {
		t.Fatal(err)
	}
}
