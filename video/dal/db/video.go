package db

import (
	"context"
	"github.com/gdan0324/ByteWeGo/api/pkg/consts"
	"github.com/gdan0324/ByteWeGo/api/pkg/errno"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	UserID        int64  `json:"user_id" gorm:"primary_key"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
}

type Video struct {
	VideoID       int64     `json:"video_id" gorm:"primary_key"`
	PlayUrl       string    `json:"play_url"`
	CoverUrl      string    `json:"cover_url"`
	Title         string    `json:"title"`
	CreateTime    time.Time `json:"create_time"`
	CommentCount  int64     `json:"comment_count"`
	FavoriteCount int64     `json:"favorite_count"`
	UserID        int64     `json:"user_id"`
	User          User      `gorm:"foreignkey:UserID"`
}

type Favorite struct {
	VideoID    int64     `json:"video_id" gorm:"primaryKey"`
	UserID     int64     `json:"user_id"`
	CreateTime time.Time `json:"create_time"`
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}

func (f *Favorite) TableName() string {
	return consts.FavoriteTableName
}

// FavoriteAction is used to increment or decrement the favorite count of specified videoID
// according to the action_type
func FavoriteAction(ctx context.Context, favorite *Favorite, actionType int32) (bool, error) {
	var err error
	if actionType == consts.FavoriteType {
		err = DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			err := tx.Create(favorite).Error
			if err != nil {
				return err
			}
			log.Printf("create a favorite to video_id = %d from user %d", favorite.VideoID, favorite.UserID)
			res := tx.Table(consts.VideoTableName).Where("video_id = ?", favorite.VideoID).
				Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
			if res.Error != nil {
				return res.Error
			}
			if res.RowsAffected != 1 {
				return errno.NewErrNo(20002, "database err")
			}
			return nil
		})
	} else if actionType == consts.NoFavoriteType {
		err = DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			err := tx.Delete(favorite).Error
			if err != nil {
				return err
			}
			log.Printf("delete a favorite to video_id = %d from user %d", favorite.VideoID, favorite.UserID)
			res := tx.Table(consts.VideoTableName).Where("video_id = ?", favorite.VideoID).
				Update("favorite_count", gorm.Expr("favorite_count + ?", -1))
			if res.Error != nil {
				return res.Error
			}
			if res.RowsAffected != 1 {
				return errno.NewErrNo(20002, "database err")
			}
			return nil
		})
	}
	if err == nil {
		return true, err
	}
	return false, err
}

// FavoriteList get all favorite video information of specified user according to the userID
func FavoriteList(ctx context.Context, userID int64) ([]*Video, error) {
	videos := make([]*Video, 0)
	err := DB.WithContext(ctx).Table("video v").
		Select("f.user_id, v.*").
		Joins("join favorite f on f.video_id = v.video_id").
		Where("f.user_id = ?", userID).
		Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
