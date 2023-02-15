package service

import (
	"context"
	"github.com/gdan0324/ByteWeGo/community/dal/db"
	"github.com/gdan0324/ByteWeGo/community/kitex_gen/communityservice"
	"github.com/gdan0324/ByteWeGo/community/pack"
)

type MGetFollowService struct {
	ctx context.Context
}

// NewMGetFollowService new MGetFollowService
func NewMGetFollowService(ctx context.Context) *MGetFollowService {
	return &MGetFollowService{ctx: ctx}
}

// MGetFollow multiple get follow info
func (s *MGetFollowService) MGetFollow(req *communityservice.GetFollowRequest) ([]*communityservice.User, error) {
	follows, err := db.MGetFollows(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	isFollow := make([]bool, len(follows))
	for i := range isFollow {
		isFollow[i] = true
	}
	return pack.Users(follows, isFollow), nil
}
