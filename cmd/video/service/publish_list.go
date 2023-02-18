package service

import (
	"context"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/pack"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
)

type PublishListService struct {
	ctx context.Context
}

// NewPublishListService new PublishListService
func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{ctx: ctx}
}

// PublishList get the videoList by user id.
func (s *PublishListService) PublishList(req *VideoServer.DouyinPublishListRequest) ([]*VideoServer.Video, error) {
	UserVideos, err := dal.MGetUserVideos(s.ctx, req.UserId) //这个ctx实际没用到，后续改。
	if err != nil {
		return nil, err
	}
	videos, err := pack.Videos(s.ctx, UserVideos, req.UserId)
	if err != nil {
		return nil, err
	}
	return videos, nil //这边pack还要改，改成只返回videosList的格式。后续每个服务再自己封装就行了
}
