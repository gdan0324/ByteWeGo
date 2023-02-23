package service

import (
	"context"
	"github.com/gdan0324/ByteWeGo/video/kitex_gen/videoservice"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

// FavoriteAction todo complete the service
func (s *FavoriteActionService) FavoriteAction(req *videoservice.FavoriteActionRequest) (bool, error) {
	return false, nil
}
