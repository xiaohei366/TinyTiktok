package pack

import (
	"github.com/xiaohei366/TinyTiktok/cmd/feed/kitex_gen/FeedServer"
	"github.com/xiaohei366/TinyTiktok/cmd/feed/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/user/kitex_gen/UserServer"
)

// Video pack video info
func Video(v *dal.Video) *FeedServer.Video { //这个没问题
	if v == nil {
		return nil
	}
	return &FeedServer.Video{
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

// FeedList pack list of videos
func FeedList(vs []*dal.Video) []*FeedServer.Video {
	feedList := make([]*FeedServer.Video, 0)
	for _, v := range vs {
		if video2 := Video(v); video2 != nil {
			feedList = append(feedList, video2)
		}
	}
	return feedList
}
