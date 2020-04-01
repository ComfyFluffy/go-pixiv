package pixiv

import (
	"net/url"
	"strconv"
)

// UserService does the fetching with user.
type UserService service

// IllustQuery defines url query struct in fetching user's illusts.
type IllustQuery struct {
	Filter string `url:"filter,omitempty"`
	Type   string `url:"type,omitempty"`
	Offset int    `url:"offset,omitempty"`
}

// BookmarkQuery defines url query struct in fetching bookmark.
type BookmarkQuery struct {
	Filter string `url:"filter,omitempty"` //for_ios
}

// UserDetailQuery defines url query struct in fetching user's detail.
type UserDetailQuery struct {
	Filter string `url:"filter,omitempty"` //for_ios
}

// FollowingQuery defines url query struct in fetching user's followings.
type FollowingQuery struct {
	Restrict Restrict `url:"restrict,omitempty"`
}

// Restrict defines the restrict query field in fetching bookmark.
// It can be "public" or "private".
type Restrict string

// Restrict can be "public" or "private".
const (
	RPublic  Restrict = "public"
	RPrivate Restrict = "private"
	RAll     Restrict = "all"
)

// Detail fetches user profile from /v1/user/detail
func (s *UserService) Detail(userID int, opts *UserDetailQuery) (*RespUserDetail, error) {
	r := &RespUserDetail{}
	err := s.api.getWithValues(r, s.api.BaseURL+"/v1/user/detail", opts, url.Values{
		"user_id": {strconv.Itoa(userID)},
	}, "user detail",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Illusts fetches user's illusts.
func (s *UserService) Illusts(userID int, opts *IllustQuery) (*RespIllusts, error) {
	r := &RespIllusts{api: s.api}
	err := s.api.getWithValues(r, s.api.BaseURL+"/v1/user/illusts", opts, url.Values{
		"user_id": {strconv.Itoa(userID)},
	}, "user's illusts",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// BookmarkedIllusts fetches user's bookmarked illusts.
func (s *UserService) BookmarkedIllusts(userID int, restrict Restrict, opts *BookmarkQuery) (*RespIllusts, error) {
	r := &RespIllusts{api: s.api}
	err := s.api.getWithValues(r, s.api.BaseURL+"/v1/user/bookmarks/illust", opts, url.Values{
		"user_id":  {strconv.Itoa(userID)},
		"restrict": {string(restrict)},
	}, "user's bookmarked illusts",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Novels fetches user's novels.
func (s *UserService) Novels(userID int) (*RespNovels, error) {
	r := &RespNovels{api: s.api}
	err := s.api.getWithValues(r, s.api.BaseURL+"/v1/user/novels", nil, url.Values{
		"user_id": {strconv.Itoa(userID)},
	}, "user's novels",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// BookmarkedNovels fetches user's bookmarked novels.
func (s *UserService) BookmarkedNovels(userID int, restrict Restrict, opts *BookmarkQuery) (*RespNovels, error) {
	r := &RespNovels{api: s.api}
	err := s.api.getWithValues(r, s.api.BaseURL+"/v1/user/bookmarks/novel", opts, url.Values{
		"user_id":  {strconv.Itoa(userID)},
		"restrict": {string(restrict)},
	}, "user's bookmarked novels",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Followings fetches user following with UserPreviews
func (s *UserService) Followings(userID int, opts *FollowingQuery) (*RespUserFollowing, error) {
	r := &RespUserFollowing{api: s.api}
	err := s.api.getWithValues(r, s.api.BaseURL+"/v1/user/following", opts, url.Values{
		"user_id": {strconv.Itoa(userID)},
	}, "user's following",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}
