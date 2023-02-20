package pack

import (
	"github.com/xiaohei366/TinyTiktok/cmd/video/initialize/db"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
)

// Video pack feed info
func VideoInfo(v *db.Video, author *UserServer.User, isFol bool) *VideoServer.Video {
	return &VideoServer.Video{
		Id: v.BaseModel.ID,
		Author: &VideoServer.User{
			Id:            author.Id,
			Name:          author.Name,
			FollowCount:   author.FollowCount,
			FollowerCount: author.FollowerCount,
			IsFollow:      isFol,
		},
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavCount,
		CommentCount:  v.ComCount,
		IsFavorite:    false, //todo
		Title:         v.Title,
	}
}

func VideoList(vs []*db.Video, authors []*UserServer.User, relation []bool) []*VideoServer.Video {
	videos := make([]*VideoServer.Video, 0)
	for i, v := range vs {
		video2 := VideoInfo(v, authors[i], relation[i])
		//todo is fav
		if video2 != nil {
			videos = append(videos, video2)
		}
	}
	return videos
}
