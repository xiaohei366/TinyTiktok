package service

import (

	"github.com/xiaohei366/TinyTiktok/cmd/favorite/rpc"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/service/pack"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/FavoriteServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// 获取当前用户的所有点赞视频
func (s *GetFavoriteService) GetFavouriteList(req *FavoriteServer.DouyinFavoriteListRequest) ([]*FavoriteServer.Video, error) {
	favs, err := dal.GetFavoriteVideoIdList(s.ctx, req.UserId)
	if err != nil {
		return nil, errno.FavoriteVideoListNotExistErr
	}
	videoIds := []int64{}
	for _, fav := range favs {
		videoIds = append(videoIds, fav.VideoId)
	}
	//调用rpc获取视频列表
	videoList, err := rpc.GetVideoListByVideoId(s.ctx, &VideoServer.DouyinVideoListByVideoId{
		VideoId: videoIds,
		UserId:  req.UserId,
	})
	if err != nil {
		return nil, errno.FavoriteVideoListErr
	}

	resp := pack.ConvertVideos(videoList)
	return resp, nil
}
