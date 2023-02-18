package ApiServer

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	ApiServer "github.com/xiaohei366/TinyTiktok/cmd/api/biz/model/ApiServer"
)


// CommentAction .
// @router /douyin/comment/action/ [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.DouyinCommentActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(ApiServer.DouyinCommentActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// CommentList .
// @router /douyin/comment/list/ [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.DouyinCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(ApiServer.DouyinCommentListResponse)

	c.JSON(consts.StatusOK, resp)
}