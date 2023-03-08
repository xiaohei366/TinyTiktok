package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

//上下文的作用是允许任务中断：任务中断之后，处理器保存上下文，以便之后根据上下文在任务的同一位置继续执行

var Ctx = context.Background()
const TextField = "CommentText"
const UserField = "UserID"

// 下面两个为了解决热Key而使用
var VedioToComment *redis.Client//vedio->comment_id
var UserList *redis.Client //comment_id->user_id--发表评论的用户
var CommentList1 *redis.Client // comment_id ->text
var CommentList2 *redis.Client // 解决大Key而生
func InitRedis() {
	VedioToComment = redis.NewClient(&redis.Options{
		Addr:     shared.RedisAddr,
		Password: shared.RedisPassword,
		DB:       shared.RedisVedioToCommentList,
	})
	UserList = redis.NewClient(&redis.Options{
		Addr:     shared.RedisAddr,
		Password: shared.RedisPassword,
		DB:       shared.RedisCommentUesrList,
	})

	CommentList1 = redis.NewClient(&redis.Options{
		Addr:     shared.RedisAddr,
		Password: shared.RedisPassword,
		DB:       shared.RedisCommentList1,
	})
	CommentList2 = redis.NewClient(&redis.Options{
		Addr:     shared.RedisAddr,
		Password: shared.RedisPassword,
		DB:       shared.RedisCommentList2,
	})
}
