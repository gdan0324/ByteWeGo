package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/gdan0324/ByteWeGo/api/pkg/consts"
	"github.com/gdan0324/ByteWeGo/api/pkg/mw"
	"github.com/gdan0324/ByteWeGo/comments/dal"
	"github.com/gdan0324/ByteWeGo/comments/kitex_gen/commentservice/commentservice"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	initService()
	//testclient()
}

func initService() {
	dal.Init()
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddress})
	if err != nil {
		klog.Fatal(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", consts.CommentServiceAddr)
	if err != nil {
		klog.Fatal(err)
	}

	svr := commentservice.NewServer(new(CommentServiceImpl),
		server.WithServiceAddr(addr),                                       // address
		server.WithMiddleware(mw.CommonMiddleware),                         // middleware
		server.WithMiddleware(mw.ServerMiddleware),                         // middleware
		server.WithRegistry(r),                                             // registry
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.CommentServiceName}))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
