package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"log"

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
	user, err := db.QueryUser(s.ctx, userName)
	log.Println("-------------")
	if err != nil {
		return 0, err
	}
	if user == nil {
		return 0, errno.AuthorizationFailedErr
	}

	if user.Password != passWord {
		return 0, errno.AuthorizationFailedErr
	}
	return int64(user.Id), nil
}
