package ApiServer

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/kitex_gen/RelationServer"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/model/ApiServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// 注册的响应报文格式
type RegisterResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

// 用户信息的响应报文格式
type UserResponse struct {
	StatusCode int32           `json:"status_code"`
	StatusMsg  string          `json:"status_msg"`
	User       *ApiServer.User `json:"user"`
}

// 关注/粉丝列表的报文格式
type FollowListResponse struct {
	StatusCode int32             `json:"status_code"`
	StatusMsg  string            `json:"status_msg"`
	UserList   []*RelationServer.User `json:"user_list"`
}

// 关注/取关动作的报文格式
type FollowActionResponse struct {
	StatusCode int32             `json:"status_code"`
	StatusMsg  string            `json:"status_msg"`
}

// 发送注册的响应报文
func SendRegisterResponse(c *app.RequestContext, err error, id int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, RegisterResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		UserId:     id,
		Token:      token,
	})
}

// 发送用户信息的响应报文
func SendUesrInfoResponse(c *app.RequestContext, err error, u *UserServer.User, isFollow bool) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, UserResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		User: &ApiServer.User{
			Id:            u.Id,
			Name:          u.Name,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      isFollow,
		},
	})
}

// 发送关注/粉丝列表的响应报文
func SendFollowListResponse(c *app.RequestContext, err error, u []*RelationServer.User) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, FollowListResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		UserList: u,
	})
}

// 发送关注/粉丝列表的响应报文
func SendFollowActionResponse(c *app.RequestContext, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, FollowActionResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
	})
}