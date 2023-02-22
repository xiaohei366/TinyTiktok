package service

import (
	"context"
	"strconv"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/initialize/redis"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/rpc"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/service/pack"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type MGetUserRelationFollowerService struct {
	ctx context.Context
}

func NewMGetUserRelationFollowerService(ctx context.Context) *MGetUserRelationFollowerService {
	return &MGetUserRelationFollowerService{ctx: ctx}
}

// 获得粉丝列表---协程池来补充完整用户信息
func (s *MGetUserRelationFollowerService) MGetUserRelationFollower(userID int64) ([]*RelationServer.User, error) {
	var wg sync.WaitGroup
	wg.Add(2)
	followerIDs := make([]int64, 0)
	//先查看Redis里面是否存在
	ids, _ := redis.Follower.MGet(redis.Ctx, strconv.Itoa(int(userID))).Result()
	if len(ids)-1 == 0 {
		//从数据库取出所有粉丝的ID
		followers, err := dal.MGetFollowerList(s.ctx, userID)
		if err != nil {
			return nil, errno.GetFollowerListErr
		}
		for _, v := range followers {
			followerIDs = append(followerIDs, v.UserID)
		}
		//更新redis
		redis.AddFollower(userID, followerIDs)
	} else {
		for _, v := range ids {
			followerIDs = append(followerIDs, v.(int64))
		}
	}
	//如果没有粉丝， 就直接返回即可
	if len(followerIDs) == 0 {
		return nil, nil
	}
	//随后通过RPC 由这些ID获得 用户信息,同时得到该用户的关注列表，用于补充用户信息
	//如果出现错误，不能直接返回失败，将默认值返回，保证稳定
	var users []*UserServer.User
	var followSet map[int64]struct{}
	go func() {
		var err error
		users, err = rpc.MGetUserInfo(s.ctx, &UserServer.DouyinMUserRequest{
			UserId: followerIDs,
		})
		if err != nil {
			klog.Errorf("调用用户RPC失败: %v", err)
		}
		wg.Done()
	}()
	//随后得到该用户的关注列表，用于补充用户信息
	go func() {
		var err error
		//先尝试使用Redis获取
		follows, _ := redis.Follow.MGet(redis.Ctx, strconv.Itoa(int(userID))).Result()
		if len(follows)-1 != 0 {
			for _, v := range follows {
				followSet[v.(int64)] = struct{}{}
			}
		} else {
			followSet, err = dal.GetFollowSet(s.ctx, userID)
			if err != nil {
				klog.Errorf("未获得关注列表：%v", err)
			}
		}
		wg.Done()
	}()
	wg.Wait()
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
