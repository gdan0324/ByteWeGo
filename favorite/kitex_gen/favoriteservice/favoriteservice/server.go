// Code generated by Kitex v0.4.4. DO NOT EDIT.
package favoriteservice

import (
	server "github.com/cloudwego/kitex/server"
	favoriteservice "github.com/gdan0324/ByteWeGo/favorite/kitex_gen/favoriteservice"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler favoriteservice.FavoriteService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
