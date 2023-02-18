package service

import (
	"context"

	"github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/rpc"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/service/pack"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type MGetUserRelationFollowService struct {
	ctx context.Context
}

func NewMGetUserRelationFollowService(ctx context.Context) *MGetUserRelationFollowService {
	return &MGetUserRelationFollowService{ctx: ctx}
}

// 获得关注列表
func (s *MGetUserRelationFollowService) MGetUserRelationFollow(userID int64) ([]*RelationServer.User, error) {
	//先取出所有关注的ID
	follows, err := dal.MGetFollowList(s.ctx, userID)
	if err != nil {
		return nil, errno.GetFollowListErr
	}
	followIDs := make([]int64, 0)
	for _, v := range follows {
		followIDs = append(followIDs, v.ToUserID)
	}
	//如果没有关注， 就直接返回即可
	if len(followIDs) == 0 {
		return nil, nil
	}
	//随后通过RPC 由这些ID获得 用户信息
	var users []*UserServer.User
	users, err = rpc.MGetUserInfo(s.ctx, &UserServer.DouyinMUserRequest{
		UserId: followIDs,
	})
	if err != nil {
		return nil, errno.UserRPCErr
	}

	//挨个转换结构体---并同时封装好“是否关注”信息
	r_users := make([]*RelationServer.User, 0)
	for _, u := range users {
		r_users = append(r_users, pack.UserInfoConvert(u, true))
	}
	if len(users) != len(r_users) {
		return nil, errno.StructConvertFailedErr
	}

	return r_users, nil
}