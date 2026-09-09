package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/liuzhengtao/auth-common-backend/api/common"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

// Auth JWT 认证中间件，宿主业务路由可直接 group.Middleware(middleware.Auth)。
func Auth(r *ghttp.Request) {
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
		g.Log().Error(r.GetCtx(), "[SYSTEM ERROR]", err)
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
