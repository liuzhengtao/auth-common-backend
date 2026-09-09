package controller

import (
	"context"

	"github.com/liuzhengtao/auth-common-backend/api/common"
	"github.com/liuzhengtao/auth-common-backend/api/v1/dept"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

type cSysDeptController struct {
	dept service.ISysDeptService
}

func NewSysDeptController() *cSysDeptController {
	return &cSysDeptController{
		dept: service.SysDeptService(),
	}
}

func (c *cSysDeptController) ListDepartments(ctx context.Context, req *dept.ListDepartmentsReq) (resp dept.ListDepartmentsRes, err error) {
	resp, err = c.dept.ListDepartments(ctx, req.Keyword, req.Status)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cSysDeptController) CreateDepartment(ctx context.Context, req *dept.AddReq) (resp *common.Res, err error) {
	err = c.dept.Add(ctx, req.DepartmentForm)
	return
}

func (c *cSysDeptController) UpdateDepartment(ctx context.Context, req *dept.UpdateReq) (resp *common.Res, err error) {
	req.DepartmentForm.Id = req.DeptId
	err = c.dept.Update(ctx, req.DepartmentForm)
	return
}
func (c *cSysDeptController) DeleteDepartment(ctx context.Context, req *dept.DeleteReq) (resp *common.Res, err error) {
	err = c.dept.Delete(ctx, req.Ids)
	return
}

func (c *cSysDeptController) GetForm(ctx context.Context, req *dept.GetFormReq) (resp *dept.GetFormRes, err error) {
	form, err := c.dept.GetForm(ctx, req.DeptId)
	if err != nil {
		return nil, err
	}
	resp = (*dept.GetFormRes)(form)
	return
}

func (c *cSysDeptController) ListDeptOptions(ctx context.Context, req *dept.DeptOptionsReq) (resp dept.DeptOptionsRes, err error) {
	return c.dept.DeptOptionsList(ctx)
}
