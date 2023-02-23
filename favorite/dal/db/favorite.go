package db

import (
	"context"
	"github.com/gdan0324/ByteWeGo/favorite/kitex_gen/favoriteservice"
	"gorm.io/gorm"
	"time"
)

type Favorite struct {
	VideoId    int64     `json:"video_id" gorm:"column:video_id"`
	UserId     int64     `json:"user_id" gorm:"column:user_id"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}

type Video struct {
	VideoId       int64     `json:"video_id"`
	UserId        int64     `json:"user_id"`
	PlayUrl       string    `json:"play_url"`
	CoverUrl      string    `json:"cover_url"`
	Title         string    `json:"title"`
	CreateTime    time.Time `json:"create_time"`
	CommentCount  int64     `json:"comment_count"`
	FavoriteCount int64     `json:"favorite_count"`
}

func NewFavorite(ctx context.Context, videoId int64, userId int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table("favorite").Create(&Favorite{
			VideoId:    videoId,
			UserId:     userId,
			CreateTime: time.Now(),
		}).Error
		if err != nil {
			return err
		}
		res := tx.Table("video").Where("video_id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	return err
}

func DisFavorite(ctx context.Context, videoId int64, userId int64) error {
	err := DB.WithContext(ctx).Table("favorite").Transaction(func(tx *gorm.DB) error {
		favorite := new(Favorite)
		if err := tx.Where("user_id = ? AND video_id = ?", userId, videoId).First(&favorite).Error; err != nil {
			return err
		}
		err := tx.Unscoped().Delete(&Favorite{}, favorite).Error
		if err != nil {
			return err
		}
		res := tx.Table("video").Where("video_id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	return err
}

func GetFavorite(ctx context.Context, userId int64) ([]*Favorite, error) {
	favorites := make([]*Favorite, 0)
	err := DB.WithContext(ctx).Table("favorite").Where("user_id = ?", userId).Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func GetVideo(ctx context.Context, videoId int64) (*Video, error) {
	var video *Video
	err := DB.WithContext(ctx).Table("video").Where("video_id = ?", videoId).Find(&video).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}

func GetUser(ctx context.Context, userId int64) (*favoriteservice.User, error) {
	var user *favoriteservice.User
	err := DB.WithContext(ctx).Table("user").Where("user_id = ?", userId).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}
