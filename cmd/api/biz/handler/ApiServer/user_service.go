package ApiServer

import (
	"context"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/handler/pack"
	mw "github.com/xiaohei366/TinyTiktok/cmd/api/biz/middleware"
	ApiServer "github.com/xiaohei366/TinyTiktok/cmd/api/biz/model/ApiServer"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/rpc"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

// Register .
// @router /douyin/user/register/ [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.DouyinUserRegisterRequest
	var user_id int64
	var token string
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	//调用PRC方法，完成注册操作
	user_id, err = rpc.Register(context.Background(), &UserServer.DouyinUserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		pack.SendRegisterResponse(c, errno.ConvertErr(err), -1, "")
		return
	}
	// 使用JWT来产生Token--注意是使用id，因为这个将作为存储信息在token中
	token, _, err = mw.JwtMiddleware.TokenGenerator(user_id)
	if err != nil {
		pack.SendRegisterResponse(c, errno.ConvertErr(err), -1, "")
		return
	}
	//成功响应
	pack.SendRegisterResponse(c, errno.Success, user_id, token)
}

// Login .
// @router /douyin/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	//先执行Authenticator--期间失败则执行unauthorized&HTTPStatusMessageFunc
	//随后创建token
	//若PayloadFunc不为空，则此时执行PayloadFunc
	//最后执行LoginResponse返回信息
	mw.JwtMiddleware.LoginHandler(ctx, c)
}

// GetUserInfo .
// @router /douyin/user/ [GET]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.DouyinUserRequest
	var wg sync.WaitGroup
	var u *UserServer.User
	var isFollow bool
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	//开协程调用两个RPC方法
	//如果出现错误，不能直接返回失败，将默认值返回，保证稳定
	wg.Add(2)
	go func() {
		u, err = rpc.GetUserInfo(context.Background(), &UserServer.DouyinUserRequest{
			UserId: req.UserId,
		})
		if err != nil {
			klog.Errorf("调用用户模块失败:%v", err)
		}
		wg.Done()
	}()
	go func() {
		v, _ := c.Get(shared.IdentityKey) // 取出token的id
		isFollow, err = rpc.QueryRelation(context.Background(), &RelationServer.DouyinQueryRelationRequest{
			UserId:   v.(*ApiServer.User).Id,
			ToUserId: req.UserId,
		})
		if err != nil {
			klog.Errorf("调用关系模块失败:%v", err)
			return
		}
		wg.Done()
	}()
	wg.Wait()
	//成功响应
	pack.SendUesrInfoResponse(c, errno.Success, u, isFollow)
}
