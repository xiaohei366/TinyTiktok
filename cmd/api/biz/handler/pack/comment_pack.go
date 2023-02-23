package pack

import (
	"reflect"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/CommentServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// 评论列表的报文格式
type CommentListResponse struct {
	StatusCode  int32                    `json:"status_code"`
	StatusMsg   string                   `json:"status_msg"`
	CommentList []*CommentServer.Comment `json:"comment_list"`
}

// 评论动作的报文格式
type CommentActionResponse struct {
	StatusCode int32                  `json:"status_code"`
	StatusMsg  string                 `json:"status_msg"`
	Comment    *CommentServer.Comment `json:"comment"`
}

// 发送评论列表的响应报文
func SendCommentListResponse(c *app.RequestContext, resp interface{}) {
	switch value := resp.(type) {
	case error:
		Err := errno.ConvertErr(value)
		c.JSON(consts.StatusOK, CommentListResponse{
			StatusCode:  Err.ErrCode,
			StatusMsg:   Err.ErrMsg,
			CommentList: nil,
		})
	case *CommentServer.DouyinCommentListResponse:
		c.JSON(consts.StatusOK, CommentListResponse{
			StatusCode:  value.BaseResp.StatusCode,
			StatusMsg:   value.BaseResp.StatusMsg,
			CommentList: value.CommentList,
		})
	default:
		klog.Error("响应报文传入未知类型%v", reflect.TypeOf(resp))
	}
}

// 创建/删除评论的响应报文
func SendCommentActionResponse(c *app.RequestContext, resp interface{}) {
	switch value := resp.(type) {
	case error:
		Err := errno.ConvertErr(value)
		c.JSON(consts.StatusOK, CommentActionResponse{
			StatusCode: Err.ErrCode,
			StatusMsg:  Err.ErrMsg,
			Comment:    nil,
		})
	case *CommentServer.DouyinCommentActionResponse:
		c.JSON(consts.StatusOK, CommentActionResponse{
			StatusCode: value.BaseResp.StatusCode,
			StatusMsg:  value.BaseResp.StatusMsg,
			Comment:    value.Comment,
		})
	default:
		klog.Error("响应报文传入未知类型%v", reflect.TypeOf(resp))
	}
}
