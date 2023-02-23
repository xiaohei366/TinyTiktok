package service

import (
	"fmt"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/service/dal"
)

// FavouriteVideoCount 根据videoId获取这个视频点赞数量
func (s *GetFavoriteService) QueryUserVideo(userid, videoId int64) (bool, error) {
	//查询两者关系
	fmt.Println("userId,videoId:", userid, videoId)
	f, err := dal.QueryUserVideo(s.ctx, userid, videoId)
	if err != nil {
		return f, nil
	}
	return f, nil
}
