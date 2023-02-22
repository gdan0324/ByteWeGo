package main

import (
	"context"
	"github.com/gdan0324/ByteWeGo/api/pkg/jwt"
	userservice "github.com/gdan0324/ByteWeGo/user/kitex_gen/userservice"
	"github.com/gdan0324/ByteWeGo/user/service"
	"strconv"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userservice.CheckUserRequest) (resp *userservice.CheckUserResponse, err error) {
	id, err := service.NewCheckUserService(ctx).CheckUser(req)

	if err != nil {
		return &userservice.CheckUserResponse{
			StatusCode: int32(userservice.ErrCode_ServiceErrCode),
			StatusMsg:  err.Error(),
		}, nil
	}

	token, _ := jwt.GnerateToken(strconv.FormatInt(id, 10))

	return &userservice.CheckUserResponse{
		StatusCode: 0,
		StatusMsg:  "OK",
		UserId:     id,
		Token:      token,
	}, nil
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userservice.CreateUserRequest) (resp *userservice.CreateUserResponse, err error) {
	id, err := service.NewCreateUserService(ctx).CreateUser(req)

	if err != nil {
		return &userservice.CreateUserResponse{
			StatusCode: int32(userservice.ErrCode_ServiceErrCode),
			StatusMsg:  err.Error(),
		}, nil
	}

	token, _ := jwt.GnerateToken(strconv.FormatInt(id, 10))

	return &userservice.CreateUserResponse{
		StatusCode: 0,
		StatusMsg:  "OK",
		UserId:     id,
		Token:      token,
	}, nil
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *userservice.GetUserRequest) (resp *userservice.GetUserResponse, err error) {
	user, err := service.NewGetUserService(ctx).GetUser(req)

	if err != nil {
		return &userservice.GetUserResponse{
			StatusCode: int32(userservice.ErrCode_ServiceErrCode),
			StatusMsg:  err.Error(),
		}, nil
	}

	return &userservice.GetUserResponse{
		StatusCode: 0,
		StatusMsg:  "OK",
		User:       user,
	}, nil
}
