package service

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/config"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/initialize/mq"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/initialize/redis"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/rpc"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type RelationActionService struct {
	ctx context.Context
}

func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

// 关注/取关操作---先保证user服务器是正常的，才能实行关注与否操作
func (s *RelationActionService) FollowAction(isFollow bool, relation db.Follow) error {
	if relation.UserID == relation.ToUserID {
		return nil
	}
	//组装成一个消息，为发送至消息队列作准备
	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(int(relation.UserID)))
	sb.WriteString("&")
	sb.WriteString(strconv.Itoa(int(relation.ToUserID)))
	//先尝试修改user模块的数据库
	err := rpc.ChangeFollowCount(s.ctx, &UserServer.DouyinChangeUserFollowRequest{
		UserId:   relation.UserID,
		ToUserId: relation.ToUserID,
		IsFollow: isFollow,
	})
	if err != nil {
		return errno.UserRPCErr
	}
	//成功后再修改本地数据库中的关系
	//延迟双删，先删除缓存中的关系，再刷新数据库，再删除缓存中的数据---保证数据库和缓存的一致性
	if isFollow {
		//step1
		redis.RedisWithDel(relation.UserID, relation.ToUserID)
		//step2
		err = mq.AddActor.Publish(context.Background(), sb.String())
		//step3
		time.Sleep(config.SleepTime)
		//step4
		redis.RedisWithDel(relation.UserID, relation.ToUserID)
	} else {
		//step1
		redis.RedisWithDel(relation.UserID, relation.ToUserID)
		//step2
		err = mq.DelActor.Publish(context.Background(), sb.String())
		//step3
		time.Sleep(config.SleepTime)
		//step4
		redis.RedisWithDel(relation.UserID, relation.ToUserID)
	}
	// 记录日志
	klog.Debug("消息打入成功。")
	if err != nil {
		return errno.FollowActionErr
	}
	return nil
}
