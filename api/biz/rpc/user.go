package rpc

import (
	"context"
	"log"

	"github.com/gdan0324/ByteWeGo/api/kitex_gen/userservice"
	userSvc "github.com/gdan0324/ByteWeGo/api/kitex_gen/userservice/userservice"
	"github.com/gdan0324/ByteWeGo/api/pkg/consts"
	"github.com/gdan0324/ByteWeGo/api/pkg/mw"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userSvc.Client

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}

	c, err := userSvc.NewClient(
		consts.UserServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// create user info
func CreateUser(ctx context.Context, req *userservice.CreateUserRequest) (*userservice.CreateUserResponse, error) {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	log.Println(resp)

	return resp, nil
}

// check user info
func CheckUser(ctx context.Context, req *userservice.CheckUserRequest) (*userservice.CheckUserResponse, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return nil, err
	}
	log.Println(resp)

	return resp, nil
}

// get user info
func GetUser(ctx context.Context, req *userservice.GetUserRequest) (*userservice.GetUserResponse, error) {
	resp, err := userClient.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	log.Println(resp)

	return resp, nil
}
