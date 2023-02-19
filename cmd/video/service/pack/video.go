package pack

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/user/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/user/service/pack"
	"github.com/xiaohei366/TinyTiktok/cmd/video/initialize/db"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
	"gorm.io/gorm"
)

// Video pack feed info
func Video(ctx context.Context, v *db.Video, uid int64) (*VideoServer.Video, error) {
	if v == nil {
		return nil, nil
	}
	//打包的时候不应该这样去拿userInfo
	user, err := dal.GetUserById(ctx, int64(v.AuthorID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	author := pack.UserInfoConvert(&user)
	klog.Info("pack video user info")
	favorite_count := int64(v.FavCount)
	comment_count := int64(v.ComCount)

	//todo is Fav
	return &VideoServer.Video{
		Id:            v.BaseModel.ID,
		Author:        author,
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: favorite_count,
		CommentCount:  comment_count,
		IsFavorite:    false,
		Title:         v.Title,
	}, nil
}

// Videos pack list of video info
func Videos(ctx context.Context, vs []*db.Video, uid int64) ([]*VideoServer.Video, error) {
	videos := make([]*VideoServer.Video, 0)
	for _, v := range vs {
		video2, err := Video(ctx, v, uid)
		if err != nil {
			return nil, err
		}
		//todo is fav
		if video2 != nil {
			videos = append(videos, video2)
		}
	}
	return videos, nil
}
