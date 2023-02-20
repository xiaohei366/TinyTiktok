package pack

import (
	"context"

	"github.com/xiaohei366/TinyTiktok/cmd/video/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/video/rpc"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// Video pack feed info
func VideoList(ctx context.Context, v *db.Video, uid int64) (*VideoServer.Video, error) {
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
	//随后通过RPC 由这些ID获得 两个用户的关系
	isFollow, err := rpc.QueryRelation(context.Background(), &RelationServer.DouyinQueryRelationRequest{
		UserId:   uid,
		ToUserId: v.AuthorID,
	})
	//fmt.Print("********************", uid, v.AuthorID)
	if err != nil {
		return nil, errno.RelationRPCErr
	}

	favorite_count := int64(v.FavCount)
	comment_count := int64(v.ComCount)
	//将两个封装成一份完整的用户信息
	//todo is Fav
	return &VideoServer.Video{
		Id: v.BaseModel.ID,
		Author: &VideoServer.User{
			Id:            authorInfo.Id,
			Name:          authorInfo.Name,
			FollowCount:   authorInfo.FollowCount,
			FollowerCount: authorInfo.FollowerCount,
			IsFollow:      isFollow,
		},
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: favorite_count,
		CommentCount:  comment_count,
		IsFavorite:    false,
		Title:         v.Title,
	}, nil
}

// Videos pack list of video info
func VideoLists(ctx context.Context, vs []*db.Video, uid int64) ([]*VideoServer.Video, error) {
	videos := make([]*VideoServer.Video, 0)
	for _, v := range vs {
		video2, err := VideoList(ctx, v, uid)
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
