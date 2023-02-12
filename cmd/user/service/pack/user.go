package pack

import (
	"github.com/xiaohei366/TinyTiktok/cmd/user/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/user/kitex_gen/UserServer"
)

/* 拼接用户信息--服务于获取用户信息的接口*/

func UserInfo(u *db.User) *UserServer.User {
	if u == nil {
		return nil
	}
	return &UserServer.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   u.FollowCount, 
		FollowerCount: u.FollowerCount, 
		IsFollow:      false,
	}
}
