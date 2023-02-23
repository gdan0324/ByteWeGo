package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gdan0324/ByteWeGo/api/pkg/consts"
	"github.com/gdan0324/ByteWeGo/api/pkg/errno"
	"github.com/gdan0324/ByteWeGo/comments/kitex_gen/commentservice"
	"github.com/gdan0324/ByteWeGo/user/dal/db"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID              int       `gorm:"primaryKey;column:comment_id"`
	VideoID         int       `gorm:"column:video_id;not null"`
	User            db.User   `gorm:"foreignkey:UserID"`
	UserID          int       `gorm:"column:user_id;not null"`
	Comment_Content string    `gorm:"type:varchar(512);not null;column:comment_content"`
	Create_date     time.Time `gorm:"column:create_time"`
}

func (Comment) TableName() string {
	return "comment"
}

type User struct {
	ID            int    `gorm:"primaryKey;column:user_id"`
	Username      string `gorm:"column:username"`
	FollowCount   int64  `gorm:"colunm:follow_count"`
	FollowerCount int64  `gorm:"colunm:follower_count"`
	IsFollow      bool
}

func (User) TableName() string {
	return consts.UserTableName
}

func getTableVideoName() string {
	return "video"
}

// NewComment creates a new Comment
func NewComment(ctx context.Context, comment *Comment, claimID int64) (resultcom *commentservice.Comment, err error) {
	err = DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(comment).Error
		if err != nil {
			return err
		}
		res := tx.Table(getTableVideoName()).Where("video_id = ?", comment.VideoID).Update("comment_count", gorm.Expr("comment_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.NewErrNo(20002, "database err")
		}
		return nil
	})
	resultcom, err = PackOneComment(ctx, comment, claimID)
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
	err := DB.WithContext(ctx).Model(&Comment{}).Where(&Comment{VideoID: int(vid)}).Order("create_time desc").Find(&comments).Error
	if err != nil {
		return nil, errno.NewErrNo(20002, "database err")
	}
	return comments, nil
}

func GetUser(ctx context.Context, uid int64, cid int64) (*User, error) {
	user := User{}
	if err := DB.WithContext(ctx).Where("user_id = ?", uid).Find(&user).Error; err != nil {
		return nil, errno.NewErrNo(20002, "database err")
	}
	follow := db.Follow{}
	if err := DB.WithContext(ctx).Table(consts.FollowTableName).Where("user_id = ? AND follow_id = ?", cid, uid).Find(&follow).Error; err != nil {
		return nil, err
	}
	user.IsFollow = follow.UserId != 0
	return &user, nil
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
	user, err := GetUser(ctx, int64(v.UserID), claimID)
	klog.Info("user = ", user)
	if err != nil {
		return nil, err
	}
	return &commentservice.Comment{
		Id: int64(v.ID),
		User: &commentservice.User{
			Id:            int64(user.ID),
			Name:          user.Username,
			FollowCount:   user.FollowerCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		},
		CommentText: v.Comment_Content,
		CommentDate: v.Create_date.Format("2006/1/2-15:04"),
	}, nil
}
