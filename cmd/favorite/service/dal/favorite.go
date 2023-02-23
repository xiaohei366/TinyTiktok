package dal

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/config"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/initialize/db"
)

// 根据videoId获取点赞userId //这儿似乎没用上
func GetFavoriteUserIdList(ctx context.Context, videoId int64) ([]*db.Favorite, error) {
	favUserIdList := make([]*db.Favorite, 0) //存所有该视频点赞用户id；
	//查询likes表对应视频id点赞用户，返回查询结果
	err := db.DB.WithContext(ctx).Where(map[string]interface{}{"video_id": videoId, "favorite": config.IsFavorite}).
		Find(&favUserIdList).Error
	//查询过程出现错误，返回默认值0，并输出错误信息
	if err != nil {
		klog.Info("获取点赞列表失败")
		return nil, err
	} else {
		//没查询到或者查询到结果，返回数量以及无报错
		return favUserIdList, nil
	}
}

// 根据userId，videoId,actionType点赞或者取消赞
func UpdateFavorite(ctx context.Context, fav *db.Favorite) error {
	//更新当前用户观看视频的点赞状态“favorite”，返回错误结果
	err := db.DB.WithContext(ctx).Model(fav).Where(map[string]interface{}{"user_id": fav.UserId, "video_id": fav.VideoId}). //这儿有问题。
		Update("favorite", fav.Favorite).Error
	if err != nil {
		klog.Info("update data failed")
		return err
	}
	return nil
}

// 插入新的点赞数据
func InsertFavorite(ctx context.Context, fav *db.Favorite) error {
	//创建点赞数据，默认为点赞，cancel为0，返回错误结果
	err := db.DB.WithContext(ctx).Model(db.Favorite{}).Create(&fav).Error
	//如果有错误结果，返回插入失败
	if err != nil {
		klog.Info("insert favorite data failed")
		return errors.New("insert favorite data fail")
	}
	return nil
}

// 根据userId,videoId查询是否含有本条点赞信息
func GetFavoriteInfo(ctx context.Context, userId int64, videoId int64) (db.Favorite, error) {
	var favInfo db.Favorite
	//根据userid,videoId查询是否有该条信息，如果有，存储在favInfo,返回查询结果
	res := db.DB.WithContext(ctx).Where(map[string]interface{}{"user_id": userId, "video_id": videoId}).
		First(&favInfo)

	if res.RowsAffected == 0 {
		//没查询到数据
		klog.Info("user video favorite info not found")
		return favInfo, nil
	}
	if res.Error != nil {
		//查询出错
		klog.Info("get favInfo failed")
		return favInfo, res.Error
	}
	return favInfo, nil
}

// 根据userId查询所属点赞全部videoId
func GetFavoriteVideoIdList(ctx context.Context, userId int64) ([]db.Favorite, error) {
	favList := make([]db.Favorite, 0)
	res := db.DB.WithContext(ctx).Where(map[string]interface{}{"user_id": userId, "favorite": 1}).Find(&favList)

	if res.RowsAffected == 0 {
		//没查询到数据
		klog.Info("user favorite video list info not found")
		return favList, nil
	}
	if res.Error != nil {
		//查询出错
		klog.Info("get fav videos list failed")
		return favList, res.Error
	}
	return favList, nil
}

func GetVideoFavCountByVideoId(ctx context.Context, videoId int64) (int64, error) {
	favList := make([]db.Favorite, 0)
	err := db.DB.WithContext(ctx).Where(map[string]interface{}{"video_id": videoId, "favorite": 1}).
		Find(&favList).Error
	if err != nil {
		return int64(len(favList)), nil
	}
	return int64(len(favList)), nil
}

func QueryUserVideo(ctx context.Context, userId int64, videoId int64) (bool, error) {
	var fav db.Favorite
	//根据userid,videoId查询是否有该条信息，如果有，返回查询结果
	res := db.DB.WithContext(ctx).Where(map[string]interface{}{"user_id": userId, "video_id": videoId}).
		First(&fav) //这儿出错了，查不出来。
	if res.Error != nil {
		return false, nil
	}
	if res.RowsAffected == 0 {
		return false, nil
	} else if fav.Favorite == 1 {
		return true, nil
	}
	return false, nil
}
