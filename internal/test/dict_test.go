package test

import (
	"testing"

	"github.com/liuzhengtao/auth-common-backend/internal/applog"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

func TestListDictOptionsService(t *testing.T) {
	list, err := service.SysDictService().ListDictOptions(ctx, "gender")
	if err != nil {
		t.Error(err)
		return
	}
	applog.Get().Info(ctx, list)
}
