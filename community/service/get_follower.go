package service

import (
	"context"
	"github.com/gdan0324/ByteWeGo/community/dal/db"
	"github.com/gdan0324/ByteWeGo/community/kitex_gen/communityservice"
	"github.com/gdan0324/ByteWeGo/community/pack"
)

type MGetFollowerService struct {
	ctx context.Context
}

// NewMGetFollowerService new MGetFollowerService
func NewMGetFollowerService(ctx context.Context) *MGetFollowerService {
	return &MGetFollowerService{ctx: ctx}
}

// MGetFollower multiple get follower info
func (s *MGetFollowerService) MGetFollower(req *communityservice.GetFollowerRequest) ([]*communityservice.User, error) {
	followers, err := db.MGetFollowers(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	isFollow := make([]bool, len(followers))
	for i := range isFollow {
		userId := followers[i].UserId
		res, _ := db.GetFollow(s.ctx, req.UserId, userId)
		isFollow[i] = res
	}
	return pack.Users(followers, isFollow), nil
}
