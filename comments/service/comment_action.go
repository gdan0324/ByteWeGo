package service

import (
	"context"
	"github.com/gdan0324/ByteWeGo/api/pkg/errno"
	"github.com/gdan0324/ByteWeGo/comments/dal/db"
	"github.com/gdan0324/ByteWeGo/comments/kitex_gen/commentservice"
	"time"
)

type CommentActionService struct {
	ctx context.Context
}

func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

func (s *CommentActionService) CommentAction(req *commentservice.CommentActionRequest, claimID int64) (comment *commentservice.Comment, err error) {
	if req.ActionType == 1 {
		//return db.NewComment(s.ctx, &db.Comment{
		//	UserID:          int(req.UserId),
		//	VideoID:         int(req.VideoId),
		//	Comment_Content: req.CommentText,
		//	Create_date:     time.Now(),
		//})
		comment, err := db.NewComment(s.ctx, &db.Comment{
			UserID:          int(req.UserId),
			VideoID:         int(req.VideoId),
			Comment_Content: req.CommentText,
			Create_date:     time.Now(),
		}, claimID)

		if err != nil {
			return nil, err
		}
		return comment, nil
	}

	if req.ActionType == 2 {
		return nil, db.DelComment(s.ctx, req.CommentId, req.VideoId)
	}
	return nil, errno.NewErrNo(20001, "err param")
}
