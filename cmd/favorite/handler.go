package main

import (
	"context"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/service"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/service/pack"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/FavoriteServer"
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
	err = service.NewGetFavoriteService(ctx).FavouriteAction(req)
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
	videos, err := service.NewGetFavoriteService(ctx).GetFavouriteList(req)
	if err != nil {
		resp = pack.BuildgetFavoriteListResp(err, videos)
		return resp, nil
	}
	if len(videos) == 0 {
		return nil, errno.Success
	}
	resp = pack.BuildgetFavoriteListResp(nil, videos)
	return resp, nil
}

// GetFavoriteUser implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) GetFavoriteUser(ctx context.Context, req *FavoriteServer.DouyinUserBeFavoriteRequest) (resp *FavoriteServer.DouyinUserBeFavoriteResponse, err error) {
	// 项目未用到..
	return resp, nil
}

// GetFavoriteVideo implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) GetFavoriteVideo(ctx context.Context, req *FavoriteServer.DouyinVideoBeFavoriteRequest) (resp *FavoriteServer.DouyinVideoBeFavoriteResponse, err error) {
	// TODO: Your code here...
	count, err := service.NewGetFavoriteService(ctx).GetVideoFavCount(req.VideoId)
	if err != nil {
		resp = pack.BuildfavoriteVideoQueryResp(err, count)
		return resp, nil
	}
	resp = pack.BuildfavoriteVideoQueryResp(errno.Success, count)
	return resp, nil
}

// QueryUserLikeVideo implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) QueryUserLikeVideo(ctx context.Context, req *FavoriteServer.DouyinQueryFavoriteRequest) (resp *FavoriteServer.DouyinQueryFavoriteResponse, err error) {
	// TODO: Your code here...
	fav, err := service.NewGetFavoriteService(ctx).QueryUserVideo(req.UserId, req.VideoId)
	if err != nil {
		return pack.BuildQueryUserFavoriteVideoResp(err, false), nil
	}
	return pack.BuildQueryUserFavoriteVideoResp(nil, fav), nil
}
