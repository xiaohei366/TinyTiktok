package service

import (
	"context"

	"github.com/xiaohei366/TinyTiktok/cmd/relation/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/rpc"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/service/dal"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type RelationActionService struct {
	ctx context.Context
}


func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

// 关注/取关操作
func (s *RelationActionService) FollowAction(isFollow bool, relation db.Follow) error {
	//先尝试修改user模块的数据库
	err := rpc.ChangeFollowCount(s.ctx, &UserServer.DouyinChangeUserFollowRequest{
		UserId: relation.UserID,
		ToUserId: relation.ToUserID,
		IsFollow: isFollow,
	})
	if(err != nil) {
		return errno.UserRPCErr
	}
	//成功后再修改本地数据库中的关系
	if isFollow {
		err = dal.AddFollow(s.ctx, relation.UserID, relation.ToUserID)
	} else {
		err = dal.DelFollow(s.ctx, relation.UserID, relation.ToUserID)
	}
	if err != nil {
		return errno.FollowActionErr
	}
	return nil
}
