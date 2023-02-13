package db

import (
	"context"

	"github.com/gdan0324/ByteWeGo/api/pkg/consts"
)

type User struct {
	UserId        int64  `json:"user_id" gorm:"primary_key"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
}

func (u *User) TableName() string {
	return consts.UserTableName
}

// GetUsers multiple get list of user info
func GetUser(ctx context.Context, userID int64) (*User, error) {
	res := User{}

	if err := DB.WithContext(ctx).Where("user_id = ?", userID).Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Create(user).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) (*User, error) {
	res := &User{}
	if err := DB.WithContext(ctx).Where("username = ?", userName).First(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
