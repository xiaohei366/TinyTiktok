package main

import (
	"context"

	"github.com/xiaohei366/TinyTiktok/cmd/relation/initialize/db"
	RelationServer "github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/service"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/service/pack"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// RelationServerImpl implements the last service interface defined in the IDL.
type RelationServerImpl struct{}

// RelationAction implements the RelationServerImpl interface.
func (s *RelationServerImpl) RelationAction(ctx context.Context, req *RelationServer.DouyinRelationActionRequest) (resp *RelationServer.DouyinRelationActionResponse, err error) {
	relation := db.Follow{
		UserID:   req.UserId,
		ToUserID: req.ToUserId,
	}
	//检验格式
	if req.ActionType != 1 && req.ActionType != 2 {
		resp = pack.BuildrelationActionResp(errno.ActionTypeErr)
		return resp, nil
	}

	//进行关注操作
	err = service.NewRelationActionService(ctx).FollowAction(req.ActionType == 1, relation)
	if err != nil {
		resp = pack.BuildrelationActionResp(err)
		return resp, nil
	}
	resp = pack.BuildrelationActionResp(errno.Success)
	return resp, nil
}

// MGetRelationFollowList implements the RelationServerImpl interface.
func (s *RelationServerImpl) MGetRelationFollowList(ctx context.Context, req *RelationServer.DouyinRelationFollowListRequest) (resp *RelationServer.DouyinRelationFollowListResponse, err error) {
	// 调用相应服务即可
	users, err := service.NewMGetUserRelationFollowerService(ctx).MGetUserRelationFollower(req.UserId)
	if err != nil {
		resp = pack.BuildgetFollowListResp(err, users)
		return resp, nil
	}
	resp = pack.BuildgetFollowListResp(errno.Success, users)
	return resp, nil
}

// MGetUserRelationFollowerList implements the RelationServerImpl interface.
func (s *RelationServerImpl) MGetUserRelationFollowerList(ctx context.Context, req *RelationServer.DouyinRelationFollowerListRequest) (resp *RelationServer.DouyinRelationFollowerListResponse, err error) {
	// 调用相应服务即可
	users, err := service.NewMGetUserRelationFollowerService(ctx).MGetUserRelationFollower(req.UserId)
	if err != nil {
		resp = pack.BuildgetFollowerListResp(err, users)
		return resp, nil
	}
	resp = pack.BuildgetFollowerListResp(errno.Success, users)
	return resp, nil
}

// QueryRelation implements the RelationServerImpl interface.
func (s *RelationServerImpl) QueryRelation(ctx context.Context, req *RelationServer.DouyinQueryRelationRequest) (resp *RelationServer.DouyinQueryRelationResponse, err error) {
	//先封装进结构体
	relation := db.Follow{
		UserID:   req.UserId,
		ToUserID: req.ToUserId,
	}

	//进行查询操作
	isFollow, err := service.NewQueryRelationService(ctx).QueryRelation(relation)
	if err != nil {
		resp = pack.BuildrelationQueryResp(err, false)
		return resp, nil
	}
	resp = pack.BuildrelationQueryResp(errno.Success, isFollow)
	return resp, nil
}
