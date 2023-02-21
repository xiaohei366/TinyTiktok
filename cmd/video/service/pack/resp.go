package pack

import (
	"errors"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

func BuildFeedResp(err error, videoList []*VideoServer.Video, nextTime int64) *VideoServer.DouyinFeedResponse {
	if err == nil {
		return feedResp(errno.Success, videoList, nextTime)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return feedResp(e, videoList, nextTime)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return feedResp(s, videoList, nextTime)
}

// 封装feed返回
func feedResp(err errno.ErrNo, videoList []*VideoServer.Video, nextTime int64) *VideoServer.DouyinFeedResponse {
	resp := new(VideoServer.DouyinFeedResponse)
	resp.BaseResp = &VideoServer.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
	resp.VideoList = videoList
	resp.NextTime = nextTime
	return resp
}
func BuildPublishActionResp(err error) *VideoServer.DouyinPublishActionResponse {
	if err == nil {
		return publishActionResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return publishActionResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return publishActionResp(s)
}

func publishActionResp(err errno.ErrNo) *VideoServer.DouyinPublishActionResponse {
	resp := new(VideoServer.DouyinPublishActionResponse)
	resp.BaseResp = &VideoServer.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
	return resp
}
func BuildPublishListResp(err error, videoList []*VideoServer.Video) *VideoServer.DouyinPublishListResponse {
	if err == nil {
		return publishListResp(errno.Success, videoList)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return publishListResp(e, videoList)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return publishListResp(s, videoList)
}

func publishListResp(err errno.ErrNo, videoList []*VideoServer.Video) *VideoServer.DouyinPublishListResponse {
	resp := new(VideoServer.DouyinPublishListResponse)
	resp.BaseResp = &VideoServer.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
	resp.VideoList = videoList
	return resp
}
