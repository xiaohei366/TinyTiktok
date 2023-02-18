package service

import (
	"context"

	"github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/rpc"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/service/pack"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/service/dal"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type MGetUserRelationFollowerService struct {
	ctx context.Context
}


func NewMGetUserRelationFollowerService(ctx context.Context) *MGetUserRelationFollowerService {
	return &MGetUserRelationFollowerService{ctx: ctx}
}

//获得粉丝列表
func (s *MGetUserRelationFollowerService) MGetUserRelationFollower(userID int64) ([]*RelationServer.User, error) {
	//先取出所有粉丝的ID
	followers, err := dal.MGetFollowerList(s.ctx, userID)
	if err != nil {
		return nil, errno.GetFollowerListErr
	}
	followerIDs := make([]int64, 0)
	for _, v := range followers {
		followerIDs = append(followerIDs, v.UserID)
	}
	//如果没有粉丝， 就直接返回即可
	if len(followerIDs) == 0 {
		return nil, nil
	}
	//随后通过RPC 由这些ID获得 用户信息
	var users []*UserServer.User
	users, err = rpc.MGetUserInfo(s.ctx, &UserServer.DouyinMUserRequest{
		UserId: followerIDs,
	})
	if(err != nil) {
		return nil, errno.UserRPCErr
	}

	//随后得到该用户的关注列表，用于补充用户信息
	followSet, err := dal.GetFollowSet(s.ctx, userID)
	if err != nil {
		return nil, err
	}

	//挨个转换结构体---并同时封装好“是否关注”信息
	r_users := make([]*RelationServer.User, 0)
	for _, u := range users {
		_, IsFollow := followSet[u.Id]
		r_users = append(r_users, pack.UserInfoConvert(u, IsFollow)) 
	}
	if len(users) != len(r_users) {
		return nil, errno.StructConvertFailedErr
	}

	return r_users, nil
}