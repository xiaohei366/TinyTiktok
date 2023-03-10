package dal

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/xiaohei366/TinyTiktok/cmd/video/config"
	"github.com/xiaohei366/TinyTiktok/cmd/video/initialize/db"
	"gorm.io/gorm"
	"time"
)

// MGetVideos multiple get list of videos info. Feed interface.
func MGetVideos(ctx context.Context, latestTime *int64) ([]*db.Video, error) {
	videoFeed := make([]*db.Video, 0)
	//处理时间
	if latestTime == nil || *latestTime == 0 {
		cur_time := int64(time.Now().UnixMilli())
		latestTime = &cur_time
	}
	res := db.DB.WithContext(ctx).Limit(config.Limit).Order("update_time desc").
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
func PublishVideo(ctx context.Context, videoModel *db.Video) error {
	err := db.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&videoModel).Error
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
func MGetUserVideos(ctx context.Context, userId int64) ([]*db.Video, error) {
	videosList := make([]*db.Video, 0)
	res := db.DB.Where(&db.Video{AuthorID: userId}).Find(&videosList) // 就是这句查询有问题。
	if res.RowsAffected == 0 {
		return nil, kerrors.NewBizStatusError(404, "User videos not exist")
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return videosList, nil
}

// GetVideosByVideosId
func GetVideoByVideosId(ctx context.Context, videoId int64) (*db.Video, error) {
	var video *db.Video
	res := db.DB.Where(&db.BaseModel{ID: videoId}).Find(&video)
	if res.RowsAffected == 0 {
		return nil, kerrors.NewBizStatusError(404, "videos not exist")
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return video, nil
}
