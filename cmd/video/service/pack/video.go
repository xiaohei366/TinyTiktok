package pack

import (
	"context"
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
	//根据视频信息从其他服务rpc调用获取信息
	ctx := context.Background()
	author, isFol, comCount, favCount, isFav := getUserInfo(ctx, v, reqId)
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
		IsFavorite:    isFav,
		Title:         v.Title,
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
	//开启协程去进行RPC调用获取视频信息其他部分。
	go func() {
		defer func() {
			if err := recover(); err != nil { //防止协程崩溃，保持健壮性
				klog.Fatalf("Work failed with %s in %v", err, v.AuthorID)
			}
		}()
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
		defer func() {
			if err := recover(); err != nil { //防止协程崩溃，保持健壮性
				klog.Fatalf("Work failed with %s in %v", err, v.AuthorID)
			}
		}()
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
		defer func() {
			if err := recover(); err != nil { //防止协程崩溃，保持健壮性
				klog.Fatalf("Work failed with %s in %v", err, v.AuthorID)
			}
		}()
		comList, err := rpc.CommentList(ctx, &CommentServer.DouyinCommentListRequest{
			VideoId: v.ID,
		})
		if err != nil {
			klog.Infof("getUserInfo rpc.CommentList failed：%v", err)
		} else {
			klog.Info("getUserInfo rpc.CommentList success")
		}
		comCount = int64(len(comList.CommentList))
		wg.Done()
	}()
	// 拿Favorite数量：
	go func() {
		defer func() {
			if err := recover(); err != nil { //防止协程崩溃，保持健壮性
				klog.Fatalf("Work failed with %s in %v", err, v.AuthorID)
			}
		}()
		favCount, err = rpc.GetVideosFavoriteCount(ctx, &FavoriteServer.DouyinVideoBeFavoriteRequest{
			VideoId: v.ID,
		})
		if err != nil {
			klog.Infof("getUserInfo rpc.GetVideosFavoriteCount failed：%v", err)
		} else {
			klog.Info("getUserInfo rpc.GetVideosFavoriteCount success")
		}
		wg.Done()
	}()
	// 拿是否点赞
	go func() {
		defer func() {
			if err := recover(); err != nil { //防止协程崩溃，保持健壮性
				klog.Fatalf("Work failed with %s in %v", err, v.AuthorID)
			}
		}()
		isFav, err = rpc.QueryUserLikeVideo(ctx, &FavoriteServer.DouyinQueryFavoriteRequest{
			UserId:  reqId,
			VideoId: v.ID,
		})
		if err != nil {
			klog.Infof("getUserInfo rpc.QueryUserLikeVideo failed：%v", err)
		} else {
			klog.Info("getUserInfo rpc.QueryUserLikeVideo success")
		}
		wg.Done()
	}()
	wg.Wait()
	return author, relation, comCount, favCount, isFav
}
