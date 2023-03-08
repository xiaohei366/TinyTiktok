package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

//上下文的作用是允许任务中断：任务中断之后，处理器保存上下文，以便之后根据上下文在任务的同一位置继续执行

var Ctx = context.Background()



// 下面两个为了解决热Key而使用
var UserLikeList *redis.Client

func InitRedis() {

	UserLikeList = redis.NewClient(&redis.Options{
		Addr:     shared.RedisAddr,
		Password: shared.RedisPassword,
		DB:       shared.RedisUserLikeList,
	})
}
