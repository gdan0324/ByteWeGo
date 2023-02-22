package service

import (
	"context"

	"github.com/gdan0324/ByteWeGo/video/dal/db"
	"github.com/gdan0324/ByteWeGo/video/kitex_gen/videoservice"
)

type GetVideoService struct {
	ctx context.Context
}

// NewGetVideoService NewMGetVideoService new MGetVideoService
func NewGetVideoService(ctx context.Context) *GetVideoService {
	return &GetVideoService{ctx: ctx}
}

// GetVideo GetUser multiple get list of user info
func (s *GetVideoService) GetVideo(req *videoservice.GetVideoListRequest) ([]*videoservice.Video, error) {
	modelVideo, err := db.GetVideo(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	resp := make([]*videoservice.Video, 0)
	for i := range modelVideo {
		user, err := db.GetUser(s.ctx, modelVideo[i].UserId)
		if err != nil {
			return nil, err
		}
		resp = append(resp, &videoservice.Video{
			Id:            modelVideo[i].Id,
			User:          user,
			PlayUrl:       modelVideo[i].PlayUrl,
			CoverUrl:      modelVideo[i].CoverUrl,
			FavoriteCount: modelVideo[i].FavoriteCount,
			CommentCount:  modelVideo[i].CommentCount,
			IsFavorite:    false,
			Title:         modelVideo[i].Title,
		})
	}
	return resp, nil
}
