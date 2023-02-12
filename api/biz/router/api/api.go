// Code generated by hertz generator. DO NOT EDIT.

package Api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	api "github.com/gdan0324/ByteWeGo/api/biz/handler/api"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		_douyin.GET("/user", append(_getuserMw(), api.GetUser)...)
		{
			_user := _douyin.Group("/user", _userMw()...)
			_user.POST("/login", append(_checkuserMw(), api.CheckUser)...)
			_user.POST("/register", append(_createuserMw(), api.CreateUser)...)
		}
	}
}
