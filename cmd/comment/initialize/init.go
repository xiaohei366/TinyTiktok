package initialize

import (
	"github.com/xiaohei366/TinyTiktok/cmd/comment/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/comment/rpc"
	"github.com/xiaohei366/TinyTiktok/cmd/comment/initialize/redis"
)

func Init() {
	//数据库初始化
	db.Init()
	//RPC int
	rpc.Init()
	redis.InitRedis()
	//klog日志初始化
	InitLogger()
}
