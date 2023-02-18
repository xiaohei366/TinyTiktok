package pack

import (
	"reflect"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// 关注/粉丝列表的报文格式
type FollowListResponse struct {
	StatusCode int32                  `json:"status_code"`
	StatusMsg  string                 `json:"status_msg"`
	UserList   []*RelationServer.User `json:"user_list"`
}

// 关注/取关动作的报文格式
type FollowActionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// 发送关注/粉丝列表的响应报文
func SendFollowListResponse(c *app.RequestContext, resp interface{}) {
	switch value := resp.(type) {
	case error:
		Err := errno.ConvertErr(value)
		c.JSON(consts.StatusOK, FollowListResponse{
			StatusCode: Err.ErrCode,
			StatusMsg:  Err.ErrMsg,
			UserList:   nil,
		})
	case *RelationServer.DouyinRelationFollowListResponse:
		c.JSON(consts.StatusOK, FollowListResponse{
			StatusCode: value.BaseResp.StatusCode,
			StatusMsg:  value.BaseResp.StatusMsg,
			UserList:   value.UserList,
		})
	case *RelationServer.DouyinRelationFollowerListResponse:
		c.JSON(consts.StatusOK, FollowListResponse{
			StatusCode: value.BaseResp.StatusCode,
			StatusMsg:  value.BaseResp.StatusMsg,
			UserList:   value.UserList,
		})
	default:
		klog.Error("响应报文传入未知类型%v", reflect.TypeOf(resp))
	}
}

// 关注/取关的响应报文
func SendFollowActionResponse(c *app.RequestContext, resp interface{}) {
	switch value := resp.(type) {
	case error:
		Err := errno.ConvertErr(value)
		c.JSON(consts.StatusOK, FollowListResponse{
			StatusCode: Err.ErrCode,
			StatusMsg:  Err.ErrMsg,
		})
	case *RelationServer.DouyinRelationActionResponse:
		c.JSON(consts.StatusOK, FollowListResponse{
			StatusCode: value.BaseResp.StatusCode,
			StatusMsg:  value.BaseResp.StatusMsg,
		})
	default:
		klog.Error("响应报文传入未知类型%v", reflect.TypeOf(resp))
	}
}
