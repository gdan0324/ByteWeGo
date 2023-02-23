package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gdan0324/ByteWeGo/video/dal/db"
	"github.com/gdan0324/ByteWeGo/video/kitex_gen/videoservice"
	"time"
)

type GetFeedService struct {
	ctx context.Context
}

func NewGetFeedService(ctx context.Context) *GetFeedService {
	return &GetFeedService{ctx: ctx}
}

func (s *GetFeedService) GetFeed(req *videoservice.GetFeedRequest) (vis []*videoservice.Video, nextTime int64, err error) {
	leastTime := req.LatestTime
	if leastTime == 0 {
		leastTime = time.Now().UnixMilli()
	}
	videos, err := db.GetFeed(s.ctx, leastTime)
	klog.Info(len(videos))
	if err != nil {
		return nil, nextTime, err
	}
	if len(videos) == 0 {
		nextTime = time.Now().UnixMilli()
		return vis, nextTime, nil
	} else {
		t, err := db.GetTime(s.ctx, videos[len(videos)-1].Id)
		if err != nil {
			nextTime = time.Now().UnixMilli()
			return vis, nextTime, nil
		}
		nextTime = t.Unix()
	}
	var res []*videoservice.Video
	for i := range videos {
		user, err := db.GetUser(s.ctx, videos[i].UserId)
		if err != nil {
			return nil, nextTime, err
		}
		res = append(res, &videoservice.Video{
			Id:            videos[i].Id,
			User:          user,
			PlayUrl:       videos[i].PlayUrl,
			CoverUrl:      videos[i].CoverUrl,
			FavoriteCount: videos[i].FavoriteCount,
			CommentCount:  videos[i].CommentCount,
			IsFavorite:    false,
			Title:         videos[i].Title,
		})
	}
	klog.Info(len(res))
	return res, nextTime, nil
}
