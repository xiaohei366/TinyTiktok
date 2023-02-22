package initialize

import (
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/initialize/db"
)

func Init() {
	//数据库初始化
	db.Init()

	//klog日志初始化
	InitLogger()
}
