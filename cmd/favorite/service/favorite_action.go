package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/service/dal"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/FavoriteServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type GetFavoriteService struct {
	ctx context.Context
}

func NewGetFavoriteService(ctx context.Context) *GetFavoriteService {
	return &GetFavoriteService{ctx: ctx}
}

// FavouriteAction 根据userId，videoId,actionType对视频进行点赞或者取消赞操作;
func (s *GetFavoriteService) FavouriteAction(req *FavoriteServer.DouyinFavoriteActionRequest) error {
	fav := &db.Favorite{
		UserId:   req.UserId,
		VideoId:  req.VideoId,
		Favorite: req.ActionType,
	}
	//先查询是否有数据
	favInfo, err := dal.GetFavoriteInfo(s.ctx, req.UserId, req.VideoId)
	if err != nil {
		klog.Info("cannot find user and video fav info")
		return errno.FavoriteActionErr
	} else {
		if favInfo == (db.Favorite{}) {
			//没有数据，插入新的数据
			if err := dal.InsertFavorite(s.ctx, fav); err != nil {
				klog.Info("Insert fav info error")
			}
		} else {
			//查到数据，更新
			if err := dal.UpdateFavorite(s.ctx, fav); err != nil {
				klog.Info("update like info error")
			}
		}
	}
	return nil
}
