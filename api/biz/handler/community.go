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
	//todo token
	//token := c.PostForm("token")
	paramVar.Token, _ = jwt.GnerateToken("112")
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
