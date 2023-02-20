package db

import (
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/config"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

var DB *gorm.DB

// 用户数据库表结构--自动创建(关注、被关注)
type Follow struct {
	gorm.Model
	UserID   int64 `json:"user_id" gorm:"not null" `
	ToUserID int64 `json:"to_user_id" gorm:"not null" `
}

func (u *Follow) TableName() string {
	return config.FollowTableName
}

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
	// AutoMigrate 会创建表、缺失的外键、约束、列和索引。 它不会删除未使用的列,只会增加没有的东西。
	err = DB.AutoMigrate(&Follow{})
	if err != nil {
		klog.Fatalf("建表失败: %s", err.Error())
	}

}
