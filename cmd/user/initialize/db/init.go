package db

import (
	"fmt"
	"time"

	"github.com/xiaohei366/TinyTiktok/cmd/user/config"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

var DB *gorm.DB

// 初始化数据库
func Init() {
	dsn := fmt.Sprintf(shared.MySqlDSN, config.SqlName, config.SqlPassword, config.SqlHost, config.SqlPort, shared.DBName)
	newLogger := logger.New(
		logrus.NewWriter(), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // 彩色打印
		},
	)

	// global mode
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		klog.Fatalf("init gorm failed: %s", err.Error())
	}
	if err := DB.Use(tracing.NewPlugin()); err != nil {
		klog.Fatalf("use tracing plugin failed: %s", err.Error())
	}
}