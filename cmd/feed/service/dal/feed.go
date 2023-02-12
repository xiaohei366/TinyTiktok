package dal

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/xiaohei366/TinyTiktok/cmd/feed/config"
	"github.com/xiaohei366/TinyTiktok/cmd/feed/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/user/service/dal"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct { //自定义model，方便加上自己的字段。
	ID        int64     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

type Video struct {
	BaseModel
	Author   dal.User `gorm:"foreignkey:AuthorID"` //这个User是要去User那边拿的
	AuthorID int64    `gorm:"index:idx_authorid;not null"`

	PlayUrl  string `gorm:"type:varchar(200);not null"`
	CoverUrl string `gorm:"type:varchar(200);not null"`

	FavCount int64 `gorm:"type:int;default:0;not null"`
	ComCount int64 `gorm:"type:int;default:0;not null"`

	IsFavorite bool `gorm:"type:bool;default:false;not null"`

	Data  []byte `gorm:"column:video_data"`
	Title string `gorm:"type:varchar(50);not null"`
}

func (v *Video) TableName() string {
	return config.FeedTableName
}

// MGetUsers multiple get list of user info // 返回的是所有的视频的信息
func MGetVideos(ctx context.Context, latestTime *int64) ([]*Video, error) {
	videoFeed := make([]*Video, 0)
	//TODO 这个时间是怎么处理的。
	res := db.DB.Limit(config.Limit).Order("update_time desc").
		Find(&videoFeed, "update_time < ?", time.UnixMilli(*latestTime))
	if res.RowsAffected == 0 {
		return nil, kerrors.NewBizStatusError(404, "Video feed not exist")
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return videoFeed, nil
}
