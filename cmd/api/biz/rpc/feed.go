package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/kitex_gen/FeedServer"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/kitex_gen/FeedServer/feedsrv"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
	mw "github.com/xiaohei366/TinyTiktok/pkg/middleware"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
	//"github.com/kitex-contrib/obs-opentelemetry/provider"
	//"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var feedClient feedsrv.Client

func initFeed() {
	r, err := etcd.NewEtcdResolver([]string{shared.ETCDAddress})
	if err != nil {
		panic(err)
	}
	// provider.NewOpenTelemetryProvider(
	// 	provider.WithServiceName(shared.ApiServiceName),
	// 	provider.WithExportEndpoint(shared.ExportEndpoint),
	// 	provider.WithInsecure(),
	// )
	c, err := feedsrv.NewClient(
		shared.FeedServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		//client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: shared.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	feedClient = c
}

// CreateNote create note info //这些后面的全部要改
func FeedVideos(ctx context.Context, req *FeedServer.DouyinFeedRequest) ([]*FeedServer.Video, error) {
	resp, err := feedClient.DouyinFeed(ctx, req) //这边是有点疑问的
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.VideoList, nil
}
