package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/video/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/pack"
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

	videoL = pack.VideoList(videos, req.UserId)
	return videoL, nil
}
