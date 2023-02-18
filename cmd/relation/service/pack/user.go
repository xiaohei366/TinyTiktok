package pack

import (
	
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
)

/* 拼接用户信息--将UserServer.User转化成RelationServer.User*/
func UserInfoConvert(u *UserServer.User, isFollow bool) *RelationServer.User {
	if u == nil {
		return nil
	}
	return &RelationServer.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   u.FollowCount, 
		FollowerCount: u.FollowerCount, 
		IsFollow:      isFollow,
	}
}
