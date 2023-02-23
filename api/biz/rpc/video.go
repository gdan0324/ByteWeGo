package rpc

import (
	"context"
	"github.com/gdan0324/ByteWeGo/api/pkg/consts"
	"github.com/gdan0324/ByteWeGo/api/pkg/mw"
	videoservice2 "github.com/gdan0324/ByteWeGo/video/kitex_gen/videoservice"
	"github.com/gdan0324/ByteWeGo/video/kitex_gen/videoservice/videoservice"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var videoClient videoservice.Client

func initVideo() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		consts.VideoServiceName,
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

// CreateVideo create video info
func CreateVideo(ctx context.Context, req *videoservice2.CreateVideoRequest) (*videoservice2.CreateVideoResponse, error) {
	resp, err := videoClient.CreateVideo(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetVideoList get video list
func GetVideoList(ctx context.Context, req *videoservice2.GetVideoListRequest) (*videoservice2.GetVideoListResponse, error) {
	resp, err := videoClient.GetVideoList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetFeed get feed
func GetFeed(ctx context.Context, req *videoservice2.GetFeedRequest) (*videoservice2.GetFeedResponse, error) {
	resp, err := videoClient.GetFeed(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
