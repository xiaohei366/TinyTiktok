package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	jwt "github.com/hertz-contrib/jwt"
)

func MiddlewareFuncOptional() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		middlewareImplOptional(ctx, c)
	}
}

func middlewareImplOptional(ctx context.Context, c *app.RequestContext) {
	claims, err := JwtMiddleware.GetClaimsFromJWT(ctx, c)
	//如果token为空，就跳过，由于只有feed流用，因此这里只针对query情况。（后续可加更多）
	if err == jwt.ErrEmptyQueryToken {
		c.Next(ctx)
		return
	}
	if err != nil {
		unauthorized(ctx, c, http.StatusUnauthorized, JwtMiddleware.HTTPStatusMessageFunc(err, ctx, c))
		return
	}
	switch v := claims["exp"].(type) {
	case nil:
		unauthorized(ctx, c, http.StatusBadRequest, JwtMiddleware.HTTPStatusMessageFunc(jwt.ErrMissingExpField, ctx, c))
		return
	case float64:
		if int64(v) < JwtMiddleware.TimeFunc().Unix() {
			unauthorized(ctx, c, http.StatusUnauthorized, JwtMiddleware.HTTPStatusMessageFunc(jwt.ErrExpiredToken, ctx, c))
			return
		}
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			unauthorized(ctx, c, http.StatusBadRequest, JwtMiddleware.HTTPStatusMessageFunc(jwt.ErrWrongFormatOfExp, ctx, c))
			return
		}
		if n < JwtMiddleware.TimeFunc().Unix() {
			unauthorized(ctx, c, http.StatusUnauthorized, JwtMiddleware.HTTPStatusMessageFunc(jwt.ErrExpiredToken, ctx, c))
			return
		}
	default:
		JwtMiddleware.Unauthorized(ctx, c, http.StatusBadRequest, JwtMiddleware.HTTPStatusMessageFunc(jwt.ErrWrongFormatOfExp, ctx, c))
	}

	c.Set("JWT_PAYLOAD", claims)
	identity := JwtMiddleware.IdentityHandler(ctx, c)

	if identity != nil {
		c.Set(JwtMiddleware.IdentityKey, identity)
	}

	if !JwtMiddleware.Authorizator(identity, ctx, c) {
		unauthorized(ctx, c, http.StatusForbidden, JwtMiddleware.HTTPStatusMessageFunc(jwt.ErrForbidden, ctx, c))
		return
	}

	c.Next(ctx)
}

func unauthorized(ctx context.Context, c *app.RequestContext, code int, message string) {
	c.Header("WWW-Authenticate", "JWT realm="+JwtMiddleware.Realm)
	if !JwtMiddleware.DisabledAbort {
		c.Abort()
	}

	JwtMiddleware.Unauthorized(ctx, c, code, message)
}
