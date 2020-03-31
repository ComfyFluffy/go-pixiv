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
	Filter string `url:"filter,omitempty"`
}

// Restrict defines the restrict query field in fetching bookmark.
// It can be "public" or "private".
type Restrict string

// Restrict can be "public" or "private".
const (
	RPublic  Restrict = "public"
	RPrivate Restrict = "private"
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
func (u *UserService) Detail(userID int) (*RespUserDetail, error) {
	r := &RespUserDetail{}
	_, err := u.api.get("/v1/user/detail?user_id="+strconv.Itoa(userID), nil, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Illusts fetches user's illusts from /v1/user/illusts
func (u *UserService) Illusts(userID int, opts *IllustQuery) (*RespIllusts, error) {
	r := &RespIllusts{}
	err := u.withValues(r, opts, url.Values{
		"user_id": []string{strconv.Itoa(userID)}},
		"/v1/user/illusts", "illusts")
	if err != nil {
		return nil, err
	}
	return r, nil
}

// BookmarkedIllusts fetches user's bookmark from /v1/user/bookmarks/illust
func (u *UserService) BookmarkedIllusts(userID int, restrict Restrict, opts *BookmarkQuery) (*RespIllusts, error) {
	r := &RespIllusts{}
	err := u.withValues(r, opts, url.Values{
		"user_id":  []string{strconv.Itoa(userID)},
		"restrict": []string{string(restrict)}},
		"/v1/user/bookmarks/illust", "bookmarked illusts")
	if err != nil {
		return nil, err
	}
	return r, nil
}
