package rpc

import (
	"context"

	"github.com/xiaohei366/TinyTiktok/cmd/relation/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/kitex_gen/UserServer/userservice"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
	mw "github.com/xiaohei366/TinyTiktok/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	//"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client 

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{shared.ETCDAddress})
	if err != nil {
		panic(err)
	}
	// provider.NewOpenTelemetryProvider(
	// 	provider.WithServiceName(shared.RelationServiceName),
	// 	provider.WithExportEndpoint(shared.ExportEndpoint),
	// 	provider.WithInsecure(),
	// )
	c, err := userservice.NewClient(
		shared.UserServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: shared.RelationServiceName}),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}



// 关注/取关操作---修改uesr数据库的数量
func ChangeFollowCount(ctx context.Context, req *UserServer.DouyinChangeUserFollowRequest) (error) {
	resp, err := userClient.ChangeUserFollowCount(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return nil
}

// 批量获得用户信息
func MGetUserInfo(ctx context.Context, req *UserServer.DouyinMUserRequest) ([]*UserServer.User, error) {
	resp, err := userClient.MGetUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.User, nil	
}