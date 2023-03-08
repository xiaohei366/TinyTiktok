package service

import (
	"context"
	"time"

	"github.com/xiaohei366/TinyTiktok/cmd/video/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/pack"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
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
func (s *FeedService) Feed(req *VideoServer.DouyinFeedRequest) (videos []*VideoServer.Video, nextTime int64, err error) {
	var latestTime *int64
	if req == nil {
		return nil, 0, nil
	}
	if req.LatestTime == 0 {
		cur_time := int64(time.Now().UnixMilli())
		latestTime = &cur_time
	}
	//查询视频
	feedModels, _ := dal.MGetVideos(s.ctx, latestTime)
	if len(feedModels) == 0 {
		nextTime = time.Now().UnixMilli()
		return videos, nextTime, nil
	} else {
		nextTime = feedModels[len(feedModels)-1].UpdatedAt.UnixMilli()
	}
	videos = pack.VideoList(feedModels, req.UserId)
	return videos, nextTime, nil
}
