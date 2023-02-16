package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gdan0324/ByteWeGo/api/biz/rpc"
	"github.com/gdan0324/ByteWeGo/comments/kitex_gen/commentservice"
	"strconv"
)

func CommentAction(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")

	paramVideoId, err := strconv.Atoi(video_id)
	if err != nil {
		SendResponse(c, &commentservice.CommentActionResponse{
			StatusCode: 20003,
			StatusMsg:  "error para bind video_id",
		})
		return
	}
	paramActionType, err := strconv.Atoi(action_type)
	if err != nil {
		SendResponse(c, &commentservice.CommentActionResponse{
			StatusCode: 20003,
			StatusMsg:  "error para bind action_type",
		})
		return
	}
	rpcReq := commentservice.CommentActionRequest{
		VideoId:    int64(paramVideoId),
		Token:      token,
		ActionType: int32(paramActionType),
		UserId:     111,
	}
	if paramActionType == 1 {
		comment_text := c.Query("comment_text")
		rpcReq.SetCommentText(comment_text)
	} else {
		comment_id := c.Query("comment_id")
		paramCommentId, err := strconv.Atoi(comment_id)
		if err != nil {
			SendResponse(c, &commentservice.CommentActionResponse{
				StatusCode: 20003,
				StatusMsg:  "error para bind comment_id",
			})
			return
		}
		rpcReq.SetCommentId(int64(paramCommentId))
	}

	resp, err := rpc.CommentAction(ctx, &rpcReq)
	if err != nil {
		SendResponse(c, &commentservice.CommentActionResponse{
			StatusCode: 20003,
			StatusMsg:  "error para bind resp",
		})
		return
	}
	SendResponse(c, resp)
}

func GetComments(ctx context.Context, c *app.RequestContext) {
	video_id := c.Query("video_id")
	paramVideoId, err := strconv.Atoi(video_id)
	if err != nil {
		SendResponse(c, &commentservice.CommentActionResponse{
			StatusCode: 20003,
			StatusMsg:  "error para bind",
		})
		return
	}
	paramToken := c.Query("token")

	if len(paramToken) == 0 || paramVideoId < 0 {
		SendResponse(c, &commentservice.CommentActionResponse{
			StatusCode: 20003,
			StatusMsg:  "error para bind",
		})
	}

	resp, err := rpc.GetComments(ctx, &commentservice.GetCommentsRequest{
		VideoId: int64(paramVideoId),
		Token:   paramToken,
	})
	if err != nil {
		SendResponse(c, &commentservice.CommentActionResponse{
			StatusCode: 20003,
			StatusMsg:  "error para bind",
		})
		return
	}
	SendResponse(c, resp)
}
