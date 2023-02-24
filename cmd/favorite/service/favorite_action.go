package service

import (
	"context"
	"github.com/xiaohei366/TinyTiktok/cmd/favorite/initialize/mq"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/FavoriteServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
	"strconv"
	"strings"
)

type GetFavoriteService struct {
	ctx context.Context
}

func NewGetFavoriteService(ctx context.Context) *GetFavoriteService {
	return &GetFavoriteService{ctx: ctx}
}

// FavouriteAction 根据userId，videoId,actionType对视频进行点赞或者取消赞操作;
func (s *GetFavoriteService) FavouriteAction(req *FavoriteServer.DouyinFavoriteActionRequest) error {
	//组装成一个消息，为发送至消息队列作准备
	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(int(req.UserId)))
	sb.WriteString("&")
	sb.WriteString(strconv.Itoa(int(req.VideoId)))
	sb.WriteString("&")
	sb.WriteString(strconv.Itoa(int(req.ActionType)))

	//修改本地数据库中的关系
	err := mq.AddActor.Publish(context.Background(), sb.String())
	if err != nil {
		return errno.FavoriteActionErr
	}
	return nil
}
