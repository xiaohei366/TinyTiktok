package main

import (
	"github.com/xiaohei366/TinyTiktok/cmd/video/initialize/db/init"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/dal"
)

// auto migrate model (video) in mysql used by gorm.DB
func main() {
	init.Init_DB()
	//定义一个表结构，将表结构直接生成对应的表-migrations
	_ = init.DB.AutoMigrate(&dal.Video{})
}
