package pixiv

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/google/go-querystring/query"
)

// UserService runs the fetching about user.
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

// NovelQuery defines url query struct in fetching novels.
type NovelQuery struct{}

// Restrict defines the restrict query field in fetching bookmark.
// It can be "public" or "private".
type Restrict string

// Restrict can be "public" or "private".
const (
	RPublic  Restrict = "public"
	RPrivate Restrict = "private"
	RAll     Restrict = "all"
)

func (u *UserService) withValues(r interface{}, opts interface{}, values url.Values, urls string, caller string) error {
	var err error
	var q url.Values

	if opts != nil {
		q, err = query.Values(opts)
		if err != nil {
			return fmt.Errorf("pixiv: user %s: query encode: %w", caller, err)
		}
		for k, v := range values {
			q[k] = v
		}
	} else {
		q = values
	}

	_, err = u.api.get(urls, q, r)
	if err != nil {
		return err
	}

	return nil
}

// Detail fetches user profile from /v1/user/detail
func (u *UserService) Detail(userID int, opts *UserDetailQuery) (*RespUserDetail, error) {
	r := &RespUserDetail{}
	err := u.withValues(r, opts, url.Values{
		"user_id": []string{strconv.Itoa(userID)}},
		u.api.BaseURL+"/v1/user/detail", "user detail")
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Illusts fetches user's illusts.
func (u *UserService) Illusts(userID int, opts *IllustQuery) (*RespIllusts, error) {
	r := &RespIllusts{api: u.api}
	err := u.withValues(r, opts, url.Values{
		"user_id": []string{strconv.Itoa(userID)}},
		u.api.BaseURL+"/v1/user/illusts", "user's illusts")
	if err != nil {
		return nil, err
	}
	return r, nil
}

// BookmarkedIllusts fetches user's bookmarked illusts.
func (u *UserService) BookmarkedIllusts(userID int, restrict Restrict, opts *BookmarkQuery) (*RespIllusts, error) {
	r := &RespIllusts{api: u.api}
	err := u.withValues(r, opts, url.Values{
		"user_id":  []string{strconv.Itoa(userID)},
		"restrict": []string{string(restrict)}},
		u.api.BaseURL+"/v1/user/bookmarks/illust", "bookmarked illusts")
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Novels fetches user's novels.
func (u *UserService) Novels(userID int, opts *NovelQuery) (*RespNovels, error) {
	r := &RespNovels{api: u.api}
	err := u.withValues(r, opts, url.Values{
		"user_id": []string{strconv.Itoa(userID)}},
		u.api.BaseURL+"/v1/user/novels", "novels")
	if err != nil {
		return nil, err
	}
	return r, nil
}

// BookmarkedNovels fetches user's bookmarked novels.
func (u *UserService) BookmarkedNovels(userID int, restrict Restrict, opts *BookmarkQuery) (*RespNovels, error) {
	r := &RespNovels{api: u.api}
	err := u.withValues(r, opts, url.Values{
		"user_id":  []string{strconv.Itoa(userID)},
		"restrict": []string{string(restrict)}},
		u.api.BaseURL+"/v1/user/bookmarks/novel", "user's novels")
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Following fetches user following with UserPreviews
func (u *UserService) Following(userID int, opts *FollowingQuery) (*RespUserFollowing, error) {
	r := &RespUserFollowing{api: u.api}
	err := u.withValues(r, opts, url.Values{
		"user_id": []string{strconv.Itoa(userID)}},
		u.api.BaseURL+"/v1/user/following", "following")
	if err != nil {
		return nil, err
	}
	return r, nil
}
