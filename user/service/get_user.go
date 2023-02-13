package service

import (
	"context"
	"log"

	"github.com/gdan0324/ByteWeGo/user/dal/db"
	"github.com/gdan0324/ByteWeGo/user/kitex_gen/userservice"
	"github.com/gdan0324/ByteWeGo/user/utils"
)

type GetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// GetUser multiple get list of user info
func (s *GetUserService) GetUser(req *userservice.GetUserRequest) (*userservice.User, error) {
	modelUser, err := db.GetUser(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	claims, err := utils.ParseToken(req.Token)
	if err != nil {
		return nil, err
	}

	log.Println(claims)
	isFollow, err := db.GetFollow(s.ctx, int64(claims["Id"].(float64)), req.UserId)
	if err != nil {
		return nil, err
	}

	user := &userservice.User{
		Id:            modelUser.UserId,
		Name:          modelUser.Username,
		FollowCount:   modelUser.FollowCount,
		FollowerCount: modelUser.FollowerCount,
		IsFollow:      isFollow,
	}
	return user, nil
}
