// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ISysUserRoleService interface {
		// SaveUserRoles 保存用户角色
		SaveUserRoles(ctx context.Context, userId int, rolesIds []int) error
		// HasAssignedUsers 判断角色是否存在绑定的用户
		HasAssignedUsers(ctx context.Context, roleId int64) (bool, error)
		DeleteByUserIds(ctx context.Context, userIds []string) error
		DeleteByRoleIds(ctx context.Context, roleIds []string) error
	}
)

var (
	localSysUserRoleService ISysUserRoleService
)

func SysUserRoleService() ISysUserRoleService {
	if localSysUserRoleService == nil {
		panic("implement not found for interface ISysUserRoleService, forgot register?")
	}
	return localSysUserRoleService
}

func RegisterSysUserRoleService(i ISysUserRoleService) {
	localSysUserRoleService = i
}
