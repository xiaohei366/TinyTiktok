package pack

import (
	"github.com/xiaohei366/TinyTiktok/kitex_gen/FavoriteServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
)

/* 拼接用户信息--服务于获取用户信息的接口*/
//这个什么用？
//func FavoriteInfo(f *FavoriteServer.User) *FavoriteServer.User {
//	if f == nil {
//		return nil
//	}
//	return &FavoriteServer.User{
//		Id:              f.Id,
//		Name:            f.Name,
//		FollowCount:     10, //test?
//		FollowerCount:   10, //test?
//		IsFollow:        false,
//		TotalBeFavorite: 10, //test?
//		FavoriteCount:   10, //test?
//	}
//}

// 将VideoServer的格式转换为FavoriteServer的格式，视频信息在videoServer已经打包好了
func ConvertVideos(videos []*VideoServer.Video) []*FavoriteServer.Video {
	videosList := make([]*FavoriteServer.Video, 0)
	for _, v := range videos {
		videosList = append(videosList, &FavoriteServer.Video{
			Id: v.Id,
			Author: &FavoriteServer.User{
				Id:            v.Author.Id,
				Name:          v.Author.Name,
				FollowCount:   v.Author.FollowCount,
				FollowerCount: v.Author.FollowerCount,
				IsFollow:      v.Author.IsFollow,
				//TotalBeFavorite: ,todo
				//FavoriteCount: ,  todo
			},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
		})
	}
	return videosList
}
