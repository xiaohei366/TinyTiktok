package dal

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/config"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/kitex_gen/FavoriteServer"
)

// GetLikeUserIdList 根据videoId获取点赞userId
func GetLikeUserIdList(ctx context.Context, videoId int64) ([]*db.Favorite, error) {
	likeUserIdList := make([]*db.Favorite, 0) //存所有该视频点赞用户id；
	//查询likes表对应视频id点赞用户，返回查询结果
	err := db.DB.WithContext(ctx).Where(map[string]interface{}{"video_id": videoId, "cancel": config.IsFavorite}).
		Find(&likeUserIdList).Error
	//查询过程出现错误，返回默认值0，并输出错误信息
	if err != nil {
		klog.Info("获取点赞列表失败")
		return nil, err
	} else {
		//没查询到或者查询到结果，返回数量以及无报错
		return likeUserIdList, nil
	}
}

// UpdateLike 根据userId，videoId,actionType点赞或者取消赞
func UpdateLike(ctx context.Context, userId int64, videoId int64, actionType int32) error {
	//更新当前用户观看视频的点赞状态“cancel”，返回错误结果
	err := db.DB.WithContext(ctx).Where(map[string]interface{}{"user_id": userId, "video_id": videoId}).
		Update("cancel", actionType).Error
	//如果出现错误，返回更新数据库失败
	if err != nil {
		klog.Info("update data failed")
		return err
	}
	//更新操作成功
	return nil
}

// InsertLike 插入点赞数据
func InsertLike(ctx context.Context, favoriteData db.Favorite) error {
	//创建点赞数据，默认为点赞，cancel为0，返回错误结果
	err := db.DB.WithContext(ctx).Create(&favoriteData).Error
	//如果有错误结果，返回插入失败
	if err != nil {
		klog.Info("insert favoriteData failed")
		return err
	}
	return nil
}

// GetLikeInfo 根据userId,videoId查询点赞信息
func GetLikeInfo(ctx context.Context, userId int64, videoId int64) (db.Favorite, error) {
	//创建一条空like结构体，用来存储查询到的信息
	var likeInfo db.Favorite
	//根据userid,videoId查询是否有该条信息，如果有，存储在likeInfo,返回查询结果
	err := db.DB.WithContext(ctx).Where(map[string]interface{}{"user_id": userId, "video_id": videoId}).
		First(&likeInfo).Error
	if err != nil {
		//查询数据为0，打印"can't find data"，返回空结构体，这时候就应该要考虑是否插入这条数据了
		if "record not found" == err.Error() {
			klog.Info("can't find data")
			return db.Favorite{}, nil
		} else {
			//如果查询数据库失败，返回获取likeInfo信息失败
			klog.Info("get likeInfo failed")
			return likeInfo, err
		}
	}
	return likeInfo, nil
}

// GetLikeVideoIdList 根据userId查询所属点赞全部videoId
func GetLikeVideoIdList(ctx context.Context, userId int64) ([]*FavoriteServer.Video, error) {
	likeVideoIdList := make([]*FavoriteServer.Video, 0)
	err := db.DB.WithContext(ctx).Where(map[string]interface{}{"user_id": userId, "cancel": config.IsFavorite}).
		Pluck("video_id", &likeVideoIdList).Error
	if err != nil {
		//查询数据为0，返回空likeVideoIdList切片，以及返回无错误
		if "record not found" == err.Error() {
			klog.Info("there are no likeVideoId")
			return likeVideoIdList, nil
		} else {
			//如果查询数据库失败，返回获取likeVideoIdList失败
			klog.Info(err.Error())
			return likeVideoIdList, errors.New("get likeVideoIdList failed")
		}
	}
	return likeVideoIdList, nil
}
