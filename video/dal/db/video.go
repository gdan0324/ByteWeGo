package db

import (
	"context"
	"github.com/gdan0324/ByteWeGo/api/pkg/consts"
	"github.com/gdan0324/ByteWeGo/user/kitex_gen/userservice"
	"github.com/gdan0324/ByteWeGo/video/kitex_gen/videoservice"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	Id            int64     `json:"video_id" gorm:"column:video_id"`
	UserId        int64     `json:"user_id" gorm:"column:user_id"`
	PlayUrl       string    `json:"play_url" gorm:"column:play_url"`
	CoverUrl      string    `json:"cover_url" gorm:"column:cover_url"`
	FavoriteCount int64     `json:"favorite_count" gorm:"column:favorite_count"`
	CommentCount  int64     `json:"comment_count" gorm:"column:comment_count"`
	Title         string    `json:"title" gorm:"column:title"`
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time"`
}

type VideoInfo struct {
	Id            int64            `json:"video_id"`
	User          userservice.User `json:"user"`
	PlayUrl       string           `json:"play_url"`
	CoverUrl      string           `json:"cover_url"`
	FavoriteCount int64            `json:"favorite_count"`
	CommentCount  int64            `json:"comment_count"`
	IsFavorite    bool             `json:"is_favorite"`
	Title         string           `json:"title"`
}

func (u *Video) TableName() string {
	return consts.VideoTableName
}

// CreateVideo CreateUser create video info
func CreateVideo(ctx context.Context, video *Video) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(video).Error
		if err != nil {
			return nil
		}
		return nil
	})
	return err
}

// GetVideo QueryUser query list of video info
func GetVideo(ctx context.Context, userId int64) ([]*Video, error) {
	var res []*Video
	err := DB.WithContext(ctx).Table("video").
		Where("user_id = ?", userId).
		Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetUser(ctx context.Context, userId int64) (*videoservice.User, error) {
	var res *videoservice.User
	err := DB.WithContext(ctx).Table("user").Where("user_id = ?", userId).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, err
}

func GetFeed(ctx context.Context, latestTime int64) ([]*Video, error) {
	var res []*Video
	err := DB.WithContext(ctx).Table("video").
		Where("create_time < ?", time.UnixMilli(latestTime)).
		Order("create_time desc").Limit(30).
		Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetTime(ctx context.Context, id int64) (time.Time, error) {
	var nextTime time.Time
	err := DB.WithContext(ctx).Table("video v").Select("v.create_time").Where("video_id = ?", id).Find(&nextTime).Error
	if err != nil {
		return time.Now(), err
	}
	return nextTime, nil
}
