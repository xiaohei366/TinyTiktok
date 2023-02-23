package service

import (
	"context"

	"github.com/xiaohei366/TinyTiktok/cmd/comment/rpc"
	"github.com/xiaohei366/TinyTiktok/cmd/comment/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/comment/service/pack"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/CommentServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type MGetCommentService struct {
	ctx context.Context
}

func NewMGetCommentService(ctx context.Context) *MGetCommentService {
	return &MGetCommentService{ctx: ctx}
}

// 获得评论列表列表
func (s *MGetCommentService) MGetComment(videoID int64) ([]*CommentServer.Comment, error) {

	//先取出所有评论的ID
	comments, err := dal.MGetCommentList(s.ctx, videoID)
	if err != nil {
		return nil, errno.GetCommentListErr
	}
	commentIDs := make([]int64, 0)
	for _, c := range comments {
		commentIDs = append(commentIDs, c.User_id)
	}
	//如果没有评论， 就直接返回即可
	if len(commentIDs) == 0 {
		return nil, nil
	}
	//随后通过RPC 由这些ID获得 用户信息
	users, err := rpc.MGetUserInfo(s.ctx, &UserServer.DouyinMUserRequest{
		UserId: commentIDs,
	})
	if err != nil {
		return nil, errno.UserRPCErr
	}
	//挨个转换结构体---并同时封装好users信息
	r_comments := make([]*CommentServer.Comment, 0)

	for i, c := range comments {
		//fmt.Print(c, "\n")
		r_comments = append(r_comments, pack.CommentInfoConvert(users[i], c))
	}
	if len(comments) != len(r_comments) {
		return nil, errno.StructConvertFailedErr
	}
	return r_comments, nil
}
