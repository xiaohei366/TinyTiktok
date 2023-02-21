package pack

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/FavoriteServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
	"reflect"
)

// 点赞列表的报文格式
type FavoriteListResponse struct {
	StatusCode int32                   `json:"status_code"`
	StatusMsg  string                  `json:"status_msg"`
	VideoList  []*FavoriteServer.Video `json:"user_list"`
}

// 点赞/取消赞的报文格式
type FavoriteActionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// 点赞/取消赞的响应报文
func SendFavoriteActionResponse(c *app.RequestContext, resp interface{}) {
	switch value := resp.(type) {
	case error:
		Err := errno.ConvertErr(value)
		c.JSON(consts.StatusOK, FavoriteActionResponse{
			StatusCode: Err.ErrCode,
			StatusMsg:  Err.ErrMsg,
		})
	case *FavoriteServer.DouyinFavoriteActionResponse:
		c.JSON(consts.StatusOK, FavoriteActionResponse{
			StatusCode: value.BaseResp.StatusCode,
			StatusMsg:  value.BaseResp.StatusMsg,
		})
	default:
		klog.Error("响应报文传入未知类型%v", reflect.TypeOf(resp))
	}
}

// 发送点赞列表的响应报文
func SendFavoriteListResponse(c *app.RequestContext, resp interface{}) {
	switch value := resp.(type) {
	case error:
		Err := errno.ConvertErr(value)
		c.JSON(consts.StatusOK, FavoriteListResponse{
			StatusCode: Err.ErrCode,
			StatusMsg:  Err.ErrMsg,
			VideoList:  nil,
		})
	case *FavoriteServer.DouyinFavoriteListResponse:
		c.JSON(consts.StatusOK, FavoriteListResponse{
			StatusCode: value.BaseResp.StatusCode,
			StatusMsg:  value.BaseResp.StatusMsg,
			VideoList:  value.VideoList,
		})
	default:
		klog.Error("响应报文传入未知类型%v", reflect.TypeOf(resp))
	}
}
