package main

import (
	"context"


	"github.com/xiaohei366/TinyTiktok/cmd/comment/service"
	"github.com/xiaohei366/TinyTiktok/cmd/comment/service/pack"
	CommentServer "github.com/xiaohei366/TinyTiktok/kitex_gen/CommentServer"

)

// CommentServerImpl implements the last service interface defined in the IDL.
type CommentServerImpl struct{}

// CommentAction implements the CommentServerImpl interface.
func (s *CommentServerImpl) CommentAction(ctx context.Context, req *CommentServer.DouyinCommentActionRequest) (resp *CommentServer.DouyinCommentActionResponse, err error) {

	//进行创建或删除评论
	//var comment *CommentServer.Comment
	comment, err := service.NewCommentActionService(ctx).CommentAction(req) //在service中创建响应的.go文件（具体内容还没看，刚看到创建文件，文件里面大致是对数据库的操作）
	if err != nil {
		resp = pack.BuildcommentActionResp(err, nil)
		return resp, nil
	}
	resp = pack.BuildcommentActionResp(nil, comment)
	return resp, nil
}

// MGetCommentList implements the CommentServerImpl interface.
func (s *CommentServerImpl) CommentList(ctx context.Context, req *CommentServer.DouyinCommentListRequest) (resp *CommentServer.DouyinCommentListResponse, err error) {
	// 调用相应服务即可
	comments, err := service.NewMGetCommentService(ctx).MGetComment(req.VideoId)
	if err != nil {
		resp = pack.BuildgetCommentListResp(err, nil)
		return resp, nil
	}
	resp = pack.BuildgetCommentListResp(nil, comments)
	return resp, nil
}
