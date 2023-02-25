package ApiServer

import (
	"bytes"
	"context"
	"io"
	"strconv"

	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/handler/pack"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/rpc"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/model/ApiServer"
)

// Feed .
// @router /douyin/feed/ [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var laststTime, useID int64
	// 获取最近的时间并判断处理
	lastst_time := c.Query("latest_time")
	if len(lastst_time) != 0 {
		if latesttime, err := strconv.Atoi(lastst_time); err != nil {
			pack.SendFeedResponse(c, errno.ConvertErr(err), nil)
			return
		} else {
			laststTime = int64(latesttime)
		}
	}
	//获取token中传来的user id
	user, _ := c.Get(shared.IdentityKey)
	if user == nil {
		useID = 0
	} else {
		useID = user.(*ApiServer.User).Id
	}
	//rpc调用获取视频feed
	videos, err := rpc.FeedVideos(context.Background(), &VideoServer.DouyinFeedRequest{
		LatestTime: laststTime,
		UserId:     useID,
	})
	if err != nil {
		pack.SendFeedResponse(c, errno.ConvertErr(err), nil)
	}
	pack.SendFeedResponse(c, errno.Success, videos)
}

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {

	var err error
	var req ApiServer.DouyinPublishActionRequest

	req.Title = c.PostForm("title")
	req.Token = c.PostForm("token")
	//接收视频文件并处理。
	fileHeader, err := c.Request.FormFile("data")
	if err != nil {
		pack.SendPublishActionResponse(c, errno.ConvertErr(err), nil)
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		pack.SendPublishActionResponse(c, errno.ConvertErr(err), nil)
		return
	}
	defer file.Close()
	//copy为[]byte格式
	buf := bytes.NewBuffer(nil)
	_, _ = io.Copy(buf, file)
	if err != nil {
		pack.SendPublishActionResponse(c, errno.ConvertErr(err), nil)
		return
	}
	//获取
	userId, _ := c.Get(shared.IdentityKey)
	request := &VideoServer.DouyinPublishActionRequest{
		Token:  req.Token,
		Title:  req.Title,
		Data:   buf.Bytes(),
		UserId: userId.(*ApiServer.User).Id,
	}
	resp, err := rpc.PublishVideos(ctx, request)
	if err != nil {
		pack.SendPublishActionResponse(c, errno.ConvertErr(err), nil)
	}
	pack.SendPublishActionResponse(c, errno.Success, resp)
}

// PublishList .
// @router /douyin/publish/list/ [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.DouyinPublishListRequest
	var useID int64
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	//拿userid
	user, _ := c.Get(shared.IdentityKey)
	if user == nil {
		useID = 0
	} else {
		useID = user.(*ApiServer.User).Id
	}
	request := &VideoServer.DouyinPublishListRequest{
		UserId:   useID,
		ToUserId: req.UserId,
	}
	//调用rpc获取该用户已经发布的视频。
	resp, err := rpc.PublishList(ctx, request)
	if err != nil {
		pack.SendFeedResponse(c, errno.ConvertErr(err), nil)
	}

	pack.SendFeedResponse(c, errno.Success, resp) //这个只有测试才知道对不对
}
