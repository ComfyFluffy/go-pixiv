package pixiv

import (
	"net/url"
	"strconv"
)

// CommentService fetches comments.
type CommentService service

// RepliesIllust fetches illust comment replies.
func (s *CommentService) RepliesIllust(commentID int) (*RespComments, error) {
	r := &RespComments{api: s.api}
	err := s.api.getWithValues(r, s.api.BaseURL+"/v1/illust/comment/replies", nil, url.Values{
		"comment_id": {strconv.Itoa(commentID)},
	}, "comment replies: illust")
	if err != nil {
		return nil, err
	}
	return r, nil
}

// RepliesNovel fetches novel comment replies.
func (s *CommentService) RepliesNovel(commentID int) (*RespComments, error) {
	r := &RespComments{api: s.api}
	err := s.api.getWithValues(r, s.api.BaseURL+"/v1/novel/comment/replies", nil, url.Values{
		"comment_id": {strconv.Itoa(commentID)},
	}, "comment replies: novel")
	if err != nil {
		return nil, err
	}
	return r, nil
}
