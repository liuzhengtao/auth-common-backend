package test

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

func TestListDictOptionsService(t *testing.T) {
	list, err := service.SysDictService().ListDictOptions(ctx, "gender")
	if err != nil {
		t.Error(err)
		return
	}
	g.Log().Info(ctx, list)
}
