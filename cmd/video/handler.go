package main

import (
	"context"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/pack"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
)

// VideoSrvImpl implements the last service interface defined in the IDL.
type VideoSrvImpl struct{}

// Feed implements the VideoSrvImpl interface.
func (s *VideoSrvImpl) Feed(ctx context.Context, req *VideoServer.DouyinFeedRequest) (resp *VideoServer.DouyinFeedResponse, err error) {
	// TODO: Your code here...

	videos, nextTime, err := service.NewFeedService(ctx).Feed(req)
	if err != nil {
		resp = pack.BuildFeedResp(err, videos, nextTime)
		return resp, err
	}
	return pack.BuildFeedResp(nil, videos, nextTime), nil
}

// PublishAction implements the VideoSrvImpl interface.
func (s *VideoSrvImpl) PublishAction(ctx context.Context, req *VideoServer.DouyinPublishActionRequest) (resp *VideoServer.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...

	err = service.NewPublishActionService(ctx).PublishAction(req) //似乎不需要提供这个信息。只需存储即可。
	if err != nil {
		return pack.BuildPublishActionResp(err), nil
	}
	return pack.BuildPublishActionResp(nil), nil
}

// PublishList implements the VideoSrvImpl interface.
func (s *VideoSrvImpl) PublishList(ctx context.Context, req *VideoServer.DouyinPublishListRequest) (resp *VideoServer.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	userVideos, err := service.NewPublishListService(ctx).PublishList(req)
	if err != nil {
		return pack.BuildPublishListResp(err, userVideos), nil
	}
	return pack.BuildPublishListResp(nil, userVideos), nil
}
