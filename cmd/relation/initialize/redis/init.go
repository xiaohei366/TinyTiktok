package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

//上下文的作用是允许任务中断：任务中断之后，处理器保存上下文，以便之后根据上下文在任务的同一位置继续执行

var Ctx = context.Background()

// 粉丝列表信息
var Follower *redis.Client

// 关注列表信息信息
var Follow *redis.Client

// 下面两个为了解决热Key而使用
var Relation1 *redis.Client
var Relation2 *redis.Client

func InitRedis() {

	Follower = redis.NewClient(&redis.Options{
		Addr:     shared.RedisAddr,
		Password: shared.RedisPassword,
		DB:       shared.RedisFollower,
	})
	Follow = redis.NewClient(&redis.Options{
		Addr:     shared.RedisAddr,
		Password: shared.RedisPassword,
		DB:       shared.RedisFollow,
	})
	Relation1 = redis.NewClient(&redis.Options{
		Addr:     shared.RedisAddr,
		Password: shared.RedisPassword,
		DB:       shared.RedisRelation1,
	})
	Relation2 = redis.NewClient(&redis.Options{
		Addr:     shared.RedisAddr,
		Password: shared.RedisPassword,
		DB:       shared.RedisRelation2, // 当前用户是否关注了自己粉丝信息.
	})
}
