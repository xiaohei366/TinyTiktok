package pack

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// FeedResponse
type FeedResponse struct {
	StatusCode int32                `json:"status_code"`
	StatusMsg  string               `json:"status_msg"`
	VideoList  []*VideoServer.Video `json:"video_list"`
}

// PublishActionResponse
type PublishActionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// SendResponse feed pack response
// 发送关注/粉丝列表的响应报文
func SendFeedResponse(c *app.RequestContext, err error, videoList []*VideoServer.Video) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, FeedResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		VideoList:  videoList,
	})
}

// 关注/取关的响应报文
func SendPublishActionResponse(c *app.RequestContext, err error, resp interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, PublishActionResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
	})
}
