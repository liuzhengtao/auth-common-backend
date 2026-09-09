package sysRoleMenu

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/liuzhengtao/auth-common-backend/internal/dao"
	"github.com/liuzhengtao/auth-common-backend/internal/model/do"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

type sSysRoleMenu struct {
	logger *glog.Logger
}

func New() *sSysRoleMenu {
	return &sSysRoleMenu{
		logger: g.Log().Line(true),
	}
}

func init() {
	service.RegisterSysRoleMenu(New())
}

func (s *sSysRoleMenu) RemoveRoleMenus(ctx context.Context, roleId int64) (err error) {
	_, err = dao.SysRoleMenu.Ctx(ctx).Delete(dao.SysRoleMenu.Columns().RoleId, roleId)
	return err
}

func (s *sSysRoleMenu) RemoveRoleMenusByMenuId(ctx context.Context, menuIds []int64) (err error) {
	_, err = dao.SysRoleMenu.Ctx(ctx).WhereIn(dao.SysRoleMenu.Columns().MenuId, menuIds).Delete()
	return err
}
func (s *sSysRoleMenu) RemoveByRoleIds(ctx context.Context, roleIds []string) (err error) {
	_, err = dao.SysRoleMenu.Ctx(ctx).WhereIn(dao.SysRoleMenu.Columns().RoleId, roleIds).Delete()
	return err
}

func (s *sSysRoleMenu) SaveBatch(ctx context.Context, roleMenus []do.SysRoleMenu) (err error) {
	_, err = dao.SysRoleMenu.Ctx(ctx).Data(roleMenus).Insert()
	return err
}
