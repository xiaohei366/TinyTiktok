package service

import (
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/service/dal"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// FavouriteVideoCount 根据videoId获取这个视频点赞数量
func (s *GetFavoriteService) GetVideoFavCount(videoId int64) (int64, error) {
	//获取
	f, err := dal.GetVideoFavCountByVideoId(s.ctx, videoId)
	if err != nil {
		return f, errno.UserNotExistErr
	}
	return f, nil
}

//
//// 根据userId获取这个用户点赞视频数量
//func (s *GetFavoriteService) FavouriteVideoCount(userId int64) (int64, error) {
//	f, err := dal.GetFavoriteUserIdList(s.ctx, userId)
//	if err != nil {
//		return int64(len(f)), errno.UserNotExistErr
//	}
//	return int64(len(f)), nil
//}
