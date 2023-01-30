package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/kitex_gen/UserServer"
	api "github.com/xiaohei366/TinyTiktok/cmd/api/biz/model/ApiServer"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/rpc"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJWT() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte(shared.SecretKey),
		TokenLookup:   "form: token, param: token, header: Authorization, query: token", //这里主要用到form和query
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Timeout:       time.Hour,          //指定了 token 有效期为一个小时
		MaxRefresh:    time.Hour,          //用于设置最大 token 刷新时间
		IdentityKey:   shared.IdentityKey, // 用于设置检索身份的键
		////用于设置登录验证的函数----专职服务login
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var id int64
			var req api.DouyinUserLoginRequest
			if err = c.BindAndValidate(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(req.Username) == 0 || len(req.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			//调用rpc方法
			id, err = rpc.Login(context.Background(), &UserServer.DouyinUserLoginRequest{
				Username: req.Username,
				Password: req.Password,
			})
			c.Set("id", id) //将ID存到上下文中，然后返回报文时再取
			return id, err
		},
		//用于设置获取身份信息的函数，此处提取 token 的负载，并配合 IdentityKey 将用户id存入上下文信息。
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &api.User{
				Id: int64(claims[shared.IdentityKey].(float64)),
			}
		},
		//它的入参就是 Authenticator 的返回值，此时负责解析 user，并将用户id注入 token 的 payload 部分
		//这个函数将配合IdentityHandler
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					shared.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		//用于设置登录的响应函数----专职服务login
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			v, _ := c.Get("id")
			c.JSON(code, utils.H{
				"status_code": errno.Success.ErrCode,
				"status_msg":  "success",
				"user_id":     v,
				"token":       token,
			})
		},

		/*下面是登录失败的响应函数*/
		//用于设置 jwt 验证流程失败的响应函数，当前 demo 返回了错误码和错误信息。
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    errno.AuthorizationFailedErr.ErrCode,
				"message": message,
			})
		},
		//用于设置 jwt 校验流程发生错误时响应所包含的错误信息，你可以自行包装这些内容
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch t := e.(type) {
			case errno.ErrNo:
				return t.ErrMsg
			default:
				return t.Error()
			}
		},
	})
}
