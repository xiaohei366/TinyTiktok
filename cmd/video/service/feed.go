package service

import (
	"context"
	"github.com/xiaohei366/TinyTiktok/cmd/video/kitex_gen/VideoServer"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/pack"
	"time"
)

type FeedService struct {
	ctx context.Context
}

// NewFeedService new FeedService
func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

// Feed used for feed service, get videos by latestTime.
func (s *FeedService) Feed(req *VideoServer.DouyinFeedRequest) (res *VideoServer.DouyinFeedResponse, err error) {
	var latestTime *int64
	if &req.LatestTime == nil || req.LatestTime == 0 {
		cur_time := int64(time.Now().UnixMilli())
		latestTime = &cur_time
	}
	feedModels, err := dal.MGetVideos(s.ctx, latestTime)
	if err != nil {
		return nil, err
	}
	feeds := pack.VideoList(feedModels)

	return &VideoServer.DouyinFeedResponse{
		VideoList: feeds,
	}, nil
}
