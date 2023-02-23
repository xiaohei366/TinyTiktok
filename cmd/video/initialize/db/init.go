package db

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

// BaseModel model ID and other info
type BaseModel struct { //自定义model，方便加上自己的字段。
	ID        int64     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

// Video model video info
type Video struct {
	BaseModel
	AuthorID int64 `gorm:"index:idx_authorid;not null"` // index, use user id to get video list

	PlayUrl  string `gorm:"type:varchar(200);not null"`
	CoverUrl string `gorm:"type:varchar(200);not null"`

	FavCount int64 `gorm:"type:int;default:0;not null"`
	ComCount int64 `gorm:"type:int;default:0;not null"`

	IsFavorite bool `gorm:"type:bool;default:false;not null"`

	Data  []byte `gorm:"column:video_data"`
	Title string `gorm:"type:varchar(50);not null"`
}

func (v *Video) TableName() string {
	return config.VideoTableName
}

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

	// global model
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //使生成表的时候是user,不是users。
		},
		Logger: newLogger,
	})
	if err != nil {
		klog.Fatalf("init gorm failed: %s", err.Error())
	}
	if err := DB.Use(tracing.NewPlugin()); err != nil {
		klog.Fatalf("use tracing plugin failed: %s", err.Error())
	}
	// AutoMigrate 会创建表、缺失的外键、约束、列和索引。 它不会删除未使用的列,只会增加没有的东西。
	err = DB.AutoMigrate(&Video{})
	if err != nil {
		klog.Fatalf("建表失败:%s", err.Error())
	}
}
