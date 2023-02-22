package redis

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/xiaohei366/TinyTiktok/cmd/relation/config"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

func AddRelation(UserId int64, ToUserId int64) {
	//将热key分散到不同的服务器中
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(2) // 有几个机器是几

	switch id {
	case 0:
		// 将查询到的关注关系注入Redis.---若键值存在，自动不增加
		Relation1.SAdd(Ctx, strconv.Itoa(int(UserId)), ToUserId)
		// 更新过期时间。
		Relation1.Expire(Ctx, strconv.Itoa(int(UserId)), shared.RedisExpireTime)
	case 1:
		// 将查询到的关注关系注入Redis.
		Relation2.SAdd(Ctx, strconv.Itoa(int(UserId)), ToUserId)
		// 更新过期时间。
		Relation2.Expire(Ctx, strconv.Itoa(int(UserId)), shared.RedisExpireTime)
	}

}

func AddFollower(userId int64, ids []int64) {
	//更新粉丝列表缓存
	for _, id := range ids {
		Follower.SAdd(Ctx, strconv.Itoa(int(userId)), id)
		//AddRelation(id, userId)
	}
	// 更新follower的过期时间。
	Follower.Expire(Ctx, strconv.Itoa(int(userId)), shared.RedisExpireTime)

}

func AddFollow(userId int64, ids []int64) {
	//不是所有粉丝都加进去，如果粉丝大于N，才加入至缓存。
	if len(ids) < int(config.FansNum) {
		return
	}
	//更新粉丝列表缓存&关系缓存
	for _, id := range ids {
		Follow.SAdd(Ctx, strconv.Itoa(int(userId)), id)
		//AddRelation(id, userId)
	}
	// 更新follow的过期时间
	Follow.Expire(Ctx, strconv.Itoa(int(userId)), shared.RedisExpireTime)
}

// 当取关时，更新redis里的信息
func RedisWithDel(UserId int64, ToUserId int64) (bool, error) {
	UserIdStr := strconv.Itoa(int(UserId))
	ToUserIdStr := strconv.Itoa(int(ToUserId))
	// 删除粉丝缓存的两者关系
	//scard计算集合大小,SRem移除关系
	if cnt, _ := Follower.SCard(Ctx, ToUserIdStr).Result(); cnt != 0 {
		Follower.SRem(Ctx, ToUserIdStr, UserId)
		Follower.Expire(Ctx, ToUserIdStr, shared.RedisExpireTime)
	}
	// 删除关注缓存的两者关系
	if cnt, _ := Follow.SCard(Ctx, UserIdStr).Result(); cnt != 0 {
		Follow.SRem(Ctx, UserIdStr, ToUserId)
		Follow.Expire(Ctx, UserIdStr, shared.RedisExpireTime)
	}
	// 删除关系缓存的两者关系
	if cnt, _ := Relation1.Exists(Ctx, UserIdStr).Result(); cnt != 0 {
		Relation1.SRem(Ctx, UserIdStr, ToUserId)
		Relation1.Expire(Ctx, UserIdStr, shared.RedisExpireTime)
	}
	if cnt, _ := Relation2.Exists(Ctx, UserIdStr).Result(); cnt != 0 {
		Relation2.SRem(Ctx, UserIdStr, ToUserId)
		Relation2.Expire(Ctx, UserIdStr, shared.RedisExpireTime)
	}
	return true, nil
}
