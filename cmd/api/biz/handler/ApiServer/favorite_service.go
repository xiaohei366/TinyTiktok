package ApiServer

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/handler/pack"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/model/ApiServer"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/rpc"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/FavoriteServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

// FavoriteAction .
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.DouyinFavoriteActionRequest
	err = c.Bind(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	// get user id
	user, _ := c.Get(shared.IdentityKey)
	// rpc user do the favorite action to the video
	resp, err := rpc.FavoriteAction(ctx, &FavoriteServer.DouyinFavoriteActionRequest{
		UserId:     user.(*ApiServer.User).Id,
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
	var req ApiServer.DouyinFavoriteListRequest
	err = c.Bind(&req)
	if err != nil {
		pack.SendFavoriteListResponse(c, errno.ConvertErr(err), nil)
		return
	}

	// rpc Get user's favorite video list
	resp, err := rpc.GetFavoriteList(ctx, &FavoriteServer.DouyinFavoriteListRequest{
		UserId: req.UserId,
	})
	fmt.Println("resp:")
	if err != nil {
		pack.SendFavoriteListResponse(c, errno.ConvertErr(err), nil)
		return
	}
	fmt.Println("api action list :", resp)
	pack.SendFavoriteListResponse(c, errno.Success, resp)
}
