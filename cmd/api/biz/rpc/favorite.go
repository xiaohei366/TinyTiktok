package rpc

import (
	"context"
	"fmt"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
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
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(shared.ApiServiceName),
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
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: shared.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}

// 传递点赞操作请求，获取rpc响应//这个暂时OK了。有消息了。
func FavoriteAction(ctx context.Context, req *FavoriteServer.DouyinFavoriteActionRequest) (resp *FavoriteServer.DouyinFavoriteActionResponse, err error) {
	fmt.Println("favoriteAction:", req.UserId, req.VideoId, req.ActionType)
	resp, err = favoriteClient.FavoriteAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}

func GetFavoriteList(ctx context.Context, req *FavoriteServer.DouyinFavoriteListRequest) ([]*FavoriteServer.Video, error) {
	fmt.Println("rpc get fav list:", req.UserId)
	resp, err := favoriteClient.GetFavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	if len(resp.VideoList) == 0 {
		return nil, nil
	}
	fmt.Println("resp videoList")
	return resp.VideoList, nil
}

// 拿取视频中点赞数量//这个可以用
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
