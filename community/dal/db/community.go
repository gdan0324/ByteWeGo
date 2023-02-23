package db

import (
	"context"
	"errors"
	"github.com/gdan0324/ByteWeGo/api/pkg/consts"
	"gorm.io/gorm"
	"log"
	"time"
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

type Friend struct {
	UserId   int64 `json:"user_id"`
	FriendId int64 `json:"friend_id"`
}

type Message struct {
	FromUserId int64  `json:"from_user_id"`
	ToUserId   int64  `json:"to_user_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
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
		err := tx.Table("follow").Create(&Follow{
			UserId:   userId,
			FollowId: followId,
		}).Error
		if err != nil {
			return err
		}

		//相互关注将记录添加到friends表（好友表）
		var resbool bool
		resbool, err = GetFollow(ctx, followId, userId)
		if resbool != false {
			err := tx.Table("friends").Create(&Friend{
				UserId:   userId,
				FriendId: followId,
			}).Error
			if err != nil {
				return err
			}

			err = tx.Table("friends").Create(&Friend{
				UserId:   followId,
				FriendId: userId,
			}).Error
			if err != nil {
				return err
			}
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

		//如果user_id和follow_id本来是相互关注状态，在某一方取消关注后，从好友表中删除这条记录
		isfriend := make([]*Friend, 0)
		res := DB.WithContext(ctx).Table("friends").Where("user_id = ? AND friend_id = ?", userId, followId).Find(&isfriend)
		if len(isfriend) == 1 {
			err := tx.Table("friends").Where("user_id = ? AND friend_id = ?", userId, followId).Delete(&Friend{
				UserId:   userId,
				FriendId: followId,
			}).Error
			if err != nil {
				return err
			}

			err = tx.Table("friends").Where("user_id = ? AND friend_id = ?", followId, userId).Delete(&Friend{
				UserId:   followId,
				FriendId: userId,
			}).Error
			if err != nil {
				return err
			}
		}

		res = tx.Table("user").Where("user_id = ?", followId).Update("follower_count", gorm.Expr("follower_count - ?", 1))
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

// MGetFriends GetFriend multiple get list of frienduser info
func MGetFriends(ctx context.Context, userId int64) ([]*FollowInfo, error) {
	res := make([]*FollowInfo, 0)
	err := DB.WithContext(ctx).Table("friends f").
		Select("u.user_id, u.username, u.follow_count, u.follower_count").
		Joins("join user u on u.user_id = f.friend_id").Where("f.user_id = ?", userId).
		Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetLastMessage(ctx context.Context, userId int64, toUserId int64) (string, error) {
	lastmessage := new(Message)
	err := DB.WithContext(ctx).Table("message").Where("from_user_id = ? AND to_user_id = ?", userId, toUserId).Find(&lastmessage)
	log.Println(err)
	content := lastmessage.Content
	return content, nil
}

// NewMessage MessageAction info
func NewMessage(ctx context.Context, userId int64, toUserId int64, content string) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		contentmessage := make([]*Message, 0)
		res := tx.Table("message").Where("from_user_id = ? AND to_user_id = ?", userId, toUserId).Find(&contentmessage)
		if res.Error != nil {
			return res.Error
		}
		if len(contentmessage) != 0 {
			res = tx.Table("message").Where("from_user_id = ? AND to_user_id = ?", userId, toUserId).Update("content", content)
			if res.Error != nil {
				return res.Error
			}
			res2 := tx.Table("message").Where("from_user_id = ? AND to_user_id = ?", userId, toUserId).Update("create_time", time.Now().Format("2006-01-02 15:04:05"))
			if res2.Error != nil {
				return res.Error
			}
		} else {
			err := tx.Table("message").Create(&Message{
				FromUserId: userId,
				ToUserId:   toUserId,
				Content:    content,
				CreateTime: time.Now().Format("2006-01-02 15:04:05"),
			}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
