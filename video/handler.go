package main

import (
	"context"
	videoservice "github.com/gdan0324/ByteWeGo/video/kitex_gen/videoservice"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// FavoriteAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteAction(ctx context.Context, req *videoservice.FavoriteActionRequest) (resp *videoservice.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteList(ctx context.Context, req *videoservice.FavoriteListRequest) (resp *videoservice.FavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}
