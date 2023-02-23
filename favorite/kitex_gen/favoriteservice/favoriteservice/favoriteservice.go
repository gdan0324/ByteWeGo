// Code generated by Kitex v0.4.4. DO NOT EDIT.

package favoriteservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	favoriteservice "github.com/gdan0324/ByteWeGo/favorite/kitex_gen/favoriteservice"
)

func serviceInfo() *kitex.ServiceInfo {
	return favoriteServiceServiceInfo
}

var favoriteServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FavoriteService"
	handlerType := (*favoriteservice.FavoriteService)(nil)
	methods := map[string]kitex.MethodInfo{
		"DoFavorite":      kitex.NewMethodInfo(doFavoriteHandler, newFavoriteServiceDoFavoriteArgs, newFavoriteServiceDoFavoriteResult, false),
		"GetFavoriteList": kitex.NewMethodInfo(getFavoriteListHandler, newFavoriteServiceGetFavoriteListArgs, newFavoriteServiceGetFavoriteListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "favoriteservice",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func doFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favoriteservice.FavoriteServiceDoFavoriteArgs)
	realResult := result.(*favoriteservice.FavoriteServiceDoFavoriteResult)
	success, err := handler.(favoriteservice.FavoriteService).DoFavorite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceDoFavoriteArgs() interface{} {
	return favoriteservice.NewFavoriteServiceDoFavoriteArgs()
}

func newFavoriteServiceDoFavoriteResult() interface{} {
	return favoriteservice.NewFavoriteServiceDoFavoriteResult()
}

func getFavoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favoriteservice.FavoriteServiceGetFavoriteListArgs)
	realResult := result.(*favoriteservice.FavoriteServiceGetFavoriteListResult)
	success, err := handler.(favoriteservice.FavoriteService).GetFavoriteList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceGetFavoriteListArgs() interface{} {
	return favoriteservice.NewFavoriteServiceGetFavoriteListArgs()
}

func newFavoriteServiceGetFavoriteListResult() interface{} {
	return favoriteservice.NewFavoriteServiceGetFavoriteListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) DoFavorite(ctx context.Context, req *favoriteservice.DoFavoriteRequest) (r *favoriteservice.DoFavoriteResponse, err error) {
	var _args favoriteservice.FavoriteServiceDoFavoriteArgs
	_args.Req = req
	var _result favoriteservice.FavoriteServiceDoFavoriteResult
	if err = p.c.Call(ctx, "DoFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFavoriteList(ctx context.Context, req *favoriteservice.GetFavoriteListRequest) (r *favoriteservice.GetFavoriteListResponse, err error) {
	var _args favoriteservice.FavoriteServiceGetFavoriteListArgs
	_args.Req = req
	var _result favoriteservice.FavoriteServiceGetFavoriteListResult
	if err = p.c.Call(ctx, "GetFavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
