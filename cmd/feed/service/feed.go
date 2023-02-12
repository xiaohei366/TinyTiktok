package service

import (
	"context"
	"github.com/xiaohei366/TinyTiktok/cmd/feed/kitex_gen/FeedServer"
	"github.com/xiaohei366/TinyTiktok/cmd/feed/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/feed/service/pack"
	"time"
)

type FeedService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

// 后面写服务就好了。
func (s *FeedService) Feed(req *FeedServer.DouyinFeedRequest) (res *FeedServer.DouyinFeedResponse, err error) {
	var latestTime *int64
	if &req.LatestTime == nil || req.LatestTime == 0 {
		cur_time := int64(time.Now().UnixMilli())
		latestTime = &cur_time
	}
	feedModels, err := dal.MGetVideos(s.ctx, latestTime)
	if err != nil {
		return nil, err
	}
	feeds := pack.FeedList(feedModels) //feed.videos自己封装成DouyinFeed就行了

	return &FeedServer.DouyinFeedResponse{ //OK,搞定
		VideoList: feeds,
	}, nil
}
