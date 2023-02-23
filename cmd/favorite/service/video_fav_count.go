package service

import (
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/service/dal"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// FavouriteVideoCount 根据videoId获取这个视频点赞数量
func (s *GetFavoriteService) GetVideoFavCount(videoId int64) (int64, error) {
	//查询两者关系
	f, err := dal.GetVideoFavCountByVideoId(s.ctx, videoId)
	if err != nil {
		return f, errno.UserNotExistErr
	}
	return f, nil
}
