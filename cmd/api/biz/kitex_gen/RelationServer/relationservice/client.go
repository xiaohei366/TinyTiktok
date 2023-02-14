// Code generated by Kitex v0.4.4. DO NOT EDIT.

package relationservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	RelationServer "github.com/xiaohei366/TinyTiktok/cmd/api/biz/kitex_gen/RelationServer"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	RelationAction(ctx context.Context, Req *RelationServer.DouyinRelationActionRequest, callOptions ...callopt.Option) (r *RelationServer.DouyinRelationActionResponse, err error)
	MGetRelationFollowList(ctx context.Context, Req *RelationServer.DouyinRelationFollowListRequest, callOptions ...callopt.Option) (r *RelationServer.DouyinRelationFollowListResponse, err error)
	MGetUserRelationFollowerList(ctx context.Context, Req *RelationServer.DouyinRelationFollowerListRequest, callOptions ...callopt.Option) (r *RelationServer.DouyinRelationFollowerListResponse, err error)
	QueryRelation(ctx context.Context, Req *RelationServer.DouyinQueryRelationRequest, callOptions ...callopt.Option) (r *RelationServer.DouyinQueryRelationResponse, err error)
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
	return &kRelationServiceClient{
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

type kRelationServiceClient struct {
	*kClient
}

func (p *kRelationServiceClient) RelationAction(ctx context.Context, Req *RelationServer.DouyinRelationActionRequest, callOptions ...callopt.Option) (r *RelationServer.DouyinRelationActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationAction(ctx, Req)
}

func (p *kRelationServiceClient) MGetRelationFollowList(ctx context.Context, Req *RelationServer.DouyinRelationFollowListRequest, callOptions ...callopt.Option) (r *RelationServer.DouyinRelationFollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MGetRelationFollowList(ctx, Req)
}

func (p *kRelationServiceClient) MGetUserRelationFollowerList(ctx context.Context, Req *RelationServer.DouyinRelationFollowerListRequest, callOptions ...callopt.Option) (r *RelationServer.DouyinRelationFollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MGetUserRelationFollowerList(ctx, Req)
}

func (p *kRelationServiceClient) QueryRelation(ctx context.Context, Req *RelationServer.DouyinQueryRelationRequest, callOptions ...callopt.Option) (r *RelationServer.DouyinQueryRelationResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryRelation(ctx, Req)
}
