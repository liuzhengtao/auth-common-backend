// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/liuzhengtao/auth-common-backend/api/v1/dept"
	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
)

type (
	ISysRoleService interface {
		ListPage(ctx context.Context, keywords string, pageNum int, pageSize int) (out *model.ListOutput, err error)
		GetMaximumDataScope(ctx context.Context, roles []string) (dataScope int, err error)
		RoleOptionsList(ctx context.Context) (list []dept.Option, err error)
		CreateRole(ctx context.Context, input *model.CreateRoleInput) (lastInsertId int64, err error)
		UpdateRole(ctx context.Context, input *model.UpdateRoleInput) (err error)
		DeleteRole(ctx context.Context, in *model.DeleteRoleInput) (err error)
		ListMenuIdsByRoleId(ctx context.Context, roleId int64) (menuIds []int64, err error)
		GetRoleInfo(ctx context.Context, id int64) (sysRole *entity.SysRole, err error)
		AssignMenusToRole(ctx context.Context, in *model.AssignMenusToRoleInput) (err error)
	}
)

var (
	localSysRoleService ISysRoleService
)

func SysRoleService() ISysRoleService {
	if localSysRoleService == nil {
		panic("implement not found for interface ISysRoleService, forgot register?")
	}
	return localSysRoleService
}

func RegisterSysRoleService(i ISysRoleService) {
	localSysRoleService = i
}
