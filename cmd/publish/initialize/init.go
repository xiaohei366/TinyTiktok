package initialize

import (
	"github.com/xiaohei366/TinyTiktok/cmd/publish/initialize/db"
	"github.com/xiaohei366/TinyTiktok/pkg/minio"
)

func Init() {
	db.Init()
	minio.Init_minio()
}
