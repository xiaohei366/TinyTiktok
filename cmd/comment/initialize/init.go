package initialize

import (
	"github.com/xiaohei366/TinyTiktok/cmd/comment/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/comment/rpc"
)

func Init() {
	//数据库初始化
	db.Init()
	//RPC int
	rpc.Init()
	//klog日志初始化
	InitLogger()
}
