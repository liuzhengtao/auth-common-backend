// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
)

type (
	IAccess interface {
		CheckAccess(ctx context.Context, user *entity.SysUser, path string) (isPass bool, err error)
	}
)

var (
	localAccess IAccess
)

func Access() IAccess {
	if localAccess == nil {
		panic("implement not found for interface IAccess, forgot register?")
	}
	return localAccess
}

func RegisterAccess(i IAccess) {
	localAccess = i
}
