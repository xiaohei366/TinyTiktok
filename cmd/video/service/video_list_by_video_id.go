package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/video/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/pack"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
)

type VideoListByVideoIdService struct {
	ctx context.Context
}

func NewVideoListByVideoIdService(ctx context.Context) *VideoListByVideoIdService {
	return &VideoListByVideoIdService{ctx: ctx}
}

func (s *VideoListByVideoIdService) GetVideoListByVideoId(req *VideoServer.DouyinVideoListByVideoId) (videoL []*VideoServer.Video, err error) {
	var videos []*db.Video
	for _, vId := range req.VideoId {
		video, err := dal.GetVideosByVideosId(s.ctx, vId)
		if err != nil {
			klog.Info("video Id not found ")
		} else {
			videos = append(videos, video)
		}
	}
	users := []*UserServer.User{}
	relations := []bool{}
	if len(videos) != 0 { //这是感觉还可以再优化的地方
		//rpc调用拿取user信息
		for _, v := range videos {
			user, relation := getUserInfo(s.ctx, v, req.UserId)
			users = append(users, user)
			relations = append(relations, relation)
		}
	} else {
		return videoL, nil //没有视频，也不传错误信息
	}

	videoL = pack.VideoList(videos, users, relations)
	return videoL, nil
}
