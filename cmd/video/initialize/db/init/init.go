package init

import (
	"fmt"
	"github.com/xiaohei366/TinyTiktok/cmd/video/config"
	"gorm.io/gorm/schema"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/xiaohei366/TinyTiktok/pkg/shared"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

var DB *gorm.DB

// Init initialize the Video database in Mysql and gorm.DB was declared to be used in service/dal to do sql.
func Init_DB() {

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
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //使生成表的时候使user,不是users。
		},
		Logger: newLogger,
	})
	if err != nil {
		klog.Fatalf("init gorm failed: %s", err.Error())
	}
	if err := DB.Use(tracing.NewPlugin()); err != nil {
		klog.Fatalf("use tracing plugin failed: %s", err.Error())
	}

}
