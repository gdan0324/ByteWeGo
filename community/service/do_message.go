package service

import (
	"context"
	"github.com/gdan0324/ByteWeGo/api/pkg/jwt"
	"github.com/gdan0324/ByteWeGo/community/dal/db"
	"github.com/gdan0324/ByteWeGo/community/kitex_gen/communityservice"
	"strconv"
)

type MessageService struct {
	ctx context.Context
}

// NewFollowService new follow-service
func NewMessageService(ctx context.Context) *MessageService {
	return &MessageService{ctx: ctx}
}

// Follow unfollow or follow
func (s *MessageService) MessageAction(req *communityservice.RelationActionRequest) (string, error) {
	actionType := req.ActionType
	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		return "fail", err
	}
	FromUserId, err := strconv.Atoi(claim["Id"].(string))
	toUserId := req.ToUserId
	content := req.Content

	if actionType == 1 {
		err := db.NewMessage(s.ctx, int64(FromUserId), toUserId, content)
		if err != nil {
			return "fail", err
		}
	}

	return "success", nil
}
