package main

import (
	"context"
	"fmt"
	"github.com/gdan0324/ByteWeGo/api/pkg/jwt"
	commentservice "github.com/gdan0324/ByteWeGo/comments/kitex_gen/commentservice"
	"github.com/gdan0324/ByteWeGo/comments/service"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *commentservice.CommentActionRequest) (resp *commentservice.CommentActionResponse, err error) {
	resp = commentservice.NewCommentActionResponse()
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.SetStatusCode(20004)
		resp.SetStatusMsg("token authorize err")
		return resp, err
	}
	claimID := int64(claims["Id"].(float64))

	if req.UserId == 0 || claimID != 0 {
		req.UserId = claimID
	}

	fmt.Println("進入commentaction请求 ", req.VideoId, " ", req.UserId, " ", req.ActionType)
	if req.UserId == 0 || req.VideoId == 0 || (req.ActionType != 1 && req.ActionType != 2) {
		resp.SetStatusCode(20001)
		resp.SetStatusMsg("err param")
		return resp, nil
	}

	newcomment, err := service.NewCommentActionService(ctx).CommentAction(req, claimID)
	if err != nil {
		resp.SetStatusCode(20002)
		resp.SetStatusMsg("database err, action failed ")
		return resp, nil
	}

	// 成功后返回comment
	resp.SetStatusCode(0)
	resp.SetStatusMsg("action success")
	if req.ActionType == 1 {
		resp.SetComment(newcomment)
	}
	return resp, err
}

// GetComments implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GetComments(ctx context.Context, req *commentservice.GetCommentsRequest) (resp *commentservice.GetCommentsResponse, err error) {
	resp = commentservice.NewGetCommentsResponse()
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.SetStatusCode(20004)
		resp.SetStatusMsg("token authorize err")
		return resp, err
	}

	fmt.Println("进入评论列表请求 video_id = ", req.VideoId, "\n")

	if req.VideoId == 0 {
		resp.SetStatusCode(20001)
		resp.SetStatusMsg("err param!")
		return resp, nil // err para
	}

	comments, err := service.NewGetCommentsService(ctx).GetComments(req, int64(claims["Id"].(float64)))
	if err != nil {
		resp.SetStatusCode(20002)
		resp.SetStatusMsg("get comments database error!")
	}

	fmt.Printf("请求评论列表")
	//success
	resp.SetStatusCode(0)
	resp.SetStatusMsg("get comments success!")
	resp.SetCommentList(comments)
	return resp, nil
}

/*
	status_code:
	20000 : success
	20001 : err param
	20002 : database err
	20003 : error para bind // rpc comment handler
	20004 : token authorize err
*/
