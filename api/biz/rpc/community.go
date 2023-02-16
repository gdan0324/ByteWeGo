package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/gdan0324/ByteWeGo/api/pkg/consts"
	"github.com/gdan0324/ByteWeGo/api/pkg/mw"
	communityservice2 "github.com/gdan0324/ByteWeGo/community/kitex_gen/communityservice"
	"github.com/gdan0324/ByteWeGo/community/kitex_gen/communityservice/communityservice"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var communityClient communityservice.Client

func initCommunity() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}

	c, err := communityservice.NewClient(
		consts.CommunityServiceName,
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.CommunityServiceName}),
	)
	if err != nil {
		panic(err)
	}
	communityClient = c
}

func Follow(ctx context.Context, req *communityservice2.FollowRequest) (resp *communityservice2.FollowResponse, err error) {
	resp, err = communityClient.Follow(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func FollowList(ctx context.Context, req *communityservice2.GetFollowRequest) (resp *communityservice2.GetFollowResponse, err error) {
	resp, err = communityClient.GetFollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func FollowerList(ctx context.Context, req *communityservice2.GetFollowerRequest) (resp *communityservice2.GetFollowerResponse, err error) {
	resp, err = communityClient.GetFollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
