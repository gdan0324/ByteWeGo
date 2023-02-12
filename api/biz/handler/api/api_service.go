// Code generated by hertz generator.

package api

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	api "github.com/gdan0324/ByteWeGo/api/biz/model/api"
)

// CheckUser .
// @router /douyin/user/login [POST]
func CheckUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CheckUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CheckUserResponse)

	c.JSON(consts.StatusOK, resp)
}

// CreateUser .
// @router /douyin/user/register [POST]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CreateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CreateUserResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetUser .
// @router /douyin/user [GET]
func GetUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.GetUserResponse)

	c.JSON(consts.StatusOK, resp)
}