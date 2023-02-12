package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/kitex_gen/UserServer/userservice"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
	mw "github.com/xiaohei366/TinyTiktok/pkg/middleware"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
	//"github.com/kitex-contrib/obs-opentelemetry/provider"
	//"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{shared.ETCDAddress})
	if err != nil {
		panic(err)
	}
	// provider.NewOpenTelemetryProvider(
	// 	provider.WithServiceName(shared.ApiServiceName),
	// 	provider.WithExportEndpoint(shared.ExportEndpoint),
	// 	provider.WithInsecure(),
	// )
	c, err := userservice.NewClient(
		shared.UserServiceName,
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
	userClient = c
}

// 注册功能
func Register(ctx context.Context, req *UserServer.DouyinUserRegisterRequest) (int64, error) {
	resp, err := userClient.Register(ctx, req)
	if err != nil {
		return -1, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return -1, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserId, nil
}

// 登录功能
func Login(ctx context.Context, req *UserServer.DouyinUserLoginRequest) (int64, error) {
	resp, err := userClient.Login(ctx, req)
	if err != nil {
		return -1, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return -1, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserId, nil
}

func GetUserInfo(ctx context.Context, req *UserServer.DouyinUserRequest) (*UserServer.User, error) {
	resp, err := userClient.GetUserInfo(ctx, req)
	if err != nil {
		return &UserServer.User{}, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return &UserServer.User{}, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.User, nil //两种不同名但相同的结构体需要用指针来传递
}
