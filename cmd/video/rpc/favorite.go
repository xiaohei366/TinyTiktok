package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/FavoriteServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/FavoriteServer/favoriteservice"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
	mw "github.com/xiaohei366/TinyTiktok/pkg/middleware"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

var favoriteClient favoriteservice.Client

func initFavorite() {
	r, err := etcd.NewEtcdResolver([]string{shared.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(shared.VideoServiceName),
		provider.WithExportEndpoint(shared.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := favoriteservice.NewClient(
		shared.FavoriteServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		//client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: shared.VideoServiceName}),
	)
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}

// 拿取视频中点赞数量
func GetVideosFavoriteCount(ctx context.Context, req *FavoriteServer.DouyinVideoFavoriteRequest) (int64, error) {
	resp, err := favoriteClient.GetFavoriteVideo(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	fmt.Println("resp videoList")
	return resp.FavoriteCount, nil
}

// 查询user是否喜欢这个视频
func QueryUserLikeVideo(ctx context.Context, req *FavoriteServer.DouyinQueryFavoriteRequest) (bool, error) {
	resp, err := favoriteClient.QueryUserLikeVideo(ctx, req)
	if err != nil {
		return false, err
	}
	return resp.Favorite, nil
}
