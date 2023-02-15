package db

import (
	"context"
	"errors"
	"github.com/gdan0324/ByteWeGo/api/pkg/consts"
	"gorm.io/gorm"
)

type Follow struct {
	UserId   int64 `json:"user_id"`
	FollowId int64 `json:"follow_id"`
}

type FollowInfo struct {
	UserId        int64  `gorm:"column:user_id"`
	UserName      string `gorm:"column:username"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
}

func (f *Follow) TableName() string {
	return consts.FollowTableName
}

func GetFollow(ctx context.Context, userId int64, followId int64) (bool, error) {
	follow := make([]*Follow, 0)
	res := DB.WithContext(ctx).Where("user_id = ? AND follow_id = ?", userId, followId).Find(&follow)
	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected > 1 {
		return false, errors.New("database error")
	}
	return res.RowsAffected == 1, nil
}

func MGetFollows(ctx context.Context, userId int64) ([]*FollowInfo, error) {
	res := make([]*FollowInfo, 0)
	err := DB.WithContext(ctx).Table("follow f").
		Select("u.user_id, u.username, u.follow_count, u.follower_count").
		Joins("join user u on u.user_id = f.follow_id").Where("f.user_id = ?", userId).
		Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func MGetFollowers(ctx context.Context, userId int64) ([]*FollowInfo, error) {
	res := make([]*FollowInfo, 0)
	err := DB.WithContext(ctx).Table("follow f").
		Select("u.user_id, u.username, u.follow_count, u.follower_count").
		Joins("join user u on u.user_id = f.user_id").Where("f.follow_id = ?", userId).
		Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewFollow(ctx context.Context, userId int64, followId int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&Follow{
			UserId:   userId,
			FollowId: followId,
		}).Error
		if err != nil {
			return err
		}

		res := tx.Table("user").Where("user_id = ?", followId).Update("follower_count", gorm.Expr("follower_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return errors.New("database error")
		}

		res = tx.Table("user").Where("user_id = ?", userId).Update("follow_count", gorm.Expr("follow_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return errors.New("database error")
		}
		return nil
	})
	return err
}

func DisFollow(ctx context.Context, userId int64, followId int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		follow := new(Follow)
		if err := tx.Where("user_id = ? AND follow_id = ?", userId, followId).First(&follow).Error; err != nil {
			return err
		}
		err := tx.Unscoped().Delete(&Follow{}, follow).Error
		if err != nil {
			return err
		}
		res := tx.Table("user").Where("user_id = ?", followId).Update("follower_count", gorm.Expr("follower_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return errors.New("database error")
		}

		res = tx.Table("user").Where("user_id = ?", userId).Update("follow_count", gorm.Expr("follow_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return errors.New("database error")
		}
		return nil
	})
	return err
}
