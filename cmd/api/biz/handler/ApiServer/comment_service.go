package ApiServer

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/handler/pack"
	ApiServer "github.com/xiaohei366/TinyTiktok/cmd/api/biz/model/ApiServer"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/rpc"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/CommentServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

// CommentAction .
// @router /douyin/comment/action/ [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.DouyinCommentActionRequest
	err = c.Bind(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	v, _ := c.Get(shared.IdentityKey) // 取出token的id
	//调用PRC方法，在follow服务器上完成关注操作
	resp, err := rpc.CommentAction(context.Background(), &CommentServer.DouyinCommentActionRequest{
		UserId:      v.(*ApiServer.User).Id,
		VideoId:     req.VideoId,
		ActionType:  req.ActionType,
		CommentText: req.CommentText,
		CommentId:   req.CommentId,
	})

	if err != nil {
		pack.SendCommentActionResponse(c, err)
		return
	}
	//成功响应
	pack.SendCommentActionResponse(c, resp)
}

// CommentList .
// @router /douyin/comment/list/ [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var useID int64
	var req ApiServer.DouyinCommentListRequest
	err = c.Bind(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	//先从token尝试取出UserId
	user, _ := c.Get(shared.IdentityKey)
	if user == nil {
		useID = 0
	} else {
		useID = user.(*ApiServer.User).Id
	}
	//调用PRC方法，获得关注列表
	resp, err := rpc.CommentList(context.Background(), &CommentServer.DouyinCommentListRequest{
		Token:   req.Token,
		VideoId: req.VideoId,
	})
	//再调用RPC方法，查询他们的关系
	for _, v := range resp.CommentList {
		isFollow, _ := rpc.QueryRelation(context.Background(), &RelationServer.DouyinQueryRelationRequest{
			UserId:   useID,
			ToUserId: v.User.Id,
		})
		v.User.IsFollow = isFollow
	}
	if err != nil {
		pack.SendCommentListResponse(c, err)
		return
	}
	pack.SendCommentListResponse(c, resp)
}
