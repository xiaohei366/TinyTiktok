package mw

import (
	"context"
	"net/http"
	"time"

	"github.com/cloudwego/biz-demo/easy_note/cmd/api/biz/model/demoapi"
	"github.com/cloudwego/biz-demo/easy_note/cmd/api/biz/rpc"
	"github.com/cloudwego/biz-demo/easy_note/kitex_gen/demouser"
	"github.com/xiaohei366/TinyTiktok/pkg/configs"
	"github.com/cloudwego/biz-demo/easy_note/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJWT() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte(configs.SecretKey),
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   configs.IdentityKey,
		// 下面两个函数联动， 一个登录成功生成token，下一个用于在token里面存入用户信息(这里是id)
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &demoapi.User{
				UserID: int64(claims[consts.IdentityKey].(float64)),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					consts.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		//这个是用来登录验证的
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req demoapi.CheckUserRequest
			if err = c.BindAndValidate(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(req.Username) == 0 || len(req.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			return rpc.CheckUser(context.Background(), &demouser.CheckUserRequest{
				Username: req.Username,
				Password: req.Password,
			})
		},
		//用于设置登录的响应函数
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"code":   errno.Success.ErrCode,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		//用于设置 jwt 授权失败后的响应函数，如下函数将参数列表中的错误码和错误信息封装成 json 响应返回。
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    errno.AuthorizationFailedErr.ErrCode,
				"message": message,
			})
		},
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