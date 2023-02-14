package dal

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/xiaohei366/TinyTiktok/cmd/video/config"
	"github.com/xiaohei366/TinyTiktok/cmd/video/initialize/db/init"
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

// MGetVideos multiple get list of videos info. Feed interface.
func MGetVideos(ctx context.Context, latestTime *int64) ([]*Video, error) {
	videoFeed := make([]*Video, 0)
	//TODO 这个时间是怎么处理的。
	res := init.DB.WithContext(ctx).Limit(config.Limit).Order("update_time desc").
		Find(&videoFeed, "update_time < ?", time.UnixMilli(*latestTime))
	if res.RowsAffected == 0 {
		return nil, kerrors.NewBizStatusError(404, "Video feed not exist")
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return videoFeed, nil
}

// PublishVideo store the video info into database.
func PublishVideo(ctx context.Context, videoModel *Video) error {
	err := init.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(videoModel).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// MGetUserVideos multi get video list by user id.
func MGetUserVideos(ctx context.Context, userId int64) ([]*Video, error) {
	videosList := make([]*Video, 0)
	res := init.DB.WithContext(ctx).Where(&Video{AuthorID: userId}).Find(&videosList)
	if res.RowsAffected == 0 {
		return nil, kerrors.NewBizStatusError(404, "User videos not exist")
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return videosList, nil
}
