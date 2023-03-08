package service

import (
	"context"
	"time"

	"github.com/xiaohei366/TinyTiktok/cmd/comment/config"
	"github.com/xiaohei366/TinyTiktok/cmd/comment/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/comment/initialize/redis"

	"github.com/xiaohei366/TinyTiktok/cmd/comment/rpc"
	"github.com/xiaohei366/TinyTiktok/cmd/comment/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/comment/service/pack"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/CommentServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type CommentActionService struct {
	ctx context.Context
}

func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

// 创建/删除操作
func (s *CommentActionService) CommentAction(req *CommentServer.DouyinCommentActionRequest) (*CommentServer.Comment, error) {

	var err error
	if req.ActionType == 1 {
		commentModel := &db.Comment{
			User_id:      req.UserId,
			Video_id:     req.VideoId,
			Comment_text: req.CommentText,
		}
		//step1
		redis.DelComment(req.VideoId, req.CommentId)
		//step2
		comment, err := dal.AddComment(s.ctx, commentModel)
		//step3
		time.Sleep(config.SleepTime)
		//step4
		redis.DelComment(req.VideoId, req.CommentId)
		if err != nil {
			return nil, errno.CommentActionErr
		}

		//添加评论成功后拼接comment返回
		user, err := rpc.GetUserInfo(s.ctx, &UserServer.DouyinUserRequest{
			UserId: req.UserId,
		})
		if err != nil {
			return nil, errno.UserRPCErr
		}

		comment1 := pack.CommentInfoConvert(user, comment)
		return comment1, nil

	} else {
		//step1
		redis.DelComment(req.VideoId, req.CommentId)
		//step2
		err = dal.DelComment(s.ctx, req)
		//step3
		time.Sleep(config.SleepTime)
		//step4
		redis.DelComment(req.VideoId, req.CommentId)
		if err != nil {
			return nil, errno.CommentActionErr
		}
		//删除评论返回空
		return nil, nil
	}

}
