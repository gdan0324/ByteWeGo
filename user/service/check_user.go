package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/gdan0324/ByteWeGo/api/pkg/errno"
	"github.com/gdan0324/ByteWeGo/user/dal/db"
	"github.com/gdan0324/ByteWeGo/user/kitex_gen/userservice"
)

type CheckUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

// CheckUser check user info
func (s *CheckUserService) CheckUser(req *userservice.CheckUserRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.Username
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}
	u := users[0]
	if u.Password != passWord {
		return 0, errno.AuthorizationFailedErr
	}
	return int64(u.ID), nil
}
