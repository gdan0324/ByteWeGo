package main

import (
	"context"
	userservice "github.com/gdan0324/ByteWeGo/user/kitex_gen/userservice"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userservice.CheckUserRequest) (resp *userservice.CheckUserResponse, err error) {
	// TODO: Your code here...
	return
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userservice.CreateUserRequest) (resp *userservice.CreateUserResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *userservice.GetUserRequest) (resp *userservice.GetUserResponse, err error) {
	// TODO: Your code here...
	return
}
