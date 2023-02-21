package main

import (
	"context"
	FavoriteServer "github.com/xiaohei366/TinyTiktok/cmd/favorite/kitex_gen/FavoriteServer"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/service"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/service/pack"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *FavoriteServer.DouyinFavoriteActionRequest) (resp *FavoriteServer.DouyinFavoriteActionResponse, err error) {
	// TODO: Your code here...
	//检验格式
	if req.ActionType != 1 && req.ActionType != 2 {
		resp = pack.BuildfavoriteActionResp(errno.FavoriteActionTypeErr)
		return resp, nil
	}
	//进行点赞操作
	err = service.NewGetFavoriteService(ctx).FavouriteAction(req.Id, req.VideoId, req.ActionType)
	if err != nil {
		resp = pack.BuildfavoriteActionResp(err)
		return resp, nil
	}
	resp = pack.BuildfavoriteActionResp(errno.Success)
	return resp, nil
}

// GetFavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) GetFavoriteList(ctx context.Context, req *FavoriteServer.DouyinFavoriteListRequest) (resp *FavoriteServer.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	videos, err := service.NewGetFavoriteService(ctx).GetFavouriteList(req.UserId)
	if err != nil {
		resp = pack.BuildgetFavoriteListResp(err, videos)
		return resp, nil
	}
	resp = pack.BuildgetFavoriteListResp(errno.Success, videos)
	return resp, nil
}

// GetFavoriteUser implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) GetFavoriteUser(ctx context.Context, req *FavoriteServer.DouyinUserFavoriteRequest) (resp *FavoriteServer.DouyinUserFavoriteResponse, err error) {
	// TODO: Your code here...
	total, err := service.NewGetFavoriteService(ctx).TotalFavourite(req.UserId)
	favoriteCount, err := service.NewGetFavoriteService(ctx).FavouriteVideoCount(req.UserId)
	if err != nil {
		resp = pack.BuildfavoriteUserQueryResp(err, total, favoriteCount)
		return resp, nil
	}
	resp = pack.BuildfavoriteUserQueryResp(errno.Success, total, favoriteCount)
	return resp, nil
}

// GetFavoriteVideo implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) GetFavoriteVideo(ctx context.Context, req *FavoriteServer.DouyinVideoFavoriteRequest) (resp *FavoriteServer.DouyinVideoFavoriteResponse, err error) {
	// TODO: Your code here...
	count, err := service.NewGetFavoriteService(ctx).GetFavouriteCount(req.Id)
	if err != nil {
		resp = pack.BuildfavoriteVideoQueryResp(err, count, count != 0)
		return resp, nil
	}
	resp = pack.BuildfavoriteVideoQueryResp(errno.Success, count, count != 0)
	return resp, nil
}
