package main

import (
	"context"
	favoriteservice "github.com/gdan0324/ByteWeGo/favorite/kitex_gen/favoriteservice"
	"github.com/gdan0324/ByteWeGo/favorite/service"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// DoFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) DoFavorite(ctx context.Context, req *favoriteservice.DoFavoriteRequest) (resp *favoriteservice.DoFavoriteResponse, err error) {
	resp, err = service.NewDoFavoriteService(ctx).DoFavorite(req)
	if err != nil {
		return &favoriteservice.DoFavoriteResponse{
			StatusCode: 1,
			StatusMsg:  "fail...",
		}, err
	}
	return resp, nil
}

// GetFavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) GetFavoriteList(ctx context.Context, req *favoriteservice.GetFavoriteListRequest) (resp *favoriteservice.GetFavoriteListResponse, err error) {
	resp, err = service.NewGetFavoriteService(ctx).GetFavorite(req)
	if err != nil {
		return &favoriteservice.GetFavoriteListResponse{
			StatusCode: 1,
			StatusMsg:  "fail..",
			VideoList:  nil,
		}, err
	}
	return resp, nil
}
