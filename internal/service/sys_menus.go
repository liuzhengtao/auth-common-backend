// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/liuzhengtao/auth-common-backend/api/v1/dept"
	"github.com/liuzhengtao/auth-common-backend/api/v1/menus"
	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
)

type (
	ISysMenusService interface {
		ListMenus(ctx context.Context, keyword, status string) (menuList []menus.MenuVO, err error)
		GetMenuForm(ctx context.Context, id int64) (out *model.MenuFormOutput, err error)
		SaveMenus(ctx context.Context, menuForm entity.SysMenu) (err error)
		ListRolePerms(ctx context.Context, roles []string) (perms []string, err error)
		DeleteMenu(ctx context.Context, id int64) error
		ListRoutes(ctx context.Context) (routeVo []*menus.RouteVO, err error)
		OptionList(ctx context.Context) (options []dept.Option, err error)
	}
)

var (
	localSysMenusService ISysMenusService
)

func SysMenusService() ISysMenusService {
	if localSysMenusService == nil {
		panic("implement not found for interface ISysMenusService, forgot register?")
	}
	return localSysMenusService
}

func RegisterSysMenusService(i ISysMenusService) {
	localSysMenusService = i
}
