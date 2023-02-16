package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/gdan0324/ByteWeGo/api/pkg/errno"
	"github.com/gdan0324/ByteWeGo/comments/kitex_gen/commentservice"
	"github.com/gdan0324/ByteWeGo/user/dal/db"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID int `gorm:"primaryKey;column:comment_id"`
	//Video           Video     `gorm:"foreignkey:VideoID"`
	VideoID         int       `gorm:"column:video_id;not null"`
	User            db.User   `gorm:"foreignkey:UserID"`
	UserID          int       `gorm:"column:user_id;not null"`
	Comment_Content string    `gorm:"type:varchar(512);not null;column:comment_content"`
	Create_date     time.Time `gorm:"column:create_time"`
}

func (Comment) TableName() string {
	return "comment"
}

func getTableVideoName() string {
	return "video"
}

// NewComment creates a new Comment
func NewComment(ctx context.Context, comment *Comment, claimID int64) (*commentservice.Comment, error) {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(comment).Error
		if err != nil {
			return err
		}
		fmt.Printf(" new comment in db , comment_id = ", comment.ID)
		res := tx.Table(getTableVideoName()).Where("video_id = ?", comment.VideoID).Update("comment_count", gorm.Expr("comment_count + ?", 1))

		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.NewErrNo(20002, "database err")
		}
		return nil
	})
	resultcom, err := PackOneComment(ctx, comment, claimID)
	return resultcom, err
}

// DelComment deletes a comment from the database.
func DelComment(ctx context.Context, commentID int64, vid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		comment := new(Comment)
		if err := tx.First(&comment, commentID).Error; err != nil {
			return errno.NewErrNo(20002, "database err")
		}
		err := tx.Unscoped().Delete(&comment).Error
		if err != nil {
			return errno.NewErrNo(20002, "database err")
		}

		res := tx.Table(getTableVideoName()).Where("video_id= ?", vid).Update("comment_count", gorm.Expr("comment_count - ?", 1))

		if res.Error != nil {
			return errno.NewErrNo(20002, "database err")
		}

		if res.RowsAffected != 1 {
			return errno.NewErrNo(20002, "database err")
		}

		return nil
	})
	return err
}

// GetComments  return list<comment>
func GetComments(ctx context.Context, vid int64) ([]*Comment, error) {
	var comments []*Comment
	err := DB.WithContext(ctx).Model(&Comment{}).Where(&Comment{VideoID: int(vid)}).Find(&comments).Error
	if err != nil {
		return nil, errno.NewErrNo(20002, "database err")
	}
	return comments, nil
}

// Comment pack Comments info.
func PackComments(ctx context.Context, vs []*Comment, claimID int64) ([]*commentservice.Comment, error) {
	comments := make([]*commentservice.Comment, 0)
	for _, v := range vs {
		comment, err := PackOneComment(ctx, v, claimID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func PackOneComment(ctx context.Context, v *Comment, claimID int64) (*commentservice.Comment, error) {
	user, err := db.GetUser(ctx, int64((v.UserID)))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	isFollow, err := db.GetFollow(ctx, claimID, int64(v.UserID))
	if err != nil {
		return nil, err
	}
	//packUser, err
	//comments = append(comments, &commentservice.Comment{
	//	Id: int64(v.ID),
	//	User: &commentservice.User{
	//		Id:            user.UserId,
	//		Name:          user.Username,
	//		FollowCount:   user.FollowerCount,
	//		FollowerCount: user.FollowerCount,
	//		IsFollow:      isFollow,
	//	},
	//	CommentText: v.Comment_Content,
	//	CommentDate: v.Create_date.Format("2006/1/2-15:04"),
	//})

	return &commentservice.Comment{
		Id: int64(v.ID),
		User: &commentservice.User{
			Id:            user.UserId,
			Name:          user.Username,
			FollowCount:   user.FollowerCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      isFollow,
		},
		CommentText: v.Comment_Content,
		CommentDate: v.Create_date.Format("2006/1/2-15:04"),
	}, nil
}

//func GetUser(ctx context.Context, uid int64) (*User, error) {
//	var user *User
//
//	if err := DB.WithContext(ctx).First(&user, uid).Error; err != nil {
//		return nil, err
//	}
//	return user, nil
//}
