package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	Init "github.com/xiaohei366/TinyTiktok/cmd/feed/initialize/db"
	FeedServer "github.com/xiaohei366/TinyTiktok/cmd/feed/kitex_gen/FeedServer/feedsrv"
	mw "github.com/xiaohei366/TinyTiktok/pkg/middleware"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
	"net"
)

// APK安装，WSA跟着视频下载。
func main() {
	//初始化
	Init.Init()
	//注册中心
	r, err := etcd.NewEtcdRegistry([]string{shared.ETCDAddress})
	if err != nil {
		panic(err)
	}
	//获得开启监听的addr
	addr, err := net.ResolveTCPAddr("tcp", shared.FeedServiceAddr)
	if err != nil {
		panic(err)
	}

	//启动服务器
	//启动
	svr := FeedServer.NewServer(new(FeedSrvImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		//server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: shared.FeedServiceName}))

	//启动服务
	err = svr.Run() //这个会阻塞，所以放在协程里
	if err != nil {
		klog.Fatal(err)
	}
}
