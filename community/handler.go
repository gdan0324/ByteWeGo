package main

import (
	"context"
	"github.com/gdan0324/ByteWeGo/community/kitex_gen/communityservice"
	"github.com/gdan0324/ByteWeGo/community/service"
)

// CommunityServiceImpl implements the last service interface defined in the IDL.
type CommunityServiceImpl struct{}

// GetFollowList implements the CommunityServiceImpl interface.
func (s *CommunityServiceImpl) GetFollowList(ctx context.Context, req *communityservice.GetFollowRequest) (resp *communityservice.GetFollowResponse, err error) {
	// TODO: Your code here...
	follow, err := service.NewMGetFollowService(ctx).MGetFollow(req)
	if err != nil {
		resp = &communityservice.GetFollowResponse{
			StatusCode: 404,
			StatusMsg:  "fail...",
		}
		return resp, err
	}
	resp = &communityservice.GetFollowResponse{
		StatusCode: 200,
		StatusMsg:  "success...",
		UserList:   follow,
	}
	return resp, nil
}

// GetFollowerList implements the CommunityServiceImpl interface.
func (s *CommunityServiceImpl) GetFollowerList(ctx context.Context, req *communityservice.GetFollowerRequest) (resp *communityservice.GetFollowerResponse, err error) {
	// TODO: Your code here...
	follower, err := service.NewMGetFollowerService(ctx).MGetFollower(req)
	if err != nil {
		resp = &communityservice.GetFollowerResponse{
			StatusCode: 404,
			StatusMsg:  "fail...",
		}
		return resp, err
	}
	resp = &communityservice.GetFollowerResponse{
		StatusCode: 200,
		StatusMsg:  "success...",
		UserList:   follower,
	}
	return resp, nil
}

// Follow implements the CommunityServiceImpl interface.
func (s *CommunityServiceImpl) Follow(ctx context.Context, req *communityservice.FollowRequest) (resp *communityservice.FollowResponse, err error) {
	// TODO: Your code here...
	msg, err := service.NewFollowService(ctx).Follow(req)
	if err != nil {
		resp = &communityservice.FollowResponse{
			StatusCode: 404,
			StatusMsg:  msg,
		}
		return resp, err
	}
	resp = &communityservice.FollowResponse{
		StatusCode: 200,
		StatusMsg:  msg,
	}
	return resp, nil
}
