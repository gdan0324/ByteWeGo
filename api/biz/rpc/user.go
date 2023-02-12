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

// CreateUser create user info
func CreateUser(ctx context.Context, req *userservice.CreateUserRequest) error {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	log.Println(resp)
	// if resp.BaseResp.Code != 0 {
	// 	return errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Message)
	// }
	return nil
}

// CheckUser check user info
func CheckUser(ctx context.Context, req *userservice.CheckUserRequest) (int64, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, err
	}
	// if resp.BaseResp.Code != 0 {
	// 	return 0, errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Message)
	// }
	return resp.UserId, nil
}
