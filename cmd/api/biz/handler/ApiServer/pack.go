package ApiServer

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/model/ApiServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// 发送注册的响应报文
func SendRegisterResponse(c *app.RequestContext, err error, id int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, ApiServer.DouyinUserRegisterResponse{
		BaseResp: &ApiServer.BaseResp{
			StatusCode: Err.ErrCode,
			StatusMsg:  Err.ErrMsg,
		},
		UserId: id,
		Token:  token,
	})
}

// 发送用户信息的响应报文
func SendUesrInfoResponse(c *app.RequestContext, err error, u *UserServer.User) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, ApiServer.DouyinUserResponse{
		BaseResp: &ApiServer.BaseResp{
			StatusCode: Err.ErrCode,
			StatusMsg:  Err.ErrMsg,
		},
		User: &ApiServer.User{
			Id:            u.Id,
			Name:          u.Name,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      u.IsFollow,
		},
	})
}
