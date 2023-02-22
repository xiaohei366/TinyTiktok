package pack

import (
	"github.com/xiaohei366/TinyTiktok/kitex_gen/FavoriteServer"
)

/* 拼接用户信息--服务于获取用户信息的接口*/

func FavoriteInfo(f *FavoriteServer.User) *FavoriteServer.User {
	if f == nil {
		return nil
	}
	return &FavoriteServer.User{
		Id:             f.Id,
		Name:           f.Name,
		FollowCount:    10, //test
		FollowerCount:  10, //test
		IsFollow:       false,
		TotalFavorited: 10, //test
		FavoriteCount:  10, //test
	}
}
