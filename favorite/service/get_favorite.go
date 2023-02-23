package service

import (
	"context"
	"github.com/gdan0324/ByteWeGo/api/pkg/errno"
	"github.com/gdan0324/ByteWeGo/api/pkg/jwt"
	"github.com/gdan0324/ByteWeGo/favorite/dal/db"
	"github.com/gdan0324/ByteWeGo/favorite/kitex_gen/favoriteservice"
	"strconv"
)

type GetFavoriteService struct {
	ctx context.Context
}

func NewGetFavoriteService(ctx context.Context) *GetFavoriteService {
	return &GetFavoriteService{ctx: ctx}
}

func (s *GetFavoriteService) GetFavorite(req *favoriteservice.GetFavoriteListRequest) (resp *favoriteservice.GetFavoriteListResponse, err error) {
	userId := req.UserId
	claim, err := jwt.ParseToken(req.Token)
	if err != nil {
		return nil, err
	}
	ui, err := strconv.Atoi(claim["Id"].(string))
	if err != nil {
		return nil, err
	}
	if int64(ui) != userId {
		return nil, errno.AuthorizationFailedErr
	}
	videos := make([]*favoriteservice.Video, 0)
	favorites, err := db.GetFavorite(s.ctx, userId)
	if err != nil {
		return nil, err
	}
	user, err := db.GetUser(s.ctx, userId)
	if err != nil {
		return nil, err
	}
	for i := range favorites {
		videoId := favorites[i].VideoId
		video, err := db.GetVideo(s.ctx, videoId)
		if err != nil {
			return nil, err
		}
		videos = append(videos, &favoriteservice.Video{
			Id:            videoId,
			Author:        user,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    true,
			Title:         video.Title,
		})
	}
	resp = &favoriteservice.GetFavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "success..",
		VideoList:  videos,
	}
	return resp, nil
}
