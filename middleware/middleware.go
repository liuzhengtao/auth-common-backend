package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/liuzhengtao/auth-common-backend/api/common"
	"github.com/liuzhengtao/auth-common-backend/internal/applog"
	"github.com/liuzhengtao/auth-common-backend/internal/config"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
	"github.com/liuzhengtao/auth-common-backend/internal/tokenblacklist"
)

// Auth JWT 认证中间件，宿主业务路由可直接 group.Middleware(middleware.Auth)。
// 分布式模式下会先校验 Token 黑名单。
func Auth(r *ghttp.Request) {
	if config.Get().Distributed {
		token := tokenblacklist.ExtractToken(r)
		if token != "" && tokenblacklist.IsBlocked(r.Context(), token) {
			r.Response.WriteJsonExit(g.Map{
				"code": 401,
				"msg":  "token has been revoked",
			})
			return
		}
	}
	service.Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}

// ErrorHandler 统一成功/失败 JSON 响应包装。
func ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	var (
		err = r.GetError()
		res = r.GetHandlerResponse()
	)
	if err != nil {
		applog.Get().Error(r.GetCtx(), "[SYSTEM ERROR]", err)
		r.Response.ClearBuffer()
		r.Response.WriteJsonExit(
			common.ResultFailed(err.Error()),
		)
	} else if res != nil {
		r.Response.WriteJsonExit(
			common.ResultSuccess(res),
		)
	}
}
