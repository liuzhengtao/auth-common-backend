package test

import (
	"testing"

	"github.com/liuzhengtao/auth-common-backend/internal/applog"
	"github.com/liuzhengtao/auth-common-backend/internal/dao"
)

func TestListRolePerms(t *testing.T) {
	perms, err := dao.SysMenu.ListRolePerms(ctx, []string{"ADMIN"})
	if err != nil {
		t.Error(err)
		return
	}
	applog.Get().Info(ctx, perms)
}
