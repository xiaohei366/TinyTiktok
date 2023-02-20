package ApiServer

import (
	"bytes"
	"context"
	"io"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
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
	var laststTime int64
	lastst_time := c.PostForm("latest_time")
	if len(lastst_time) != 0 {
		if latesttime, err := strconv.Atoi(lastst_time); err != nil {
			pack.SendFeedResponse(c, errno.ConvertErr(err), nil)
			return
		} else {
			laststTime = int64(latesttime)
		}
	}
	user, _ := c.Get(shared.IdentityKey)
	videos, err := rpc.FeedVideos(context.Background(), &VideoServer.DouyinFeedRequest{
		LatestTime: laststTime,
		UserId:     user.(*ApiServer.User).Id,
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
	//_ = c.BindAndValidate(&req) //验证参数

	req.Title = c.PostForm("title")
	req.Token = c.PostForm("token")
	//拿取视频文件。
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

	buf := bytes.NewBuffer(nil)
	_, _ = io.Copy(buf, file)
	if err != nil {
		pack.SendPublishActionResponse(c, errno.ConvertErr(err), nil)
		return
	}
	//拿userid
	userId, _ := c.Get(shared.IdentityKey)
	klog.Info("publish action:", req.Title, userId)
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
	//测试可用
	var err error
	var req ApiServer.DouyinPublishListRequest
	userid := c.Query("user_id")
	req.Token = c.Query("token")
	klog.Info("user_id", userid)
	if userid != "" {
		uid, err := strconv.Atoi(userid)
		if err != nil {
			return
		}
		req.UserId = int64(uid)
	}
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	//拿userid

	request := &VideoServer.DouyinPublishListRequest{
		UserId: req.UserId,
		Token:  req.Token,
	}

	resp, err := rpc.PublishList(ctx, request)
	if err != nil {
		pack.SendFeedResponse(c, errno.ConvertErr(err), nil)
	}
	pack.SendFeedResponse(c, errno.Success, resp) //这个只有测试才知道对不对
}
