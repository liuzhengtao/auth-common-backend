package dao

import "github.com/liuzhengtao/auth-common-backend/internal/dao/internal"

// Init 按数据库配置组重建全部 DAO 全局实例。空串回落 default。
func Init(group string) {
	if group == "" {
		group = "default"
	}
	SysUser = sysUserDao{internal.NewSysUserDao(group)}
	SysDept = sysDeptDao{internal.NewSysDeptDao(group)}
	SysRole = sysRoleDao{internal.NewSysRoleDao(group)}
	SysMenu = sysMenuDao{internal.NewSysMenuDao(group)}
	SysDict = sysDictDao{internal.NewSysDictDao(group)}
	SysDictType = sysDictTypeDao{internal.NewSysDictTypeDao(group)}
	SysUserRole = sysUserRoleDao{internal.NewSysUserRoleDao(group)}
	SysRoleMenu = sysRoleMenuDao{internal.NewSysRoleMenuDao(group)}
}
