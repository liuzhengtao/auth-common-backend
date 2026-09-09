// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/liuzhengtao/auth-common-backend/api/v1/user"
	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
)

type (
	ISysUserService interface {
		ListPageUsers(ctx context.Context, queryParams *model.UserPageQueryInput) (rows []*user.UserPageVO, total int, err error)
		CreateUser(ctx context.Context, in *model.CreateUserInput) (lastInsertId int64, err error)
		GetUserFormData(ctx context.Context, userId int64) (form *user.Form, err error)
		GetUser(ctx context.Context, userId int64) (user *entity.SysUser, err error)
		UpdateUser(ctx context.Context, in *model.UpdateUserInput) (err error)
		DeleteUsers(ctx context.Context, in *model.DeleteUserInput) (err error)
		UpdatePasswd(ctx context.Context, in *model.UpdatePasswdInput) (err error)
		UpdateStatus(ctx context.Context, in *model.UpdateStatusInput) (err error)
		GetCurrentUserInfo(ctx context.Context) (out *model.GetCurrentInfoOutput, err error)
		GetUserToToken(ctx context.Context, username, password string) (user *entity.SysUser, err error)
		GetUserAuthInfo(ctx context.Context, username string) (*model.UserAuthInfo, error)
	}
)

var (
	localSysUserService ISysUserService
)

func SysUserService() ISysUserService {
	if localSysUserService == nil {
		panic("implement not found for interface ISysUserService, forgot register?")
	}
	return localSysUserService
}

func RegisterSysUserService(i ISysUserService) {
	localSysUserService = i
}
