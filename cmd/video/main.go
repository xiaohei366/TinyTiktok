package main

import (
	mw "github.com/xiaohei366/TinyTiktok/pkg/middleware"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/xiaohei366/TinyTiktok/cmd/video/initialize"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer/videosrv"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

func main() {
	//init db and minio
	initialize.Init()
	//init etcd register
	r, err := etcd.NewEtcdRegistry([]string{shared.ETCDAddress})
	if err != nil {
		panic(err)
	}
	// get Tcp addr.
	addr, err := net.ResolveTCPAddr("tcp", shared.VideoServiceAddr)
	if err != nil {
		panic(err)
	}
	//init video server
	svr := videosrv.NewServer(new(VideoSrvImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: shared.VideoServiceName}),
	)
	//run video server
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
