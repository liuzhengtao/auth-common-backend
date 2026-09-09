// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/liuzhengtao/auth-common-backend/internal/model"
)

type (
	ICaptchaService interface {
		NewAndStore(ctx context.Context, captchaStoreKey string) error
		GetCaptcha(ctx context.Context) (out *model.CaptchaResult, err error)
		Store(ctx context.Context, captchaStoreKey, captchaStoreVal string) error
		VerifyAndClear(r *ghttp.Request, captchaStoreKey string, value string) bool
	}
)

var (
	localCaptchaService ICaptchaService
)

func CaptchaService() ICaptchaService {
	if localCaptchaService == nil {
		panic("implement not found for interface ICaptchaService, forgot register?")
	}
	return localCaptchaService
}

func RegisterCaptchaService(i ICaptchaService) {
	localCaptchaService = i
}
