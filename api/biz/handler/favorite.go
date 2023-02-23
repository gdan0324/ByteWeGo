package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gdan0324/ByteWeGo/api/biz/rpc"
	"github.com/gdan0324/ByteWeGo/favorite/kitex_gen/favoriteservice"
	"strconv"
)

func Favorite(ctx context.Context, c *app.RequestContext) {
	var paramVar FavoriteActionParam
	token := c.PostForm("token")
	paramVar.Token = token
	video_id := c.PostForm("video_id")
	vi, err := strconv.Atoi(video_id)
	if err != nil {
		if err != nil {
			c.String(consts.StatusBadRequest, err.Error())
			return
		}
	}
	paramVar.VideoId = int64(vi)
	action_type := c.PostForm("action_type")
	at, err := strconv.Atoi(action_type)
	if err != nil {
		if err != nil {
			c.String(consts.StatusBadRequest, err.Error())
			return
		}
	}
	paramVar.ActionType = int32(at)
	resp, err := rpc.DoFavorite(ctx, &favoriteservice.DoFavoriteRequest{
		Token:      paramVar.Token,
		VideoId:    paramVar.VideoId,
		ActionType: paramVar.ActionType,
	})
	if err != nil {
		if err != nil {
			c.String(consts.StatusBadRequest, err.Error())
			return
		}
	}
	c.JSON(consts.StatusOK, resp)
}

func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var paramPar UserParam
	userId := c.Query("user_id")
	ui, err := strconv.Atoi(userId)
	if err != nil {
		if err != nil {
			c.String(consts.StatusBadRequest, err.Error())
			return
		}
	}
	paramPar.UserId = int64(ui)
	token := c.Query("token")
	paramPar.Token = token
	resp, err := rpc.GetFavoriteList(ctx, &favoriteservice.GetFavoriteListRequest{
		UserId: paramPar.UserId,
		Token:  paramPar.Token,
	})
	if err != nil {
		if err != nil {
			c.String(consts.StatusBadRequest, err.Error())
			return
		}
	}
	c.JSON(consts.StatusOK, resp)
}
