package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gdan0324/ByteWeGo/api/pkg/errno"
	"github.com/gdan0324/ByteWeGo/user/dal/db"
	"github.com/gdan0324/ByteWeGo/user/kitex_gen/userservice"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *userservice.CreateUserRequest) (int64, error) {
	log.Println("CreateUser")
	user, err := db.QueryUser(s.ctx, req.Username)
	if err != nil && err.Error() != "record not found" {
		return 0, err
	}

	if user != nil && user.UserId != 0 {
		return 0, errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))

	user = &db.User{
		Username:   req.Username,
		Password:   password,
		CreateTime: time.Now(),
	}
	err = db.CreateUser(s.ctx, user)
	if err != nil {
		return 0, err
	}

	return user.UserId, nil
}
