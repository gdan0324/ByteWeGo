package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/gdan0324/ByteWeGo/api/pkg/consts"
	"github.com/gdan0324/ByteWeGo/api/pkg/mw"
	"github.com/gdan0324/ByteWeGo/video/dal"
	videoservice "github.com/gdan0324/ByteWeGo/video/kitex_gen/videoservice/videoservice"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func Init() {
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(consts.TCP, consts.VideoServiceAddr)
	if err != nil {
		panic(err)
	}
	Init()

	svr := videoservice.NewServer(new(VideoServiceImpl),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.VideoServiceName}),
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
