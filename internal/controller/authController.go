package controller

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/liuzhengtao/auth-common-backend/api/v1/auth"
	v1 "github.com/liuzhengtao/auth-common-backend/api/v1/user"
	"github.com/liuzhengtao/auth-common-backend/internal/consts"
	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

type cAuthController struct {
	auth service.IAuthService
}

func NewAuthController() *cAuthController {
	return &cAuthController{
		auth: service.AuthService(),
	}
}
func (c *cAuthController) Captcha(ctx context.Context, param *auth.CaptchaReq) (resp *auth.CaptchaRes, err error) {
	output, err := service.CaptchaService().GetCaptcha(ctx)
	if err != nil {
		return nil, err
	}
	return output.CaptchaRes, nil
}

func (c *cAuthController) Login(ctx context.Context, param *auth.LoginReq) (resp *auth.LoginRes, err error) {
	r := g.RequestFromCtx(ctx)
	if !service.CaptchaService().VerifyAndClear(r, param.CaptchaId, param.CaptchaCode) {
		return nil, gerror.New(consts.VERIFY_CODE_ERROR)
	}
	output, err := c.auth.Login(ctx, &model.LoginInput{Username: param.Username, Password: param.Password})
	if err != nil {
		return nil, err
	}
	resp = output.LoginRes
	return
}

func (c *cAuthController) Logout(ctx context.Context, req *auth.LogoutReq) (res *v1.NoResultRes, err error) {
	err = c.auth.Logout(ctx)
	if err != nil {
		return nil, err
	}
	return
}
