package pixiv

import (
	"net/url"
	"strconv"
)

// NovelService does ops with novels.
type NovelService service

// AddHistory adds novel browsing history.
func (s *NovelService) AddHistory(novelIDs []int) error {
	return s.api.postWithValues(nil,
		s.api.BaseURL+"/v2/user/browsing-history/novel/add",
		nil, url.Values{
			"novel_ids[]": intsToStrings(novelIDs),
		}, "novel: add history",
	)
}

// AddBookmark adds novel to public or private bookmark.
func (s *NovelService) AddBookmark(novelID int, restrict Restrict, opts *AddBookmarkOptions) error {
	return s.api.postWithValues(nil,
		s.api.BaseURL+"/v2/novel/bookmark/add",
		opts, url.Values{
			"novel_id": {strconv.Itoa(novelID)},
			"restrict": {string(restrict)},
		}, "novel: bookmark add",
	)
}

// DeleteBookmark deletes novel from public and private bookmark
func (s *NovelService) DeleteBookmark(novelID int) error {
	return s.api.postWithValues(nil,
		s.api.BaseURL+"/v1/novel/bookmark/delete",
		nil, url.Values{
			"novel_id": {strconv.Itoa(novelID)},
		}, "novel: bookmark add",
	)
}

// Text fetches text of the novel.
func (s *NovelService) Text(novelID int) (*RespNovelText, error) {
	r := &RespNovelText{}
	err := s.api.getWithValues(r,
		s.api.BaseURL+"/v1/novel/text",
		nil, url.Values{
			"novel_id": {strconv.Itoa(novelID)},
		}, "novel: text",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Comments fetches comments of the novel.
func (s *NovelService) Comments(novelID int) (*RespComments, error) {
	r := &RespComments{api: s.api}
	err := s.api.getWithValues(r,
		s.api.BaseURL+"/v2/novel/comments",
		nil, url.Values{
			"novel_id": {strconv.Itoa(novelID)},
		}, "novel: comments",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Detail fetches novel's detail by it's id.
//
// The API seems not used in Pixiv's app.
func (s *NovelService) Detail(novelID int) (*RespNovel, error) {
	r := &RespNovel{}
	err := s.api.getWithValues(r,
		s.api.BaseURL+"/v2/novel/detail",
		nil, url.Values{
			"novel_id": {strconv.Itoa(novelID)},
		}, "novel: detail",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}
