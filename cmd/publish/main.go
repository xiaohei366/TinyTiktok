package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/xiaohei366/TinyTiktok/cmd/publish/initialize"
	"github.com/xiaohei366/TinyTiktok/cmd/publish/kitex_gen/PublishServer/publishsrv"
	mw "github.com/xiaohei366/TinyTiktok/pkg/middleware"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"

	"net"
)

func main() {
	//初始化
	initialize.Init()
	//注册中心
	r, err := etcd.NewEtcdRegistry([]string{shared.ETCDAddress})
	if err != nil {
		panic(err)
	}
	//获得开启监听的addr
	addr, err := net.ResolveTCPAddr("tcp", shared.PublishServiceAddr)
	if err != nil {
		panic(err)
	}
	//链路追踪相关设置
	// provider.NewOpenTelemetryProvider(
	// 	provider.WithServiceName(shared.UserServiceName),
	// 	provider.WithExportEndpoint(shared.ExportEndpoint),
	// 	provider.WithInsecure(),
	// )
	//启动服务器
	svr := publishsrv.NewServer(new(PublishSrvImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		//server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: shared.PublishServiceName}),
	)
	//启动
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
