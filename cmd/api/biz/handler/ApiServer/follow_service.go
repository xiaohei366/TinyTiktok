package ApiServer

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/handler/pack"
	ApiServer "github.com/xiaohei366/TinyTiktok/cmd/api/biz/model/ApiServer"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/rpc"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

// RelationAction .
// @router /douyin/relation/action/ [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.DouyinRelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	v, _ := c.Get(shared.IdentityKey) // 取出token的id
	//调用PRC方法，在follow服务器上完成关注操作
	resp, err := rpc.FollowAction(context.Background(), &RelationServer.DouyinRelationActionRequest{
		UserId:     v.(*ApiServer.User).Id,
		ToUserId:   req.ToUserId,
		ActionType: req.ActionType,
	})
	if err != nil {
		pack.SendFollowActionResponse(c, err)
		return
	}
	//成功响应
	pack.SendFollowActionResponse(c, resp)
}

// RelationFollowList .
// @router /douyin/relation/follow/list/ [GET]
func RelationFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.DouyinRelationFollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	//调用PRC方法，获得关注列表
	resp, err := rpc.GetFollowList(context.Background(), &RelationServer.DouyinRelationFollowListRequest{
		UserId: req.UserId,
	})
	if err != nil {
		pack.SendFollowListResponse(c, err)
		return
	}
	pack.SendFollowListResponse(c, resp)
}

// RelationFollowerList .
// @router /douyin/relation/follower/list/ [GET]
func RelationFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.DouyinRelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	//调用PRC方法，获得粉丝列表
	resp, err := rpc.GetFollowerList(context.Background(), &RelationServer.DouyinRelationFollowerListRequest{
		UserId: req.UserId,
	})
	if err != nil {
		pack.SendFollowListResponse(c, err)
		return
	}
	//成功响应
	pack.SendFollowListResponse(c, resp)
}
