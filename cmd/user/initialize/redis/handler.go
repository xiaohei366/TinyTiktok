package redis

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

func UpdateCount(UserId int64, FollowNum int64, FollowerNum int64) {
	//将热key分散到不同的服务器中
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(2) // 有几个机器是几

	switch id {
	case 0:
		// 将查询到的数量修改至Redis.---若键值存在，自动修改
		Count1.HSet(Ctx, strconv.Itoa(int(UserId)), FollowField, FollowNum)
		Count1.HSet(Ctx, strconv.Itoa(int(UserId)), FollowerField, FollowerNum)
		// 更新过期时间。
		Count1.Expire(Ctx, strconv.Itoa(int(UserId)), shared.RedisExpireTime)
	case 1:
		// 将查询到的数量修改至Redis.---若键值存在，自动修改
		Count2.HSet(Ctx, strconv.Itoa(int(UserId)), FollowField, FollowNum)
		Count2.HSet(Ctx, strconv.Itoa(int(UserId)), FollowerField, FollowerNum)
		// 更新过期时间。
		Count2.Expire(Ctx, strconv.Itoa(int(UserId)), shared.RedisExpireTime)
	}

}

func AddName(userId int64, name string) {
	//string类型数据操作
	Name.Set(Ctx, strconv.Itoa(int(userId)), name, shared.RedisExpireTime)
}

// 删除UserId的关注数和ToUserId的粉丝数
func DelCount(UserId int64, ToUserId int64) (bool, error) {
	UserIdStr := strconv.Itoa(int(UserId))
	ToUserIdStr := strconv.Itoa(int(ToUserId))
	// 删除粉丝缓存的两者关系
	//scard计算集合大小,SRem移除关系
	if ok, _ := Count1.HExists(Ctx, UserIdStr, FollowField).Result(); ok {
		Count1.HDel(Ctx, UserIdStr, FollowField)
		Count1.HDel(Ctx, ToUserIdStr, FollowerField)
		Count1.Expire(Ctx, UserIdStr, shared.RedisExpireTime)
		Count1.Expire(Ctx, ToUserIdStr, shared.RedisExpireTime)
	}
	if ok, _ := Count2.HExists(Ctx, UserIdStr, FollowField).Result(); ok {
		Count2.HDel(Ctx, UserIdStr, FollowField)
		Count2.HDel(Ctx, ToUserIdStr, FollowerField)
		Count2.Expire(Ctx, UserIdStr, shared.RedisExpireTime)
		Count2.Expire(Ctx, ToUserIdStr, shared.RedisExpireTime)
	}
	return true, nil
}
