package service

import (
	"context"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/pack"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
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
	users := []*UserServer.User{}
	relations := []bool{}
	if len(UserVideos) != 0 { //这是感觉还可以再优化的地方
		//rpc调用拿取user信息
		for _, v := range UserVideos {
			user, relation := getUserInfo(s.ctx, v, req.UserId)
			users = append(users, user)
			relations = append(relations, relation)
		}
	} else {
		return videoList, nil //没有视频，也不传错误信息
	}

	videoList = pack.VideoList(UserVideos, users, relations)
	return videoList, nil
}
