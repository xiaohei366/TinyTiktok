package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/publish/kitex_gen/PublishServer"
	"github.com/xiaohei366/TinyTiktok/cmd/publish/service"
	"github.com/xiaohei366/TinyTiktok/cmd/publish/service/pack"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// PublishSrvImpl implements the last service interface defined in the IDL.
type PublishSrvImpl struct{}

// PublishAction implements the PublishSrvImpl interface.
func (s *PublishSrvImpl) PublishAction(ctx context.Context, req *PublishServer.DouyinPublishActionRequest) (resp *PublishServer.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	_, err = service.NewPublishActionService(ctx).PublishAction(req) //似乎不需要提供这个信息。只需存储即可。
	if err != nil {
		return &PublishServer.DouyinPublishActionResponse{
			BaseResp: pack.BuildPublishBaseResp(err),
		}, nil
	}
	return &PublishServer.DouyinPublishActionResponse{
		BaseResp: pack.BuildPublishBaseResp(errno.Success),
	}, nil
}

// PublishList implements the PublishSrvImpl interface.
func (s *PublishSrvImpl) PublishList(ctx context.Context, req *PublishServer.DouyinPublishListRequest) (resp *PublishServer.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	klog.Info("handler/PublishList req:", req)
	userVideos, err := service.NewPublishListService(ctx).PublishList(req)
	if err != nil {
		return &PublishServer.DouyinPublishListResponse{
			BaseResp: pack.BuildPublishBaseResp(err),
		}, err
	}
	return &PublishServer.DouyinPublishListResponse{
		BaseResp:  pack.BuildPublishBaseResp(errno.Success),
		VideoList: userVideos,
	}, nil
}
