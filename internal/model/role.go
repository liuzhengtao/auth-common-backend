package model

import (
	"github.com/liuzhengtao/auth-common-backend/api/v1/role"
)

type ListOutput struct {
	role.PageQueryRes
}
type CreateRoleInput struct {
	*role.SaveRoleReq
}

type UpdateRoleInput struct {
	*role.UpdateRoleReq
}

type DeleteRoleInput struct {
	*role.DeleteRoleReq
}

type AssignMenusToRoleInput struct {
	*role.MenusReq
}
