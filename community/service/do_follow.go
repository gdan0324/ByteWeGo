package service

import (
	"context"
	"github.com/gdan0324/ByteWeGo/api/pkg/jwt"
	"github.com/gdan0324/ByteWeGo/community/dal/db"
	"github.com/gdan0324/ByteWeGo/community/kitex_gen/communityservice"
	"strconv"
)

type FollowService struct {
	ctx context.Context
}

// NewFollowService new follow-service
func NewFollowService(ctx context.Context) *FollowService {
	return &FollowService{ctx: ctx}
}

// Follow unfollow or follow
func (s *FollowService) Follow(req *communityservice.FollowRequest) (string, error) {
	actionType := req.ActionType
	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		return "fail", err
	}
	userId, err := strconv.Atoi(claim["Id"].(string))
	if err != nil {
		return "fail", err
	}
	followId := req.ToUserId
	if actionType == 1 {
		err := db.NewFollow(s.ctx, int64(userId), followId)
		if err != nil {
			return "fail", err
		}
	}

	if actionType == 2 {
		err := db.DisFollow(s.ctx, int64(userId), followId)
		if err != nil {
			return "fail", err
		}
	}
	return "success", nil
}
