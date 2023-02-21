package service

import (
	"context"
	"strconv"

	"github.com/xiaohei366/TinyTiktok/cmd/relation/initialize/redis"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/rpc"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/service/pack"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
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
	followIDs := make([]int64, 0)
	//先尝试使用Redis
	ids, _ := redis.Follow.MGet(redis.Ctx, strconv.Itoa(int(userID))).Result()
	if len(ids) == 0 {
		//不行再用数据库
		follows, err := dal.MGetFollowList(s.ctx, userID)
		if err != nil {
			return nil, errno.GetFollowListErr
		}
		for _, v := range follows {
			followIDs = append(followIDs, v.ToUserID)
		}
	} else {
		for _, v := range ids {
			followIDs = append(followIDs, v.(int64))
		}
	}
	//如果没有关注， 就直接返回即可
	if len(followIDs) == 0 {
		return nil, nil
	}
	//随后通过RPC 由这些ID获得 用户信息
	var users []*UserServer.User
	var err error
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
