package test

import (
	"testing"

	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var ctx = gctx.New()

func TestLoginService(t *testing.T) {
	info, err := service.AuthService().Login(ctx, &model.LoginInput{
		Username: "admin",
		Password: "a123456",
	})
	if err != nil {
		t.Error(err)
		return
	}
	g.Log().Info(ctx, info)
}
