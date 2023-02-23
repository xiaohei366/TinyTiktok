package pack

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/video/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/video/rpc"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/CommentServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/FavoriteServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
	"sync"
)

// Video pack feed info
func VideoInfo(v *db.Video, reqId int64) *VideoServer.Video {
	//明天这个地方是要优化的，拿Favcount和ComCount
	//开启协程
	ctx := context.Background()
	author, isFol, comCount, favCount, isFav := getUserInfo(ctx, v, reqId) //这个是请求这条视频的人拿的token信息
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
		FavoriteCount: favCount,
		CommentCount:  comCount,
		IsFavorite:    isFav, //这个地方请求的是请求PublishList的人 Id去查询这些视频

		Title: v.Title,
	}
}

func VideoList(vs []*db.Video, reqId int64) []*VideoServer.Video {
	videos := make([]*VideoServer.Video, 0)
	for _, v := range vs {
		video2 := VideoInfo(v, reqId)
		if video2 != nil {
			videos = append(videos, video2)
		}
	}
	return videos
}

func getUserInfo(ctx context.Context, v *db.Video, reqId int64) (author *UserServer.User, relation bool, comCount int64, favCount int64, isFav bool) {
	var wg sync.WaitGroup
	wg.Add(5)
	var err error
	//插入Author，这里需要将视频的发布者和当前登录的用户传入，才能正确获得isFollow，
	//如果出现错误，不能直接返回失败，将默认值返回，保证稳定
	go func() {
		author, err = rpc.GetUserInfo(ctx, &UserServer.DouyinUserRequest{
			UserId: v.AuthorID,
		})
		if err != nil {
			klog.Infof("getUserInfo rpc.GetUserInfo failed：%v", err)
		} else {
			klog.Info("getUserInfo rpc.GetUserInfo success")
		}
		wg.Done()
	}()
	//拿是否关注的信息
	go func() {
		relation, err = rpc.QueryRelation(ctx, &RelationServer.DouyinQueryRelationRequest{
			UserId:   reqId,
			ToUserId: v.AuthorID,
		})
		if err != nil {
			klog.Infof("getUserInfo rpc.QueryRealtion failed：%v", err)
		} else {
			klog.Info("getUserInfo rpc.QueryRealtion success")
		}
		wg.Done()
	}()
	//拿评论数量
	go func() {
		comList, err := rpc.CommentList(ctx, &CommentServer.DouyinCommentListRequest{
			VideoId: v.ID,
		})
		if err != nil {
			klog.Infof("getUserInfo rpc.CommentList failed：%v", err)
		} else {
			klog.Info("getUserInfo rpc.CommentList success")
		}
		comCount = int64(len(comList.CommentList))
		fmt.Println("comment count:", comCount)
		wg.Done()
	}()
	// 拿Favorite数量：
	go func() {
		favCount, err = rpc.GetVideosFavoriteCount(ctx, &FavoriteServer.DouyinVideoFavoriteRequest{
			VideoId: v.ID,
		})
		if err != nil {
			klog.Infof("getUserInfo rpc.GetVideosFavoriteCount failed：%v", err)
		} else {
			klog.Info("getUserInfo rpc.GetVideosFavoriteCount success")
		}
		fmt.Println("fav count:", favCount)
		wg.Done()
	}()
	go func() {
		fmt.Println("fav query:")
		isFav, err = rpc.QueryUserLikeVideo(ctx, &FavoriteServer.DouyinQueryFavoriteRequest{
			UserId:  reqId,
			VideoId: v.ID,
		})
		if err != nil {
			klog.Infof("getUserInfo rpc.QueryUserLikeVideo failed：%v", err)
		} else {
			klog.Info("getUserInfo rpc.QueryUserLikeVideo success")
		}
		fmt.Println("fav query:", isFav)
		wg.Done()
	}()
	wg.Wait()
	return author, relation, comCount, favCount, isFav
}
