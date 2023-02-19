package initialize

import (
	"github.com/xiaohei366/TinyTiktok/cmd/video/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/video/rpc"
	"github.com/xiaohei366/TinyTiktok/pkg/minio"
)

// Init Video database and minio
func Init() {
	//DB init
	db.Init_DB()
	minio.Init_minio()

	//RPC init
	rpc.Init()
	//klog日志初始化
	InitLogger()
}
