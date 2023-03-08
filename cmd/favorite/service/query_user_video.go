package service

import (
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/service/dal"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// 查询user是否点赞该视频
func (s *GetFavoriteService) QueryUserVideo(userid, videoId int64) (bool, error) {

	f, err := dal.QueryUserVideo(s.ctx, userid, videoId)
	if err != nil {
		return f, errno.QueryUserLikeVideoErr
	}
	return f, nil
}
