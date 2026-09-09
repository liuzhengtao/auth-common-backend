// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/liuzhengtao/auth-common-backend/internal/model/do"
)

type (
	ISysRoleMenu interface {
		RemoveRoleMenus(ctx context.Context, roleId int64) (err error)
		RemoveRoleMenusByMenuId(ctx context.Context, menuIds []int64) (err error)
		RemoveByRoleIds(ctx context.Context, roleIds []string) (err error)
		SaveBatch(ctx context.Context, roleMenus []do.SysRoleMenu) (err error)
	}
)

var (
	localSysRoleMenu ISysRoleMenu
)

func SysRoleMenu() ISysRoleMenu {
	if localSysRoleMenu == nil {
		panic("implement not found for interface ISysRoleMenu, forgot register?")
	}
	return localSysRoleMenu
}

func RegisterSysRoleMenu(i ISysRoleMenu) {
	localSysRoleMenu = i
}
