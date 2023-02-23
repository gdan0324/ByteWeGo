package main

import (
	"context"
	"github.com/gdan0324/ByteWeGo/video/service"
	//"github.com/gdan0324/ByteWeGo/user/utils"
	"github.com/gdan0324/ByteWeGo/video/kitex_gen/videoservice"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, req *videoservice.CreateVideoRequest) (resp *videoservice.CreateVideoResponse, err error) {
	// TODO: Your code here...
	err = service.NewCreateVideoService(ctx).CreateVideo(req)
	if err != nil {
		resp = &videoservice.CreateVideoResponse{
			StatusCode: 404,
			StatusMsg:  "fail...",
		}
		return resp, err
	}
	resp = &videoservice.CreateVideoResponse{
		StatusCode: 0,
		StatusMsg:  "success...",
	}
	return resp, nil
}

// GetVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoList(ctx context.Context, req *videoservice.GetVideoListRequest) (resp *videoservice.GetVideoListResponse, err error) {
	// TODO: Your code here...
	msg, err := service.NewGetVideoService(ctx).GetVideo(req)
	if err != nil {
		resp = &videoservice.GetVideoListResponse{
			StatusCode: 404,
			StatusMsg:  "fail...",
		}
		return resp, err
	}
	resp = &videoservice.GetVideoListResponse{
		StatusCode: 0,
		StatusMsg:  "success...",
		VideoList:  msg,
	}
	return resp, nil
}

// GetFeed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFeed(ctx context.Context, req *videoservice.GetFeedRequest) (resp *videoservice.GetFeedResponse, err error) {
	res, nextTime, err := service.NewGetFeedService(ctx).GetFeed(req)
	if err != nil {
		resp = &videoservice.GetFeedResponse{
			StatusCode: 500,
			StatusMsg:  "fail..",
			VideoList:  nil,
			NextTime:   nextTime,
		}
		return resp, err
	}
	resp = &videoservice.GetFeedResponse{
		StatusCode: 0,
		StatusMsg:  "success..",
		VideoList:  res,
		NextTime:   nextTime,
	}
	return resp, nil
}
