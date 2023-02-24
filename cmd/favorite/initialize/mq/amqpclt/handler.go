package amqpclt

import (
	"context"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/service/dal"
	"strconv"
	"strings"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
)

// 更改用户与视频的点赞关系
func (a *Actor) FavoriteActionAdd(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		// 取出user Id和视频video Id
		params := strings.Split(string(d.Body), "&")
		UserId, err := strconv.Atoi(params[0])
		if err != nil {
			klog.Errorf("转换错误：(%v)", err)
		}
		VideoId, err := strconv.Atoi(params[1])
		if err != nil {
			klog.Errorf("转换错误：(%v)", err)
		}
		ActionType, err := strconv.Atoi(params[2])
		if err != nil {
			klog.Errorf("转换错误：(%v)", err)
		}
		klog.Infof("调用数据库更新点赞操作(%v,%v,%v)", UserId, VideoId, ActionType)
		// 执行SQL，注必须scan，该SQL才能被执行。
		if err := dal.FavoriteAction(context.Background(), int64(UserId), int64(VideoId), int32(ActionType)); err != nil {
			// 执行出错，打印日志。
			klog.Errorf("调用数据库增加关注失败：(%v)", err)
		}
	}
}
