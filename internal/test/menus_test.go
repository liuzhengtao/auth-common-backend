package test

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/liuzhengtao/auth-common-backend/internal/dao"
)

func TestListRolePerms(t *testing.T) {
	perms, err := dao.SysMenu.ListRolePerms(ctx, []string{"ADMIN"})
	if err != nil {
		t.Error(err)
		return
	}
	g.Log().Info(ctx, perms)
}
