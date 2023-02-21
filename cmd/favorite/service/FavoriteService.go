package service

import (
	"context"
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

// IsFavorite 根据userId,videoId查询点赞状态;
func (s *GetFavoriteService) IsFavorite(userId int64, videoId int64) (db.Favorite, error) {
	f, err := dal.GetLikeInfo(s.ctx, userId, videoId)
	if err != nil {
		return f, errno.FavoriteNotExistErr
	}
	return f, nil
}

// GetFavouriteCount 根据videoId获取对应点赞数量;
func (s *GetFavoriteService) GetFavouriteCount(videoId int64) (int64, error) {
	f, err := dal.GetLikeUserIdList(s.ctx, videoId)
	if err != nil {
		return int64(len(f)), errno.FavoriteNotExistErr
	}
	return int64(len(f)), nil
}

// FavouriteAction 根据userId，videoId,actionType对视频进行点赞或者取消赞操作;
func (s *GetFavoriteService) FavouriteAction(userId int64, videoId int64, actionType int32) error {
	err := dal.UpdateLike(s.ctx, userId, videoId, actionType)
	if err != nil {
		return errno.FavouriteActionErr
	}
	return nil
}

// TotalFavourite 根据userId获取这个用户总共被点赞数量
func (s *GetFavoriteService) TotalFavourite(userId int64) (int64, error) {
	//查询两者关系
	f, err := dal.GetLikeVideoIdList(s.ctx, userId)
	if err != nil {
		return int64(len(f)), errno.UserNotExistErr
	}
	return int64(len(f)), nil
}

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
func (s *GetFavoriteService) GetFavouriteList(userId int64) ([]*FavoriteServer.Video, error) {
	f, err := dal.GetLikeVideoIdList(s.ctx, userId)
	if err != nil {
		return nil, errno.FavoriteNotExistErr
	}
	return f, nil
}
