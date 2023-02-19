package pack

import (
	"context"
	"github.com/xiaohei366/TinyTiktok/cmd/video/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/video/rpc"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// Video pack feed info
func Video(ctx context.Context, v *db.Video, uid int64) (*VideoServer.Video, error) {
	if v == nil {
		return nil, nil
	}

	//随后通过RPC 由这些ID获得 用户信息
	authorInfo, err := rpc.GetUserInfo(ctx, &UserServer.DouyinUserRequest{
		UserId: v.AuthorID,
	})
	if err != nil {
		return nil, errno.UserRPCErr
	}

	favorite_count := int64(v.FavCount)
	comment_count := int64(v.ComCount)

	//todo is Fav
	return &VideoServer.Video{
		Id:            v.BaseModel.ID,
		Author:        authorInfo,
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
