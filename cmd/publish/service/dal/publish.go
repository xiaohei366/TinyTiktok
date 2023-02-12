package dal

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/feed/service/dal"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

//func (v *dal.Video) TableName() string {
//	return config.PublishTableName
//}

// 这个实际不需要，就是publish
func PublishVideo(ctx context.Context, videoModel *dal.Video) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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

// 拿取User的视频的操作
func MGetUserVideos(ctx context.Context, userId int64) ([]*dal.Video, error) {
	klog.Info("dal/publish userId:", userId)
	videosList := make([]*dal.Video, 0)
	res := DB.WithContext(ctx).Model(&dal.Video{}).Where(&dal.Video{AuthorID: userId}).Find(&videosList)
	klog.Info("adl/publish res:", videosList)
	if res.RowsAffected == 0 {
		return nil, kerrors.NewBizStatusError(404, "User videos not exist")
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return videosList, nil
}
