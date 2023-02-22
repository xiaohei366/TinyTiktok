package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	Init "github.com/xiaohei366/TinyTiktok/cmd/favorite/initialize"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/FavoriteServer/favoriteservice"
	mw "github.com/xiaohei366/TinyTiktok/pkg/middleware"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
	"net"
)

func main() {
	//初始化
	Init.Init()
	//注册中心
	r, err := etcd.NewEtcdRegistry([]string{shared.ETCDAddress})
	if err != nil {
		panic(err)
	}
	//获得开启监听的addr
	addr, err := net.ResolveTCPAddr("tcp", shared.FavoriteServiceAddr)
	if err != nil {
		panic(err)
	}
	//链路追踪相关设置
	// provider.NewOpenTelemetryProvider(
	// 	provider.WithServiceName(shared.FavoriteServiceName),
	// 	provider.WithExportEndpoint(shared.ExportEndpoint),
	// 	provider.WithInsecure(),
	// )
	//启动服务器
	svr := favoriteservice.NewServer(new(FavoriteServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		//server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: shared.FavoriteServiceName}),
	)
	//启动
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
