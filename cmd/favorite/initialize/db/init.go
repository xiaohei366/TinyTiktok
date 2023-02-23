package db

import (
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/config"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

var DB *gorm.DB

// Favorite 表的结构。
type Favorite struct {
	gorm.Model
	UserId  int64 `json:"user_id" gorm:"not null"`   //点赞用户id
	VideoId int64 `json:"video_id" gorm:"not null""` //视频id
	Cancel  int32 `json:"cancel"`                    //是否点赞，0为点赞，1为取消赞
}

// TableName 修改表名映射
func (Favorite) TableName() string {
	return config.FavoriteTableName
}

// Init 初始化数据库
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
	err = DB.AutoMigrate(&Favorite{})
	if err != nil {
		klog.Fatalf("建表失败: %s", err.Error())
	}
}
