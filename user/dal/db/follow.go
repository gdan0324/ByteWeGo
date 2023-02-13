package db

import (
	"context"
	"log"

	"github.com/gdan0324/ByteWeGo/api/pkg/consts"
)

type Follow struct {
	UserId   int64 `json:"user_id"`
	FollowId int64 `json:"follow_id"`
}

func (u *Follow) TableName() string {
	return consts.UserTableName
}

// GetUsers multiple get list of user info
func GetFollow(ctx context.Context, userID int64, followerId int64) (bool, error) {
	follow := Follow{}
	if err := DB.WithContext(ctx).Where(&Follow{UserId: userID, FollowId: followerId}).Find(&follow).Error; err != nil {
		return false, err
	}
	log.Println(follow)
	return true, nil
}
