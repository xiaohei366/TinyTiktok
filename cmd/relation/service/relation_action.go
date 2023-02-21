package service

import (
	"context"
	"strconv"
	"strings"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/initialize/mq"
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
	if isFollow {
		err = mq.AddActor.Publish(context.Background(), sb.String())
	} else {
		err = mq.DelActor.Publish(context.Background(), sb.String())
	}
	// 记录日志
	klog.Debug("消息打入成功。")
	if err != nil {
		return errno.FollowActionErr
	}
	return nil
}
