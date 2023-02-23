package service

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/xiaohei366/TinyTiktok/cmd/user/config"
	"github.com/xiaohei366/TinyTiktok/cmd/user/initialize/mq"
	"github.com/xiaohei366/TinyTiktok/cmd/user/initialize/redis"
	"github.com/xiaohei366/TinyTiktok/cmd/user/rpc"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type ChangeUserFollowCountService struct {
	ctx context.Context
}

// GetUserService new MGetUserService
func NewChangeUserFollowCountService(ctx context.Context) *ChangeUserFollowCountService {
	return &ChangeUserFollowCountService{ctx: ctx}
}

// 根据userId获得TableUser对象--这里不处理"是否关注(都默认false)"，放到relation模块处理
func (s *ChangeUserFollowCountService) ChangeUserFollowCount(userID int64, toUserID int64, isFollow bool) error {
	if userID == toUserID {
		return nil
	}
	//组装成一个消息，为发送至消息队列作准备
	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(int(userID)))
	sb.WriteString("&")
	sb.WriteString(strconv.Itoa(int(toUserID)))
	var err error
	flag, err := rpc.QueryRelation(context.Background(), &RelationServer.DouyinQueryRelationRequest{
		UserId:   userID,
		ToUserId: toUserID,
	})
	if err != nil {
		return err
	}
	//延迟双删，先删除缓存中的关系，再刷新数据库，再删除缓存中的数据---保证数据库和缓存的一致性
	if isFollow {
		//只有未关注的人才能加
		if flag {
			return nil
		}
		//step1
		redis.DelCount(userID, toUserID)
		//step2
		err = mq.AddActor.Publish(context.Background(), sb.String())
		//step3
		time.Sleep(config.SleepTime)
		//step4
		redis.DelCount(userID, toUserID)
	} else {
		//只有关注的人才能减
		if !flag {
			return nil
		}
		//step1
		redis.DelCount(userID, toUserID)
		//step2
		err = mq.DelActor.Publish(context.Background(), sb.String())
		//step3
		time.Sleep(config.SleepTime)
		//step4
		redis.DelCount(userID, toUserID)
	}
	if err != nil {
		return errno.ChangeUserFollowCountErr
	}
	return nil
}
