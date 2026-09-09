// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"time"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
)

type (
	IAuth interface {
		MiddlewareFunc() ghttp.HandlerFunc
		GetIdentity(ctx context.Context) (info *model.UserAuthInfo)
		PayLoadFunc(data any) jwt.MapClaims
		IdentityHandler(ctx context.Context) any
		// Unauthorized 登录验证失败，处理不进行授权的逻辑
		Unauthorized(ctx context.Context, code int, message string)
		GetPayload(ctx context.Context) string
		GeneratorToken(ctx context.Context, data model.UserAuthInfo) (token string, expire time.Time, err error)
		RefreshHandler(ctx context.Context) (token string, expire time.Time)
		LogoutHandler(ctx context.Context)
		GetUser(ctx context.Context) (*entity.SysUser, error)
		GetRoles(ctx context.Context) (roles []string, err error)
	}
	IAuthService interface {
		Login(ctx context.Context, in *model.LoginInput) (out *model.LoginOutput, err error)
		Logout(ctx context.Context) (err error)
	}
)

var (
	localAuth        IAuth
	localAuthService IAuthService
)

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}

func AuthService() IAuthService {
	if localAuthService == nil {
		panic("implement not found for interface IAuthService, forgot register?")
	}
	return localAuthService
}

func RegisterAuthService(i IAuthService) {
	localAuthService = i
}
