package redis

import (
	"strconv"

	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

func AddUserLikeList(userId int64, videoId int64) {
	//更新缓存
	UserLikeList.SAdd(Ctx, strconv.Itoa(int(userId)), videoId)
	// 更新过期时间。
	UserLikeList.Expire(Ctx, strconv.Itoa(int(userId)), shared.RedisExpireTime)

}

// 当取消点赞时，更新redis里的信息
func DelUserLikeList(userId int64, videoId int64) (bool, error) {
	UserIdStr := strconv.Itoa(int(userId))
	// 删除粉丝缓存的两者关系
	//scard计算集合大小,SRem移除关系
	// 删除关注缓存的两者关系
	if exist, _ := UserLikeList.SIsMember(Ctx, UserIdStr, videoId).Result(); exist {
		UserLikeList.SRem(Ctx, UserIdStr, videoId)
		UserLikeList.Expire(Ctx, UserIdStr, shared.RedisExpireTime)
	}

	return true, nil
}
