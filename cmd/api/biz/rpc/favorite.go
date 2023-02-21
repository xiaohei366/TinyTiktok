package rpc

import (
	"context"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/FavoriteServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/FavoriteServer/favoriteservice"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
	mw "github.com/xiaohei366/TinyTiktok/pkg/middleware"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
	//"github.com/kitex-contrib/obs-opentelemetry/provider"
	//"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var favoriteClient favoriteservice.Client

func initFavorite() {
	r, err := etcd.NewEtcdResolver([]string{shared.ETCDAddress})
	if err != nil {
		panic(err)
	}
	// provider.NewOpenTelemetryProvider(
	// 	provider.WithServiceName(shared.ApiServiceName),
	// 	provider.WithExportEndpoint(shared.ExportEndpoint),
	// 	provider.WithInsecure(),
	// )
	c, err := favoriteservice.NewClient(
		shared.FavoriteServiceName,
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
	favoriteClient = c
}

// 传递点赞操作请求，获取rpc响应
func FavoriteAction(ctx context.Context, req *FavoriteServer.DouyinFavoriteActionRequest) (resp *FavoriteServer.DouyinFavoriteActionResponse, err error) {
	resp, err = favoriteClient.FavoriteAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}

func GetFavoriteList(ctx context.Context, req *FavoriteServer.DouyinFavoriteListRequest) (resp *FavoriteServer.DouyinFavoriteListResponse, err error) {
	resp, err = favoriteClient.GetFavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}

func GetFavoriteUser(ctx context.Context, req *FavoriteServer.DouyinUserFavoriteRequest) (resp *FavoriteServer.DouyinUserFavoriteResponse, err error) {
	resp, err = favoriteClient.GetFavoriteUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}

func GetFavoriteVideo(ctx context.Context, req *FavoriteServer.DouyinVideoFavoriteRequest) (resp *FavoriteServer.DouyinVideoFavoriteResponse, err error) {
	resp, err = favoriteClient.GetFavoriteVideo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}
