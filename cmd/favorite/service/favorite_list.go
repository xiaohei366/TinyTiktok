package service

import (
	"strconv"

	"github.com/xiaohei366/TinyTiktok/cmd/favorite/initialize/redis"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/rpc"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/service/pack"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/FavoriteServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// 获取当前用户的所有点赞视频
func (s *GetFavoriteService) GetFavouriteList(req *FavoriteServer.DouyinFavoriteListRequest) ([]*FavoriteServer.Video, error) {
	//先尝试使用redis
	ids, err := redis.UserLikeList.MGet(redis.Ctx, strconv.Itoa(int(req.UserId))).Result()
	videoIds := []int64{}
	if err != nil || len(ids)-1 == 0 {
		//不行再用数据库
		favs, err := dal.GetFavoriteVideoIdList(s.ctx, req.UserId)
		if err != nil {
			return nil, errno.FavoriteVideoListNotExistErr
		}
		for _, fav := range favs {
			videoIds = append(videoIds, fav.VideoId)
			//更新Redis
			redis.AddUserLikeList(req.UserId, fav.VideoId)
		}
	} else {
		for _, v := range ids {
			videoIds = append(videoIds, v.(int64))
		}
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
