package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/gdan0324/ByteWeGo/api/pkg/consts"
	"github.com/gdan0324/ByteWeGo/api/pkg/errno"
	"github.com/gdan0324/ByteWeGo/api/pkg/mw"
	commentservice2 "github.com/gdan0324/ByteWeGo/comments/kitex_gen/commentservice"
	"github.com/gdan0324/ByteWeGo/comments/kitex_gen/commentservice/commentservice"
	"github.com/kitex-contrib/registry-etcd"
	"gorm.io/plugin/opentelemetry/provider"
)

var commentClient commentservice.Client

func initComment() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := commentservice.NewClient(
		consts.CommentServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.CommentServiceName}),
	)
	if err != nil {
		panic(err)
	}
	commentClient = c
}

func CommentAction(ctx context.Context, req *commentservice2.CommentActionRequest) (resp *commentservice2.CommentActionResponse, err error) {
	resp, err = commentClient.CommentAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}

func GetComments(ctx context.Context, req *commentservice2.GetCommentsRequest) (resp *commentservice2.GetCommentsResponse, err error) {
	resp, err = commentClient.GetComments(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}
