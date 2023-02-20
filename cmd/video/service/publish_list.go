package service

import (
	"context"
	"github.com/xiaohei366/TinyTiktok/cmd/video/rpc"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"

	"github.com/xiaohei366/TinyTiktok/cmd/video/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/pack"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
)

type PublishListService struct {
	ctx context.Context
}

// NewPublishListService new PublishListService
func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{ctx: ctx}
}

// PublishList get the videoList by user id.
func (s *PublishListService) PublishList(req *VideoServer.DouyinPublishListRequest) (videoList []*VideoServer.Video, err error) {
	UserVideos, err := dal.MGetUserVideos(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	//rpc调用拿取user信息
	users := []*UserServer.User{}
	for _, v := range UserVideos {
		user, err := rpc.GetUserInfo(s.ctx, &UserServer.DouyinUserRequest{
			UserId: v.AuthorID,
		})

		if err != nil {
			return videoList, errno.UserRPCErr
		}
		users = append(users, user)
	}

	//queryRelation
	relations := []bool{}
	for _, u := range users {
		relation, err := rpc.QueryRelation(s.ctx, &RelationServer.DouyinQueryRelationRequest{
			UserId:   u.Id,
			ToUserId: req.UserId,
		})
		if err != nil {
			return videoList, errno.RelationRPCErr
		}
		relations = append(relations, relation)
	}

	videoList = pack.VideoList(UserVideos, users, relations)
	return videoList, nil
}
