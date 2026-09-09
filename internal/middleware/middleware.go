package middleware

import (
	publicmw "github.com/liuzhengtao/auth-common-backend/middleware"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Auth 兼容旧内部引用，转调公共 middleware 包。
func Auth(r *ghttp.Request) {
	publicmw.Auth(r)
}

// ErrorHandler 兼容旧内部引用，转调公共 middleware 包。
func ErrorHandler(r *ghttp.Request) {
	publicmw.ErrorHandler(r)
}
