package pack

import (
	"github.com/xiaohei366/TinyTiktok/cmd/user/service/dal"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
)

/* 拼接用户信息--服务于获取用户信息的接口*/

func UserInfo(u *dal.User) *UserServer.User {
	if u == nil {
		return nil
	}
	return &UserServer.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   10, //test
		FollowerCount: 10, //test
		IsFollow:      false,
	}
}
