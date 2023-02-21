package ApiServer

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/handler/pack"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/kitex_gen/FavoriteServer"
	api "github.com/xiaohei366/TinyTiktok/cmd/api/biz/model/ApiServer"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/rpc"
)

// FavoriteAction .
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinFavoriteActionRequest
	err = c.Bind(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := rpc.FavoriteAction(ctx, &FavoriteServer.DouyinFavoriteActionRequest{
		VideoId:    req.VideoId,
		ActionType: req.ActionType,
	})
	if err != nil {
		pack.SendFavoriteActionResponse(c, err)
		return
	}

	pack.SendFavoriteActionResponse(c, resp)
}

// FavoriteList .
// @router /douyin/favorite/list/ [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinFavoriteListRequest
	err = c.Bind(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := rpc.GetFavoriteList(ctx, &FavoriteServer.DouyinFavoriteListRequest{
		UserId: req.UserId,
	})
	if err != nil {
		pack.SendFavoriteListResponse(c, err)
		return
	}

	pack.SendFavoriteListResponse(c, resp)
}
