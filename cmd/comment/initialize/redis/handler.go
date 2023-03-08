package redis

import (
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

func AddComment(videoId int64, userid int64, text string, id int64, t time.Time) bool {
	mid := len(text) / 2
	// 将查询到的数量修改至Redis.---若键值存在，自动修改
	Mem := &redis.Z{
		Score:  float64(t.UnixNano()),
		Member: id,
	}
	CommentList1.Set(Ctx, strconv.Itoa(int(id)), text[:mid], shared.RedisExpireTime)
	CommentList2.Set(Ctx, strconv.Itoa(int(id)), text[mid:], shared.RedisExpireTime)

	VedioToComment.ZAdd(Ctx, strconv.Itoa(int(videoId)), Mem)
	VedioToComment.Expire(Ctx, strconv.Itoa(int(videoId)), shared.RedisExpireTime)
	UserList.Set(Ctx, strconv.Itoa(int(id)), userid, shared.RedisExpireTime)
	return true
}

func DelComment(videoId int64, commentId int64) (bool, error) {
	VedioIdStr := strconv.Itoa(int(videoId))
	CommentIdStr := strconv.Itoa(int(videoId))
	//scard计算集合大小,SRem移除关系
	CommentList1.Del(Ctx, CommentIdStr)
	CommentList2.Del(Ctx, CommentIdStr)
	UserList.Del(Ctx, CommentIdStr)

	VedioToComment.ZRem(Ctx, VedioIdStr)
	VedioToComment.Expire(Ctx, VedioIdStr, 0)
	return true, nil
}
