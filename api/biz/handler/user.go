package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gdan0324/ByteWeGo/api/biz/rpc"
	"github.com/gdan0324/ByteWeGo/api/kitex_gen/userservice"
)

func CheckUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req userservice.CheckUserRequest
	// err = c.BindAndValidate(&req)
	// if err != nil {
	// 	c.String(consts.StatusBadRequest, err.Error())
	// 	return
	// }

	_, err = rpc.CheckUser(context.Background(), &userservice.CheckUserRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	c.JSON(consts.StatusOK, utils.H{
		"message": "user",
	})
}
