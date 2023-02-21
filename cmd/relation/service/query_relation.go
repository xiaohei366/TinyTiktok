package service

import (
	"context"
	"strconv"

	"github.com/xiaohei366/TinyTiktok/cmd/relation/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/initialize/redis"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/service/dal"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

type QueryRelationService struct {
	ctx context.Context
}

// GetUserService new MGetUserService
func NewQueryRelationService(ctx context.Context) *QueryRelationService {
	return &QueryRelationService{ctx: ctx}
}

func (s *QueryRelationService) QueryRelation(relation db.Follow) (bool, error) {
	//先在Redis里找到它俩之间是否有关系--多个机器查询
	if flag, err := redis.Relation1.SIsMember(redis.Ctx, strconv.Itoa(int(relation.UserID)), relation.ToUserID).Result(); flag {
		//刷新过期时间
		redis.Relation1.Expire(redis.Ctx, strconv.Itoa(int(relation.UserID)), shared.RedisExpireTime)
		return true, err
	}
	if flag, err := redis.Relation2.SIsMember(redis.Ctx, strconv.Itoa(int(relation.UserID)), relation.ToUserID).Result(); flag {
		redis.Relation2.Expire(redis.Ctx, strconv.Itoa(int(relation.UserID)), shared.RedisExpireTime)
		return true, err
	}
	//若关系Redis没有，则取出UserId的关注列表进行对比比较
	//先去关注Redis看
	ids, _ := redis.Follower.MGet(redis.Ctx, strconv.Itoa(int(relation.UserID))).Result()
	if len(ids) != 0 {
		for _, v := range ids {
			if v.(int64) == relation.ToUserID {
				//加入缓存
				redis.AddRelation(relation.UserID, relation.ToUserID)
				return true, nil
			}
		}
		return false, nil
	}
	//再去sql看
	follows, err := dal.MGetFollowList(s.ctx, relation.UserID)
	if err != nil {
		return false, errno.QueryFollowErr
	}
	for _, v := range follows {
		if v.UserID == relation.ToUserID {
			//加入缓存
			redis.AddRelation(relation.UserID, relation.ToUserID)
			return true, nil
		}
	}
	return false, nil
}
