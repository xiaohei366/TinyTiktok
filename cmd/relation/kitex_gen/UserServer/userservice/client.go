// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	UserServer "github.com/xiaohei366/TinyTiktok/cmd/relation/kitex_gen/UserServer"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Register(ctx context.Context, Req *UserServer.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *UserServer.DouyinUserRegisterResponse, err error)
	Login(ctx context.Context, Req *UserServer.DouyinUserLoginRequest, callOptions ...callopt.Option) (r *UserServer.DouyinUserLoginResponse, err error)
	GetUserInfo(ctx context.Context, Req *UserServer.DouyinUserRequest, callOptions ...callopt.Option) (r *UserServer.DouyinUserResponse, err error)
	MGetUserInfo(ctx context.Context, Req *UserServer.DouyinMUserRequest, callOptions ...callopt.Option) (r *UserServer.DouyinMUserResponse, err error)
	ChangeUserFollowCount(ctx context.Context, Req *UserServer.DouyinChangeUserFollowRequest, callOptions ...callopt.Option) (r *UserServer.BaseResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) Register(ctx context.Context, Req *UserServer.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *UserServer.DouyinUserRegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, Req)
}

func (p *kUserServiceClient) Login(ctx context.Context, Req *UserServer.DouyinUserLoginRequest, callOptions ...callopt.Option) (r *UserServer.DouyinUserLoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, Req)
}

func (p *kUserServiceClient) GetUserInfo(ctx context.Context, Req *UserServer.DouyinUserRequest, callOptions ...callopt.Option) (r *UserServer.DouyinUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, Req)
}

func (p *kUserServiceClient) MGetUserInfo(ctx context.Context, Req *UserServer.DouyinMUserRequest, callOptions ...callopt.Option) (r *UserServer.DouyinMUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MGetUserInfo(ctx, Req)
}

func (p *kUserServiceClient) ChangeUserFollowCount(ctx context.Context, Req *UserServer.DouyinChangeUserFollowRequest, callOptions ...callopt.Option) (r *UserServer.BaseResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChangeUserFollowCount(ctx, Req)
}
