package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/feed/kitex_gen/FeedServer"
	"github.com/xiaohei366/TinyTiktok/cmd/publish/kitex_gen/PublishServer"
	"github.com/xiaohei366/TinyTiktok/cmd/publish/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/publish/service/pack"
)

type PublishListService struct {
	ctx context.Context
}

func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{ctx: ctx}
}

// 这个功能写好了
func (s *PublishListService) PublishList(req *PublishServer.DouyinPublishListRequest) ([]*FeedServer.Video, error) {
	klog.Info("service/PublishList Req:", req)
	UserVideos, err := dal.MGetUserVideos(s.ctx, req.UserId) //这个ctx实际没用到，后续改。
	if err != nil {
		return nil, err
	}
	return pack.PublishList(UserVideos), nil //这边pack还要改，改成只返回videosList的格式。后续每个服务再自己封装就行了
}
