package service

import (
	"context"
	"github.com/gdan0324/ByteWeGo/community/dal/db"
	"github.com/gdan0324/ByteWeGo/community/kitex_gen/communityservice"
	"github.com/gdan0324/ByteWeGo/community/pack"
)

type MGetFriendService struct {
	ctx context.Context
}

// NewMGetFollowerService new MGetFollowerService
func NewMGetFriendService(ctx context.Context) *MGetFriendService {
	return &MGetFriendService{ctx: ctx}
}

// MGetFollower multiple get follower info
func (s *MGetFriendService) MGetFriends(req *communityservice.CheckFriendRequest) ([]*communityservice.FriendUser, error) {

	friends, err := db.MGetFriends(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	isFollow := make([]bool, len(friends))
	for i := range isFollow {
		userId := friends[i].UserId
		res, _ := db.GetFollow(s.ctx, req.UserId, userId)
		isFollow[i] = res
	}
	lastMessage := make([]string, len(friends))
	for i := range lastMessage {
		userId := friends[i].UserId
		res, _ := db.GetLastMessage(s.ctx, req.UserId, userId)
		lastMessage[i] = res
	}
	return pack.Friends(friends, isFollow, lastMessage), nil
}
