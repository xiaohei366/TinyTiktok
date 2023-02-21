package db

import (
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/user/config"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

var DB *gorm.DB

// 用户数据库表结构--自动创建
type User struct {
	Id            int64  `json:"id" gorm:"primarykey"`
	Name          string `json:"name" gorm:"unique;type:varchar(40);not null"`
	Password      string `json:"password" gorm:"type:varchar(256);not null"`
	FollowCount   int64  `json:"follow_count" gorm:"default:0"`
	FollowerCount int64  `json:"follower_count" gorm:"default:0"`
}

func (u *User) TableName() string {
	return config.UserTableName
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
	err = DB.AutoMigrate(&User{})
	if err != nil {
		klog.Fatalf("建表失败: %s", err.Error())
	}

}
