package pack

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// Feed响应报文格式
type FeedResponse struct {
	StatusCode int32                `json:"status_code"`
	StatusMsg  string               `json:"status_msg"`
	VideoList  []*VideoServer.Video `json:"video_list"`
}

// 发布视频操作响应报文格式
type PublishActionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// 发送Feed视频流响应报文
func SendFeedResponse(c *app.RequestContext, err error, videoList []*VideoServer.Video) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, FeedResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		VideoList:  videoList,
	})
}

// 发送发布视频操作响应报文
func SendPublishActionResponse(c *app.RequestContext, err error, resp interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, PublishActionResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
	})
}
