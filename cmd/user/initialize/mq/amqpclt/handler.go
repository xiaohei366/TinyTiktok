package amqpclt

import (
	"context"
	"strconv"
	"strings"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
	"github.com/xiaohei366/TinyTiktok/cmd/user/service/dal"
)

// 数量添加的消费方式。
func (a *Actor) FollowAdd(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		// 取出两个人的id
		params := strings.Split(string(d.Body), "&")
		UserId, err := strconv.Atoi(params[0])
		if err != nil {
			klog.Errorf("转换错误：(%v)", err)
		}
		ToUserId, err := strconv.Atoi(params[1])
		if err != nil {
			klog.Errorf("转换错误：(%v)", err)
		}
		klog.Infof("调用数据库增加关注数和粉丝数操作(%v,%v)", UserId, ToUserId)
		if err := dal.IncreaseFollowCount(context.Background(), int64(UserId), int64(ToUserId)); err != nil {
			// 执行出错，打印日志。
			klog.Errorf("调用数据库增加关注数和粉丝数失败：(%v)", err)
		}
	}
}

// 数量删除的消费方式。
func (a *Actor) FollowDel(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		// 取出两个人的id
		params := strings.Split(string(d.Body), "&")
		UserId, err := strconv.Atoi(params[0])
		if err != nil {
			klog.Errorf("转换错误：(%v)", err)
		}
		ToUserId, err := strconv.Atoi(params[1])
		if err != nil {
			klog.Errorf("转换错误：(%v)", err)
		}
		klog.Infof("调用数据库减少关注和粉丝数操作(%v,%v)", UserId, ToUserId)
		if err := dal.DecreaseFollowCount(context.Background(), int64(UserId), int64(ToUserId)); err != nil {
			// 执行出错，打印日志。
			klog.Errorf("调用数据库减少关注和粉丝数失败：(%v)", err)
		}
	}
}
