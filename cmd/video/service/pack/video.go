package pack

import (
	"github.com/xiaohei366/TinyTiktok/cmd/video/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/cmd/video/kitex_gen/VideoServer"
	"github.com/xiaohei366/TinyTiktok/cmd/video/service/dal"
)

// Video pack video info
func Video(v *dal.Video) *VideoServer.Video { //这个没问题
	if v == nil {
		return nil
	}
	return &VideoServer.Video{
		Author: &UserServer.User{
			Id: v.AuthorID,
		},
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavCount,
		CommentCount:  v.ComCount, //这个地方是写错字了。
		IsFavorite:    v.IsFavorite,
		Title:         v.Title,
	}
}

// VideoList pack list of videos
func VideoList(vs []*dal.Video) []*VideoServer.Video {
	feedList := make([]*VideoServer.Video, 0)
	for _, v := range vs {
		if video2 := Video(v); video2 != nil {
			feedList = append(feedList, video2)
		}
	}
	return feedList
}
