package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/video/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/pack"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type VideoListByVideoIdService struct {
	ctx context.Context
}

func NewVideoListByVideoIdService(ctx context.Context) *VideoListByVideoIdService {
	return &VideoListByVideoIdService{ctx: ctx}
}

// 通过视频ID获取视频列表
func (s *VideoListByVideoIdService) GetVideoListByVideoId(req *VideoServer.DouyinVideoListByVideoId) (videoL []*VideoServer.Video, err error) {
	var videos []*db.Video
	for _, vId := range req.VideoId {
		video, err := dal.GetVideoByVideosId(s.ctx, vId)
		if err != nil {
			klog.Info("video Id not found ")
			return nil, errno.GetVideoListByVideoIdErr
		} else {
			videos = append(videos, video)
		}
	}

	videoL = pack.VideoList(videos, req.UserId)
	return videoL, nil
}
