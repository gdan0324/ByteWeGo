package service

import (
	"context"
	"github.com/gdan0324/ByteWeGo/api/pkg/errno"
	"github.com/gdan0324/ByteWeGo/comments/dal/db"
	"github.com/gdan0324/ByteWeGo/comments/kitex_gen/commentservice"
)

type GetCommentsService struct {
	ctx context.Context
}

// NewCommentActionService new CommentActionService
func NewGetCommentsService(ctx context.Context) *GetCommentsService {
	return &GetCommentsService{
		ctx: ctx,
	}
}

// CommentList return comment list
func (s *GetCommentsService) GetComments(req *commentservice.GetCommentsRequest, claimID int64) ([]*commentservice.Comment, error) {
	Comments, err := db.GetComments(s.ctx, req.VideoId)
	if err != nil {
		return nil, errno.NewErrNo(20002, "database err")
	}

	comments, err := db.PackComments(s.ctx, Comments, claimID)
	if err != nil {
		return nil, errno.NewErrNo(20002, "database err")
	}
	return comments, nil
}
