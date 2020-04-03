package pixiv

import "net/url"

// SearchService searches pixiv content.
type SearchService service

// TrendingTagsQuery defines url query for IllustTrendingTags.
type TrendingTagsQuery struct {
	Filter string `url:"filter,omitempty"`
}

// SearchTarget defines search_target field in SearchQuery
type SearchTarget string

// SearchTarget values
const (
	STPartialMatchTags SearchTarget = "partial_match_for_tags"
	STExactMatchTags   SearchTarget = "exact_match_for_tags"

	// For novel search only.
	STText SearchTarget = "text"

	// For illust search only.
	STTitleCaption SearchTarget = "title_and_caption"
)

// Sort defines sort field in SearchQuery
type Sort string

// Sort values
const (
	SDateAsc     Sort = "date_asc"
	SDateDesc    Sort = "date_desc"
	SPopularDesc Sort = "popular_desc"
)

// SearchQuery defines url query in illust and novel searching
type SearchQuery struct {
	SearchTarget SearchTarget `url:"search_target,omitempty"`
	Sort         Sort         `url:"sort,omitempty"`
	// MergePlainKeywordResults bool         `url:"merge_plain_keyword_results,omitempty"`
	Filter string `url:"filter,omitempty"`

	StartDate Date `url:"start_date,omitempty"`
	EndDate   Date `url:"end_date,omitempty"`
	Offset    int  `url:"offset,omitempty"`
}

// SearchUserQuery defines url query struct used in user searching
type SearchUserQuery struct {
	Filter string `url:"filter,omitempty"`
	Offset int    `url:"offset,omitempty"`
}

// IllustTrendingTags fetches trending tags of illusts and manga.
func (s *SearchService) IllustTrendingTags(opts *TrendingTagsQuery) (*RespTrendingTags, error) {
	r := &RespTrendingTags{}
	err := s.api.getWithValues(r,
		s.api.BaseURL+"/v1/trending-tags/illust",
		opts, nil, "search: illust trending-tags",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// NovelTrendingTags fetches trending tags of novels.
func (s *SearchService) NovelTrendingTags(opts *TrendingTagsQuery) (*RespTrendingTags, error) {
	r := &RespTrendingTags{}
	err := s.api.getWithValues(r,
		s.api.BaseURL+"/v1/trending-tags/novel",
		opts, nil, "search: novel trending-tags",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *SearchService) illusts(urls, word string, opts *SearchQuery, caller string) (*RespIllusts, error) {
	r := &RespIllusts{api: s.api}
	err := s.api.getWithValues(r,
		s.api.BaseURL+urls, opts, url.Values{
			"word":                           {word},
			"include_translated_tag_results": {"true"},
			"merge_plain_keyword_results":    {"true"},
		}, "search: "+caller,
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Illusts searches illusts with options.
func (s *SearchService) Illusts(word string, opts *SearchQuery) (*RespIllusts, error) {
	return s.illusts("/v1/search/illust", word, opts, "illusts")
}

// PopularIllustsPreview searches 30 illusts sort by popularity
func (s *SearchService) PopularIllustsPreview(word string, opts *SearchQuery) (*RespIllusts, error) {
	// copy opts and clear sort field
	opts2 := *opts
	opts2.Sort = ""
	return s.illusts("/v1/search/popular-preview/illust", word, &opts2, "illusts popular preview")
}

func (s *SearchService) novels(ep, word string, opts *SearchQuery, caller string) (*RespNovels, error) {
	r := &RespNovels{api: s.api}
	err := s.api.getWithValues(r,
		s.api.BaseURL+ep, opts, url.Values{
			"word":                           {word},
			"include_translated_tag_results": {"true"},
			"merge_plain_keyword_results":    {"true"},
		}, "search: "+caller,
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Novels searches novels with options.
func (s *SearchService) Novels(word string, opts *SearchQuery) (*RespNovels, error) {
	return s.novels("/v1/search/novel", word, opts, "novels")
}

// PopularNovelsPreview searches 30 novels sort by popularity
func (s *SearchService) PopularNovelsPreview(word string, opts *SearchQuery) (*RespNovels, error) {
	opts2 := *opts
	opts2.Sort = ""
	return s.novels("/v1/search/popular-preview/novel", word, &opts2, "novels popular preview")
}

// TagsStartWith fetches tags start with word.
func (s *SearchService) TagsStartWith(word string) (*RespTags, error) {
	r := &RespTags{}
	err := s.api.getWithValues(r,
		s.api.BaseURL+"/v2/search/autocomplete", nil, url.Values{
			"word":                        {word},
			"merge_plain_keyword_results": {"true"},
		}, "search: tag autocomplete",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Users searches user previews by options.
func (s *SearchService) Users(word string, opts *SearchUserQuery) (*RespUserPreviews, error) {
	r := &RespUserPreviews{api: s.api}
	err := s.api.getWithValues(r,
		s.api.BaseURL+"/v1/search/user", opts, url.Values{
			"word": {word},
		}, "search: user",
	)
	if err != nil {
		return nil, err
	}
	return r, err
}
