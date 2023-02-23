package service

import (
	"context"
	"github.com/gdan0324/ByteWeGo/api/pkg/jwt"
	"github.com/gdan0324/ByteWeGo/favorite/dal/db"
	"github.com/gdan0324/ByteWeGo/favorite/kitex_gen/favoriteservice"
	"strconv"
)

type DoFavoriteService struct {
	ctx context.Context
}

func NewDoFavoriteService(ctx context.Context) *DoFavoriteService {
	return &DoFavoriteService{ctx: ctx}
}

func (s *DoFavoriteService) DoFavorite(req *favoriteservice.DoFavoriteRequest) (resp *favoriteservice.DoFavoriteResponse, err error) {
	videoId := req.VideoId
	actionType := req.ActionType
	token := req.Token
	claim, err := jwt.ParseToken(token)
	if err != nil {
		return nil, err
	}
	userId, err := strconv.Atoi(claim["Id"].(string))
	if err != nil {
		return nil, err
	}
	if actionType == 1 {
		err = db.NewFavorite(s.ctx, videoId, int64(userId))
		if err != nil {
			return nil, err
		}
	}
	if actionType == 2 {
		err = db.DisFavorite(s.ctx, videoId, int64(userId))
		if err != nil {
			return nil, err
		}
	}
	resp = &favoriteservice.DoFavoriteResponse{
		StatusCode: 0,
		StatusMsg:  "success...",
	}
	return resp, nil
}
