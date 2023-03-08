package initialize

import (
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/initialize/mq"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/rpc"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/initialize/redis"
)

func Init() {
	//数据库初始化
	db.Init()
	//rpc Init
	rpc.Init()
	//mq init
	mq.InitMq()
	redis.InitRedis()
	//klog日志初始化
	InitLogger()
}
