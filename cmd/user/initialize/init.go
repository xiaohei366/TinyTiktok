package initialize

import (
	"github.com/xiaohei366/TinyTiktok/cmd/user/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/user/initialize/mq"
	"github.com/xiaohei366/TinyTiktok/cmd/user/initialize/redis"
	"github.com/xiaohei366/TinyTiktok/cmd/user/initialize/rpc"
)

func Init() {
	//数据库初始化
	db.Init()
	//mq和redis的初始化
	redis.InitRedis()
	mq.InitMq()
	rpc.Init()
	//klog日志初始化
	InitLogger()
}
