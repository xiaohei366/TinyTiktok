package initialize

import (
	"github.com/xiaohei366/TinyTiktok/cmd/relation/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/initialize/mq"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/rpc"
)

func Init() {
	//数据库初始化
	db.Init()
	//RPC int
	rpc.Init()
	//阻塞队列初始化
	mq.InitMq()
	//klog日志初始化
	InitLogger()
}
