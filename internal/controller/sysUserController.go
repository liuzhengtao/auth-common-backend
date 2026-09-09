package controller

import (
	"context"

	v1 "github.com/liuzhengtao/auth-common-backend/api/v1/user"
	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

type cSysUserController struct {
	userService service.ISysUserService
}

func NewSysUserController() *cSysUserController {
	return &cSysUserController{
		userService: service.SysUserService(),
	}
}

func (cUser *cSysUserController) ListPageUsers(ctx context.Context, queryParams *v1.PageQueryReq) (list *v1.ListPagedUsersRes, err error) {
	rows, total, err := cUser.userService.ListPageUsers(ctx, &model.UserPageQueryInput{
		PageQueryReq: queryParams,
	})
	if err != nil {
		return nil, err
	}
	list = &v1.ListPagedUsersRes{
		List:  rows,
		Total: total,
	}
	return
}

func (cUser *cSysUserController) SaveUser(ctx context.Context, userForm *v1.SaveUserReq) (res *v1.NoResultRes, err error) {
	_, err = cUser.userService.CreateUser(ctx, &model.CreateUserInput{SaveUserReq: userForm})
	if err != nil {
		return nil, err
	}
	return
}

func (cUser *cSysUserController) GetUserForm(ctx context.Context, userForm *v1.GetUserFormReq) (res *v1.GetUserFormRes, err error) {
	form, err := cUser.userService.GetUserFormData(ctx, userForm.UserId)
	if err != nil {
		return nil, err
	}
	res = &v1.GetUserFormRes{
		Form: form,
	}
	return
}

func (cUser *cSysUserController) UpdateUser(ctx context.Context, userForm *v1.UpdateUserReq) (res *v1.NoResultRes, err error) {
	err = cUser.userService.UpdateUser(ctx, &model.UpdateUserInput{
		UpdateUserReq: userForm,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (cUser *cSysUserController) DeleteUsers(ctx context.Context, userForm *v1.DeleteUserReq) (res *v1.NoResultRes, err error) {
	err = cUser.userService.DeleteUsers(ctx, &model.DeleteUserInput{
		DeleteUserReq: userForm,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (cUser *cSysUserController) UpdatePasswd(ctx context.Context, userForm *v1.UpdatePasswdReq) (res *v1.NoResultRes, err error) {
	err = cUser.userService.UpdatePasswd(ctx, &model.UpdatePasswdInput{
		UpdatePasswdReq: userForm,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (cUser *cSysUserController) UpdateStatus(ctx context.Context, userForm *v1.UpdateUserStatusReq) (res *v1.NoResultRes, err error) {
	err = cUser.userService.UpdateStatus(ctx, &model.UpdateStatusInput{
		UpdateUserStatusReq: userForm,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (cUser *cSysUserController) GetCurrentUserInfo(ctx context.Context, userForm *v1.GetCurrentUserInfoReq) (res *v1.GetCurrentUserInfoRes, err error) {
	info, err := cUser.userService.GetCurrentUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.GetCurrentUserInfoRes{
		UserInfoVO: info.UserInfoVO,
	}
	return
}
