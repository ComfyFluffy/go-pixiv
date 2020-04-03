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
	}, "comment: replies illust")
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
	}, "comment: replies novel")
	if err != nil {
		return nil, err
	}
	return r, nil
}

// AddToIllust adds comment to illust.
func (s *CommentService) AddToIllust(illustID int, comment string) (*RespComment, error) {
	r := &RespComment{}
	err := s.api.postWithValues(r,
		s.api.BaseURL+"/v1/illust/comment/add", nil, url.Values{
			"illust_id": {strconv.Itoa(illustID)},
			"comment":   {comment},
		}, "comment: add to illust",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// AddToNovel adds comment to novel.
func (s *CommentService) AddToNovel(novelID int, comment string) (*RespComment, error) {
	r := &RespComment{}
	err := s.api.postWithValues(r,
		s.api.BaseURL+"/v1/novel/comment/add", nil, url.Values{
			"novel_id": {strconv.Itoa(novelID)},
			"comment":  {comment},
		}, "comment: add to novel",
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// DeleteFromIllust deletes illust comment by id.
func (s *CommentService) DeleteFromIllust(commentID int) error {
	return s.api.postWithValues(nil,
		s.api.BaseURL+"/v1/illust/comment/delete", nil, url.Values{
			"comment_id": {strconv.Itoa(commentID)},
		}, "comment: delete from illust",
	)
}

// DeleteFromNovel deletes novel comment by id.
func (s *CommentService) DeleteFromNovel(commentID int) error {
	return s.api.postWithValues(nil,
		s.api.BaseURL+"/v1/novel/comment/delete", nil, url.Values{
			"comment_id": {strconv.Itoa(commentID)},
		}, "comment: delete from novel",
	)
}
