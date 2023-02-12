package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/xiaohei366/TinyTiktok/cmd/user/initialize/db"
)

func Init() {
	//数据库初始化
	db.Init()

	//klog日志初始化
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
}
