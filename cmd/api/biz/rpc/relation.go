package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer/relationservice"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
	mw "github.com/xiaohei366/TinyTiktok/pkg/middleware"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

var relationClient relationservice.Client

func initRelation() {
	r, err := etcd.NewEtcdResolver([]string{shared.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(shared.ApiServiceName),
		provider.WithExportEndpoint(shared.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := relationservice.NewClient(
		shared.RelationServiceName,
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
	relationClient = c
}

// 查询两个用户之间的关系
func QueryRelation(ctx context.Context, req *RelationServer.DouyinQueryRelationRequest) (bool, error) {
	resp, err := relationClient.QueryRelation(ctx, req)
	if err != nil {
		return false, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return false, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.IsFollow, nil
}

// 传递关注用户请求，获取rpc响应
func FollowAction(ctx context.Context, req *RelationServer.DouyinRelationActionRequest) (resp *RelationServer.DouyinRelationActionResponse, err error) {
	resp, err = relationClient.RelationAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}

// 传递获取关注列表请求，获取rpc响应
func GetFollowList(ctx context.Context, req *RelationServer.DouyinRelationFollowListRequest) (resp *RelationServer.DouyinRelationFollowListResponse, err error) {
	resp, err = relationClient.MGetRelationFollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}

// 传递获取粉丝列表请求，获取rpc响应
func GetFollowerList(ctx context.Context, req *RelationServer.DouyinRelationFollowerListRequest) (resp *RelationServer.DouyinRelationFollowerListResponse, err error) {
	resp, err = relationClient.MGetUserRelationFollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}
