package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/rpc"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/service/pack"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/FavoriteServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type GetFavoriteService struct {
	ctx context.Context
}

func NewGetFavoriteService(ctx context.Context) *GetFavoriteService {
	return &GetFavoriteService{ctx: ctx}
}

// QueryFavorite 根据userId,videoId查询点赞状态;??没用啊
//func (s *GetFavoriteService) IsFavorite(userId int64, videoId int64) (db.Favorite, error) {
//	f, err := dal.GetLikeInfo(s.ctx, userId, videoId)
//	if err != nil {
//		return f, errno.FavoriteNotExistErr
//	}
//	return f, nil
//}

//// GetFavouriteCount 根据videoId获取对应点赞数量;
//func (s *GetFavoriteService) GetFavouriteCount(videoId int64) (int64, error) {
//	f, err := dal.GetLikeUserIdList(s.ctx, videoId) //??
//	if err != nil {
//		return int64(len(f)), errno.FavoriteNotExistErr
//	}
//	return int64(len(f)), nil
//}

// FavouriteAction 根据userId，videoId,actionType对视频进行点赞或者取消赞操作;
func (s *GetFavoriteService) FavouriteAction(req *FavoriteServer.DouyinFavoriteActionRequest) error {
	fmt.Println("Favourite Action:", req.UserId, req.VideoId, req.ActionType)
	fav := &db.Favorite{
		UserId:  req.UserId,
		VideoId: req.VideoId,
		Cancel:  req.ActionType,
	}
	//先查询是否有数据
	favInfo, err := dal.GetLikeInfo(s.ctx, req.UserId, req.VideoId)
	if err != nil {
		klog.Info("cannot find user and video fav info")
	} else {
		if favInfo == (db.Favorite{}) {
			//没有数据，新建
			if err := dal.InsertLike(s.ctx, fav); err != nil {
				klog.Info("Insert like info error")
			}
		} else {
			//查到数据，更新
			if err := dal.UpdateLike(s.ctx, fav); err != nil {
				klog.Info("update like info error")
			}
		}
	}
	return nil
}

//// TotalFavourite 根据userId获取这个用户总共被点赞数量
//func (s *GetFavoriteService) TotalFavourite(userId int64) (int64, error) {
//	//查询两者关系
//	f, err := dal.GetLikeVideoIdList(s.ctx, userId)
//	if err != nil {
//		return int64(len(f)), errno.UserNotExistErr
//	}
//	return int64(len(f)), nil
//}

// FavouriteVideoCount 根据userId获取这个用户点赞视频数量
func (s *GetFavoriteService) FavouriteVideoCount(userId int64) (int64, error) {
	//查询两者关系
	f, err := dal.GetLikeVideoIdList(s.ctx, userId)
	if err != nil {
		return int64(len(f)), errno.UserNotExistErr
	}
	return int64(len(f)), nil
}

// GetFavouriteList 获取当前用户的所有点赞视频
func (s *GetFavoriteService) GetFavouriteList(req *FavoriteServer.DouyinFavoriteListRequest) ([]*FavoriteServer.Video, error) {
	favs, err := dal.GetLikeVideoIdList(s.ctx, req.UserId)
	if err != nil {
		return nil, errno.FavoriteVideoListNotExistErr
	}
	videoIds := []int64{}
	for _, fav := range favs {
		videoIds = append(videoIds, fav.VideoId)
	}
	fmt.Println("videoIds:", videoIds)
	videoList, err := rpc.GetVideoListByVideoId(s.ctx, &VideoServer.DouyinVideoListByVideoId{
		VideoId: videoIds,
		UserId:  req.UserId, //获取与这些视频相关的信息
	})
	fmt.Println("videoList:", videoList)
	if err != nil {
		return nil, errno.FavoriteVideoListErr //这个要改。
	}
	resp := pack.ConvertVideos(videoList)
	//fmt.Println("server convert resp:", resp)//这个可以，会打印。
	return resp, nil
}
