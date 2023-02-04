package db

import (
	"context"

	"github.com/gdan0324/ByteWeGo/api/pkg/consts"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string `json:"username"`
	Password      string `json:"password"`
	FollowCount   string `json:"follow_count"`
	FollowerCount string `json:"follower_count"`
}

func (u *User) TableName() string {
	return consts.UserTableName
}

// GetUsers multiple get list of user info
func GetUser(ctx context.Context, userID int64) (*User, error) {
	res := &User{}

	if err := DB.WithContext(ctx).Where("id = ?", userID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
