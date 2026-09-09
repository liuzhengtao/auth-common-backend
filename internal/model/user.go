package model

import (
	"github.com/liuzhengtao/auth-common-backend/api/common"
	"github.com/liuzhengtao/auth-common-backend/api/v1/user"
)

type UserPageQueryInput struct {
	*user.PageQueryReq
}

type UserPage struct {
	*common.Page
}

type CreateUserInput struct {
	*user.SaveUserReq
}

type UpdateUserInput struct {
	*user.UpdateUserReq
}

type DeleteUserInput struct {
	*user.DeleteUserReq
}
type UpdatePasswdInput struct {
	*user.UpdatePasswdReq
}

type UpdateStatusInput struct {
	*user.UpdateUserStatusReq
}

type GetCurrentInfoOutput struct {
	*user.UserInfoVO
}

type UserAuthInfo struct {
	UserId    int      `json:"userId"      ` // 用户ID
	Username  string   `json:"username"   `  // 用户名
	Nickname  string   `json:"nickname"   `  // 昵称
	Password  string   `json:"password"   `  // 密码
	DeptId    int      `json:"deptId"     `  // 部门ID
	Status    int      `json:"status"     `  // 用户状态((1:正常;0:禁用))
	Roles     []string `json:"roles"      `  // 角色
	Perms     []string `json:"perms"`        // 权限
	DataScope int      `json:"dataScope"`    // 数据权限
}
