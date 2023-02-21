package service

import (
	"context"
	"time"

	"github.com/xiaohei366/TinyTiktok/cmd/video/rpc"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"

	"github.com/xiaohei366/TinyTiktok/cmd/video/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/pack"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
)

type FeedService struct {
	ctx context.Context
}

// NewFeedService new FeedService
func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

// Feed used for feed service, get videos by latestTime.
func (s *FeedService) Feed(req *VideoServer.DouyinFeedRequest) (videos []*VideoServer.Video, nextTime int64, err error) {
	var latestTime *int64
	if req.LatestTime == 0 {
		cur_time := int64(time.Now().UnixMilli())
		latestTime = &cur_time
	}
	feedModels, err := dal.MGetVideos(s.ctx, latestTime)
	if len(feedModels) == 0 {
		nextTime = time.Now().UnixMilli()
		return videos, nextTime, err
	} else {
		nextTime = feedModels[len(feedModels)-1].UpdatedAt.UnixMilli()
	}
	//rpc 拿取视频信息
	users := []*UserServer.User{}
	for _, v := range feedModels {
		user, err := rpc.GetUserInfo(s.ctx, &UserServer.DouyinUserRequest{
			UserId: v.AuthorID,
		})

		if err != nil {
			return videos, nextTime, errno.UserRPCErr
		}
		users = append(users, user)
	}

	//queryRelation
	relations := []bool{}
	for _, u := range users {
		relation, err := rpc.QueryRelation(s.ctx, &RelationServer.DouyinQueryRelationRequest{
			UserId:   req.UserId,
			ToUserId: u.Id,
		})
		if err != nil {
			return videos, nextTime, errno.RelationRPCErr
		}
		relations = append(relations, relation)
	}

	videos = pack.VideoList(feedModels, users, relations)

	return videos, nextTime, nil
}
