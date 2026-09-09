package test

import (
	"testing"

	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

func TestGetMaximumDataScope(t *testing.T) {
	scope, err := service.SysRoleService().GetMaximumDataScope(ctx, []string{"ADMIN"})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(scope)
}
