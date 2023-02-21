package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/video/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/video/rpc"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"sync"
)

func getUserInfo(ctx context.Context, v *db.Video, reqId int64) (author *UserServer.User, relation bool) {
	var wg sync.WaitGroup
	wg.Add(2)
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
	wg.Wait()
	return author, relation
}
