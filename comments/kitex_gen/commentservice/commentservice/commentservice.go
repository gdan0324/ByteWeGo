// Code generated by Kitex v0.4.4. DO NOT EDIT.

package commentservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	commentservice "github.com/gdan0324/ByteWeGo/comments/kitex_gen/commentservice"
)

func serviceInfo() *kitex.ServiceInfo {
	return commentServiceServiceInfo
}

var commentServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CommentService"
	handlerType := (*commentservice.CommentService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CommentAction": kitex.NewMethodInfo(commentActionHandler, newCommentServiceCommentActionArgs, newCommentServiceCommentActionResult, false),
		"GetComments":   kitex.NewMethodInfo(getCommentsHandler, newCommentServiceGetCommentsArgs, newCommentServiceGetCommentsResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "commentservice",
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

func commentActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*commentservice.CommentServiceCommentActionArgs)
	realResult := result.(*commentservice.CommentServiceCommentActionResult)
	success, err := handler.(commentservice.CommentService).CommentAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceCommentActionArgs() interface{} {
	return commentservice.NewCommentServiceCommentActionArgs()
}

func newCommentServiceCommentActionResult() interface{} {
	return commentservice.NewCommentServiceCommentActionResult()
}

func getCommentsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*commentservice.CommentServiceGetCommentsArgs)
	realResult := result.(*commentservice.CommentServiceGetCommentsResult)
	success, err := handler.(commentservice.CommentService).GetComments(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceGetCommentsArgs() interface{} {
	return commentservice.NewCommentServiceGetCommentsArgs()
}

func newCommentServiceGetCommentsResult() interface{} {
	return commentservice.NewCommentServiceGetCommentsResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CommentAction(ctx context.Context, req *commentservice.CommentActionRequest) (r *commentservice.CommentActionResponse, err error) {
	var _args commentservice.CommentServiceCommentActionArgs
	_args.Req = req
	var _result commentservice.CommentServiceCommentActionResult
	if err = p.c.Call(ctx, "CommentAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetComments(ctx context.Context, req *commentservice.GetCommentsRequest) (r *commentservice.GetCommentsResponse, err error) {
	var _args commentservice.CommentServiceGetCommentsArgs
	_args.Req = req
	var _result commentservice.CommentServiceGetCommentsResult
	if err = p.c.Call(ctx, "GetComments", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}