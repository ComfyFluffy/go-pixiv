package pixiv

import (
	"net/url"
	"strconv"
)

// IllustService does ops with illust.
type IllustService service

// AddBookmarkOptions defines form body in AddBookmark.
type AddBookmarkOptions struct {
	Tags []string `url:"tags[],omitempty"`
}

// RelatedQuery defines url query of related illusts.
type RelatedQuery struct {
	Filter string `url:"filter,omitempty"`
}

// NewIllustsQuery defines url query of new illusts from everyone.
type NewIllustsQuery struct {
	ContentType string `url:"content_type,omitempty"`
	Filter      string `url:"filter,omitempty"`
}

// AddBookmark adds illust to public or private bookmark.
func (s *IllustService) AddBookmark(illustID int, restrict Restrict, opts *AddBookmarkOptions) error {
	return s.api.postWithValues(nil,
		s.api.BaseURL+"/v2/illust/bookmark/add",
		opts, url.Values{
			"illust_id": {strconv.Itoa(illustID)},
			"restrict":  {string(restrict)},
		}, "illust: bookmark add",
	)
}

// DeleteBookmark deletes illust from public and private bookmark
func (s *IllustService) DeleteBookmark(illustID int) error {
	return s.api.postWithValues(nil,
		s.api.BaseURL+"/v1/illust/bookmark/delete",
		nil, url.Values{
			"illust_id": {strconv.Itoa(illustID)},
		}, "illust: bookmark add",
	)
}

// AddHistory adds illust browsing history.
func (s *IllustService) AddHistory(illustIDs []int) error {
	return s.api.postWithValues(nil,
		s.api.BaseURL+"/v2/user/browsing-history/illust/add",
		nil, url.Values{
			"illust_ids[]": intsToStrings(illustIDs),
		}, "illust: history add",
	)
}

// Comments fetches comments of the illust.
func (s *IllustService) Comments(illustID int) (*RespComments, error) {
	r := &RespComments{api: s.api}
	err := s.api.getWithValues(r,
		s.api.BaseURL+"/v2/illust/comments",
		nil, url.Values{
			"illust_id": {strconv.Itoa(illustID)},
		}, "illust: comments",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Detail fetches illust's detail by it's id.
//
// The API seems not used in Pixiv's app.
func (s *IllustService) Detail(illustID int) (*RespIllust, error) {
	r := &RespIllust{}
	err := s.api.getWithValues(r,
		s.api.BaseURL+"/v1/illust/detail",
		nil, url.Values{
			"illust_id": {strconv.Itoa(illustID)},
		}, "illust: detail",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Related fetches related illusts.
func (s *IllustService) Related(illustID int, opts *RelatedQuery) (*RespIllusts, error) {
	r := &RespIllusts{api: s.api}
	err := s.api.getWithValues(r,
		s.api.BaseURL+"/v2/illust/related",
		nil, url.Values{
			"illust_id": {strconv.Itoa(illustID)},
		}, "illust: related",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// NewFromFollowings fetches new illusts from followings.
func (s *IllustService) NewFromFollowings(restrict Restrict) (*RespIllusts, error) {
	r := &RespIllusts{api: s.api}
	err := s.api.getWithValues(r,
		s.api.BaseURL+"/v2/illust/follow",
		nil, url.Values{
			"restrict": {string(restrict)},
		}, "illust: new from followings",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// NewFromAll fetches new illusts from everyone.
func (s *IllustService) NewFromAll(opts *NewIllustsQuery) (*RespIllusts, error) {
	r := &RespIllusts{api: s.api}
	err := s.api.getWithValues(r,
		s.api.BaseURL+"/v1/illust/new",
		opts, nil, "illust: new from all",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// NewFromMyPixiv fetches new illusts from my-pixiv.
func (s *IllustService) NewFromMyPixiv() (*RespIllusts, error) {
	r := &RespIllusts{api: s.api}
	err := s.api.getWithValues(r,
		s.api.BaseURL+"/v2/illust/mypixiv",
		nil, nil, "illust: new from following",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// UgoiraMetadata fetches ugoira metadata.
func (s *IllustService) UgoiraMetadata(illustID int) (*RespUgoiraMetadata, error) {
	r := &RespUgoiraMetadata{}
	err := s.api.getWithValues(r,
		s.api.BaseURL+"/v1/ugoira/metadata", nil, url.Values{
			"illust_id": {strconv.Itoa(illustID)},
		}, "illust: ugoira metadata",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}
