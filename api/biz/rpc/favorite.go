package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/gdan0324/ByteWeGo/api/pkg/consts"
	"github.com/gdan0324/ByteWeGo/api/pkg/mw"
	"github.com/gdan0324/ByteWeGo/favorite/kitex_gen/favoriteservice"
	favoriteSrv "github.com/gdan0324/ByteWeGo/favorite/kitex_gen/favoriteservice/favoriteservice"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var favoriteClient favoriteSrv.Client

func initFavorite() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}

	c, err := favoriteSrv.NewClient(
		consts.FavoriteServiceName,
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}

func DoFavorite(ctx context.Context, req *favoriteservice.DoFavoriteRequest) (*favoriteservice.DoFavoriteResponse, error) {
	resp, err := favoriteClient.DoFavorite(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetFavoriteList(ctx context.Context, req *favoriteservice.GetFavoriteListRequest) (*favoriteservice.GetFavoriteListResponse, error) {
	resp, err := favoriteClient.GetFavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
