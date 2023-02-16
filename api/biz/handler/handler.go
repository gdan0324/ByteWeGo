package handler

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func SendResponse(c *app.RequestContext, response interface{}) {
	c.JSON(consts.StatusOK, response)
}

type CommentActionParam struct {
	Token       string `json:"token,omitempty"`
	VideoId     int64  `json:"video_id,omitempty"`
	ActionType  int32  `json:"action_type,omitempty"`
	CommentText string `json:"comment_text,omitempty"`
	CommentId   int64  `json:"comment_id,omitempty"`
}

type GetCommentsParam struct {
	Token   string `json:"token,omitempty"`
	VideoId int64  `json:"video_id,omitempty"`
}
