package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gdan0324/ByteWeGo/api/biz/rpc"
	"github.com/gdan0324/ByteWeGo/api/pkg/jwt"
	communityservice2 "github.com/gdan0324/ByteWeGo/community/kitex_gen/communityservice"
	"strconv"
)

func Follow(ctx context.Context, c *app.RequestContext) {
	var paramVar RelationActionParam
	token := c.PostForm("token")
	toUserId := c.PostForm("to_user_id")
	actionType := c.PostForm("action_type")
	tid, err := strconv.Atoi(toUserId)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	act, err := strconv.Atoi(actionType)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	paramVar.Token = token
	paramVar.ToUserId = int64(tid)
	paramVar.ActionType = int32(act)
	resp, err := rpc.Follow(ctx, &communityservice2.FollowRequest{
		Token:      paramVar.Token,
		ToUserId:   paramVar.ToUserId,
		ActionType: paramVar.ActionType,
	})
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	c.JSON(consts.StatusOK, resp)
}

func FollowList(ctx context.Context, c *app.RequestContext) {
	var paramVar UserParam
	uid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	paramVar.UserId = int64(uid)
	paramVar.Token = c.Query("token")
	resp, err := rpc.FollowList(ctx, &communityservice2.GetFollowRequest{
		UserId: paramVar.UserId,
		Token:  paramVar.Token,
	})
	c.JSON(consts.StatusOK, resp)
}

func FollowerList(ctx context.Context, c *app.RequestContext) {
	var paramVar UserParam
	uid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	paramVar.UserId = int64(uid)
	paramVar.Token = c.Query("token")
	resp, err := rpc.FollowerList(ctx, &communityservice2.GetFollowerRequest{
		UserId: paramVar.UserId,
		Token:  paramVar.Token,
	})
	c.JSON(consts.StatusOK, resp)
}

// CheckFriend .
// @router /douyin/relation/friend/list/ [GET]
func CheckFriend(ctx context.Context, c *app.RequestContext) {
	var paramVar UserParam
	uid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	paramVar.UserId = int64(uid)
	paramVar.Token = c.Query("token")
	resp, err := rpc.CheckFriend(ctx, &communityservice2.CheckFriendRequest{
		UserId: paramVar.UserId,
		Token:  paramVar.Token,
	})
	c.JSON(consts.StatusOK, resp)
}

// MessageAction .
// @router /douyin/message/action/ [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var paramVar RelationActionParam
	//todo token
	paramVar.Token, _ = jwt.GnerateToken("1")

	toUserId := c.PostForm("to_user_id")
	actionType := c.PostForm("action_type")
	content := c.PostForm("content")
	tid, err := strconv.Atoi(toUserId)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	act, err := strconv.Atoi(actionType)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	paramVar.ToUserId = int64(tid)
	paramVar.ActionType = int32(act)
	paramVar.Content = content
	resp, err := rpc.MessageAction(ctx, &communityservice2.RelationActionRequest{
		Token:      paramVar.Token,
		ToUserId:   paramVar.ToUserId,
		ActionType: paramVar.ActionType,
		Content:    paramVar.Content,
	})
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	c.JSON(consts.StatusOK, resp)
}
