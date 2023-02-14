package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/xiaohei366/TinyTiktok/cmd/video/initialize/db/init"
	"github.com/xiaohei366/TinyTiktok/pkg/minio"
)

// Init Video database and minio
func Init() {
	init.Init_DB()
	minio.Init_minio()

	//klog日志初始化
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
}
