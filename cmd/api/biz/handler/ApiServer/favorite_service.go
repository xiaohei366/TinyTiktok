package ApiServer

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	ApiServer "github.com/xiaohei366/TinyTiktok/cmd/api/biz/model/ApiServer"
)

// FavoriteAction .
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(ApiServer.DouyinFavoriteActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// FavoriteList .
// @router /douyin/favorite/list/ [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.DouyinFavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(ApiServer.DouyinFavoriteActionResponse)

	c.JSON(consts.StatusOK, resp)
}
