package initialize

import (
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/rpc"
)

func Init() {
	//数据库初始化
	db.Init()
	//rpc Init
	rpc.Init()
	//klog日志初始化
	InitLogger()
}
