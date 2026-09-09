// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/liuzhengtao/auth-common-backend/api/v1/dept"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
)

type (
	ISysDeptService interface {
		DeptOptionsList(ctx context.Context) (list []dept.Option, err error)
		ListDepartments(ctx context.Context, keywords string, status int) (list []dept.DeptVo, err error)
		Add(ctx context.Context, formData dept.DepartmentForm) error
		Update(ctx context.Context, formData dept.DepartmentForm) error
		Delete(ctx context.Context, ids string) error
		FindOne(ctx context.Context, id int) (dept *entity.SysDept, err error)
		GetForm(ctx context.Context, id int) (dept *dept.DepartmentForm, err error)
	}
)

var (
	localSysDeptService ISysDeptService
)

func SysDeptService() ISysDeptService {
	if localSysDeptService == nil {
		panic("implement not found for interface ISysDeptService, forgot register?")
	}
	return localSysDeptService
}

func RegisterSysDeptService(i ISysDeptService) {
	localSysDeptService = i
}
