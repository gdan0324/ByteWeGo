// Code generated by Kitex v0.4.4. DO NOT EDIT.
package communityservice

import (
	server "github.com/cloudwego/kitex/server"
	communityservice "github.com/gdan0324/ByteWeGo/api/kitex_gen/communityservice"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler communityservice.CommunityService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}