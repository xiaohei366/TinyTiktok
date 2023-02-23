package pack

import (
	"github.com/xiaohei366/TinyTiktok/cmd/user/initialize/db"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
)

/* 拼接用户信息--将db.User转化成user.User*/
func UserInfoConvert(u *db.User) *UserServer.User {
	if u == nil {
		return nil
	}
	return &UserServer.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   u.FollowCount, 
		FollowerCount: u.FollowerCount, 
	}
}
