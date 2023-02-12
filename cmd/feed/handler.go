package main

import (
	"context"
	"github.com/xiaohei366/TinyTiktok/cmd/feed/kitex_gen/FeedServer"
	"github.com/xiaohei366/TinyTiktok/cmd/feed/service"
	"github.com/xiaohei366/TinyTiktok/cmd/feed/service/pack"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// FeedSrvImpl implements the last service interface defined in the IDL.
type FeedSrvImpl struct{}

// DouyinFeed implements the FeedSrvImpl interface.
func (s *FeedSrvImpl) DouyinFeed(ctx context.Context, req *FeedServer.DouyinFeedRequest) (resp *FeedServer.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	videos, err := service.NewFeedService(ctx).Feed(req)
	if err != nil {
		return pack.BuildBaseResp(err), nil
	}
	resp = pack.BuildBaseResp(errno.Success) // 把错误和成功的消息封装进去
	resp.VideoList = videos.VideoList
	return
}
