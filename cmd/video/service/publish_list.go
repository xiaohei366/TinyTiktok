package service

import (
	"context"

	"github.com/xiaohei366/TinyTiktok/cmd/video/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/pack"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type PublishListService struct {
	ctx context.Context
}

// NewPublishListService new PublishListService
func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{ctx: ctx}
}

// PublishList get the videoList by user id.
func (s *PublishListService) PublishList(req *VideoServer.DouyinPublishListRequest) (videoList []*VideoServer.Video, err error) {
	UserVideos, err := dal.MGetUserVideos(s.ctx, req.ToUserId)
	if err != nil {
		return nil, errno.PublishListErr
	}
	videoList = pack.VideoList(UserVideos, req.ToUserId)
	return videoList, nil
}
