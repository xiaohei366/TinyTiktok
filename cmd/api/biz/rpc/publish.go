package rpc

import (
	"context"
	mw "github.com/xiaohei366/TinyTiktok/pkg/middleware"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer/videosrv"

	"github.com/xiaohei366/TinyTiktok/pkg/errno"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"

	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var publishClient videosrv.Client

// initPublishRpc init the publishClient.
func initPublishRpc() {
	r, err := etcd.NewEtcdResolver([]string{shared.ETCDAddress})
	if err != nil {
		panic(err)
	}

	c, err := videosrv.NewClient(
		shared.VideoServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: shared.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	publishClient = c
}

// PublishVideos is used to publish videos by user.
func PublishVideos(ctx context.Context, req *VideoServer.DouyinPublishActionRequest) (*VideoServer.DouyinPublishActionResponse, error) {
	resp, err := publishClient.PublishAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}

// PublishList get the video list by user id or token.
func PublishList(ctx context.Context, req *VideoServer.DouyinPublishListRequest) ([]*VideoServer.Video, error) {
	resp, err := publishClient.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.VideoList, nil
}

// FeedVideos Get the videos by latestTime without user id.
func FeedVideos(ctx context.Context, req *VideoServer.DouyinFeedRequest) ([]*VideoServer.Video, error) {
	resp, err := publishClient.Feed(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.VideoList, nil
}
