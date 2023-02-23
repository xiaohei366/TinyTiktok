package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/CommentServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/CommentServer/commentservice"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
	mw "github.com/xiaohei366/TinyTiktok/pkg/middleware"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

var commentClient commentservice.Client

func initComment() {
	r, err := etcd.NewEtcdResolver([]string{shared.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(shared.VideoServiceName),
		provider.WithExportEndpoint(shared.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := commentservice.NewClient(
		shared.CommentServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: shared.VideoServiceName}),
	)
	if err != nil {
		panic(err)
	}
	commentClient = c
}

// 查询评论列表功能
func CommentList(ctx context.Context, req *CommentServer.DouyinCommentListRequest) (resp *CommentServer.DouyinCommentListResponse, err error) {
	resp, err = commentClient.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}

	return resp, nil
}
