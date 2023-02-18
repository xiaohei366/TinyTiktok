package db

import (
	"github.com/xiaohei366/TinyTiktok/cmd/video/config"
	"gorm.io/gorm"
	"time"
)

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
