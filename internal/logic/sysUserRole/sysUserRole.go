package sysUserRole

import (
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/liuzhengtao/auth-common-backend/internal/dao"
	"github.com/liuzhengtao/auth-common-backend/internal/model/do"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

type sSysUserRoleService struct{}

func New() *sSysUserRoleService {
	return &sSysUserRoleService{}
}

func init() {
	service.RegisterSysUserRoleService(New())
}

// SaveUserRoles 保存用户角色
func (s *sSysUserRoleService) SaveUserRoles(ctx context.Context, userId int, rolesIds []int) error {
	if len(rolesIds) == 0 {
		return gerror.New("roleIds参数不能为空")
	}
	// 用户原角色ID集合
	result, err := dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.Columns().UserId, userId).Fields(dao.SysUserRole.Columns().RoleId).Array()
	if err != nil {
		return err
	}
	userRoleIds := garray.NewIntArray()
	for _, value := range result {
		userRoleIds.Append(value.Int())
	}
	var saveRoleIds []int
	if userRoleIds.IsEmpty() {
		saveRoleIds = rolesIds
	} else {
		for _, value := range rolesIds {
			if !userRoleIds.Contains(value) {
				saveRoleIds = append(saveRoleIds, value)
			}
		}
	}
	var userRoles []do.SysUserRole
	for _, id := range saveRoleIds {
		userRoles = append(userRoles, do.SysUserRole{
			UserId: userId,
			RoleId: id,
		})
	}
	if len(userRoles) > 0 {
		_, err = dao.SysUserRole.Ctx(ctx).Data(userRoles).Insert()
		if err != nil {
			return err
		}
	}
	if !userRoleIds.IsEmpty() {
		var removeRoleIds []int
		for _, id := range userRoleIds.Slice() {
			if !garray.NewIntArrayFrom(rolesIds).Contains(id) {
				removeRoleIds = append(removeRoleIds, id)
			}
		}
		if len(removeRoleIds) > 0 {
			_, err = dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.Columns().UserId, userId).WhereIn(dao.SysUserRole.Columns().RoleId, removeRoleIds).Delete()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// HasAssignedUsers 判断角色是否存在绑定的用户
func (s *sSysUserRoleService) HasAssignedUsers(ctx context.Context, roleId int64) (bool, error) {
	count, err := dao.SysUserRole.Ctx(ctx).As("t1").InnerJoin("sys_role t2", "t1.role_id = t2.id AND t2.deleted = 0").InnerJoin("sys_user t3", "t1.user_id = t3.id AND t3.deleted = 0").Where("t1.role_id", roleId).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *sSysUserRoleService) DeleteByUserIds(ctx context.Context, userIds []string) error {
	_, err := dao.SysUserRole.Ctx(ctx).WhereIn(dao.SysUserRole.Columns().UserId, userIds).Delete()
	return err
}

func (s *sSysUserRoleService) DeleteByRoleIds(ctx context.Context, roleIds []string) error {
	result, err := dao.SysUserRole.Ctx(ctx).WhereIn(dao.SysUserRole.Columns().RoleId, roleIds).Delete()
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected > 0 {
		return service.SysRoleMenu().RemoveByRoleIds(ctx, roleIds)
	}
	return err
}
