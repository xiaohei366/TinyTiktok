package main

import (
	"context"
	"github.com/xiaohei366/TinyTiktok/cmd/video/kitex_gen/VideoServer"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/pack"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// VideoSrvImpl implements the last service interface defined in the IDL.
type VideoSrvImpl struct{}

// Feed implements the VideoSrvImpl interface.
func (s *VideoSrvImpl) Feed(ctx context.Context, req *VideoServer.DouyinFeedRequest) (resp *VideoServer.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	videos, err := service.NewFeedService(ctx).Feed(req)
	if err != nil {
		return &VideoServer.DouyinFeedResponse{
			BaseResp: pack.BuildBaseResp(err),
		}, nil
	}
	return &VideoServer.DouyinFeedResponse{
		BaseResp:  pack.BuildBaseResp(errno.Success),
		VideoList: videos.VideoList,
	}, nil
}

// PublishAction implements the VideoSrvImpl interface.
func (s *VideoSrvImpl) PublishAction(ctx context.Context, req *VideoServer.DouyinPublishActionRequest) (resp *VideoServer.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	_, err = service.NewPublishActionService(ctx).PublishAction(req) //似乎不需要提供这个信息。只需存储即可。
	if err != nil {
		return &VideoServer.DouyinPublishActionResponse{
			BaseResp: pack.BuildBaseResp(err),
		}, nil
	}
	return &VideoServer.DouyinPublishActionResponse{
		BaseResp: pack.BuildBaseResp(errno.Success),
	}, nil
}

// PublishList implements the VideoSrvImpl interface.
func (s *VideoSrvImpl) PublishList(ctx context.Context, req *VideoServer.DouyinPublishListRequest) (resp *VideoServer.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	userVideos, err := service.NewPublishListService(ctx).PublishList(req)
	if err != nil {
		return &VideoServer.DouyinPublishListResponse{
			BaseResp: pack.BuildBaseResp(err),
		}, err
	}
	return &VideoServer.DouyinPublishListResponse{
		BaseResp:  pack.BuildBaseResp(errno.Success),
		VideoList: userVideos,
	}, nil
}
