package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

//上下文的作用是允许任务中断：任务中断之后，处理器保存上下文，以便之后根据上下文在任务的同一位置继续执行

var Ctx = context.Background()

// 用户id和用户名对应信息表
var Name1 *redis.Client
var Name2 *redis.Client
// 下面两个为了解决热Key而使用
var Count1 *redis.Client
var Count2 *redis.Client
const FollowField = "FollowNum"
const FollowerField = "FollowerNum"


func InitRedis() {

	Name1 = redis.NewClient(&redis.Options{
		Addr:     shared.RedisAddr,
		Password: shared.RedisPassword,
		DB:       shared.RedisName1,
	})
	Name2 = redis.NewClient(&redis.Options{
		Addr:     shared.RedisAddr,
		Password: shared.RedisPassword,
		DB:       shared.RedisName2,
	})
	Count1 = redis.NewClient(&redis.Options{
		Addr:     shared.RedisAddr,
		Password: shared.RedisPassword,
		DB:       shared.RedisCount1,
	})
	Count2 = redis.NewClient(&redis.Options{
		Addr:     shared.RedisAddr,
		Password: shared.RedisPassword,
		DB:       shared.RedisCount2, // 当前用户是否关注了自己粉丝信息.
	})
}
