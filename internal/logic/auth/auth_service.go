package auth

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/liuzhengtao/auth-common-backend/api/v1/auth"
	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

type sAuthService struct {
}

func NewAuthService() *sAuthService {
	return &sAuthService{}
}

func init() {
	service.RegisterAuthService(NewAuthService())
}

func (s *sAuthService) Login(ctx context.Context, in *model.LoginInput) (out *model.LoginOutput, err error) {
	loginUser, err := service.SysUserService().GetUserToToken(ctx, in.Username, in.Password)
	if err != nil {
		return nil, err
	}
	info, err := service.SysUserService().GetUserAuthInfo(ctx, loginUser.Username)
	if err != nil {
		return nil, err
	}
	token, t, err := service.Auth().GeneratorToken(ctx, *info)
	if err != nil {
		return nil, err
	}
	out = &model.LoginOutput{}
	out.LoginRes = &auth.LoginRes{
		AccessToken: token,
		TokenType:   "Bearer",
		Expires:     t.UnixMilli()}
	return
}

func (s *sAuthService) Logout(ctx context.Context) (err error) {
	return g.Try(ctx, func(ctx context.Context) {
		service.Auth().LogoutHandler(ctx)
	})
}
