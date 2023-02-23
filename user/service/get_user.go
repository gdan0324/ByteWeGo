package service

import (
	"context"
	"github.com/gdan0324/ByteWeGo/api/pkg/jwt"
	"log"
	"strconv"

	"github.com/gdan0324/ByteWeGo/user/dal/db"
	"github.com/gdan0324/ByteWeGo/user/kitex_gen/userservice"
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

	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		return nil, err
	}
	userId, err := strconv.Atoi(claims["Id"].(string))
	if err != nil {
		return nil, err
	}
	log.Println(claims)
	isFollow, err := db.GetFollow(s.ctx, int64(userId), req.UserId)
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
