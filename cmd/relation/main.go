package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"

	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	Init "github.com/xiaohei366/TinyTiktok/cmd/relation/initialize"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/kitex_gen/RelationServer/relationservice"
	mw "github.com/xiaohei366/TinyTiktok/pkg/middleware"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
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
	addr, err := net.ResolveTCPAddr("tcp", shared.RelationServiceAddr)
	if err != nil {
		panic(err)
	}
	//链路追踪相关设置
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(shared.RelationServiceName),
		provider.WithExportEndpoint(shared.ExportEndpoint),
		provider.WithInsecure(),
	)
	//启动服务器
	svr := relationservice.NewServer(new(RelationServerImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: shared.RelationServiceName}),
	)
	//启动
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
