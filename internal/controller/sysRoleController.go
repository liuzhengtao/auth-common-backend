package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/liuzhengtao/auth-common-backend/api/common"
	"github.com/liuzhengtao/auth-common-backend/api/v1/role"
	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

type cRoleController struct {
	role service.ISysRoleService
}

func NewRoleController() *cRoleController {
	return &cRoleController{
		role: service.SysRoleService(),
	}
}

func (c *cRoleController) ListPage(ctx context.Context, req *role.PageQueryReq) (res *role.PageQueryRes, err error) {
	page, err := c.role.ListPage(ctx, req.Keyword, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	res = &page.PageQueryRes
	return
}
func (c *cRoleController) ListRolesOptions(ctx context.Context, req *role.ListRolesOptionsReq) (res role.ListRolesOptionsRes, err error) {
	return c.role.RoleOptionsList(ctx)
}

func (c *cRoleController) SaveRole(ctx context.Context, req *role.SaveRoleReq) (res *common.Res, err error) {
	_, err = c.role.CreateRole(ctx, &model.CreateRoleInput{SaveRoleReq: req})
	if err != nil {
		return nil, err
	}
	return nil, err
}

func (c *cRoleController) DeleteRole(ctx context.Context, req *role.DeleteRoleReq) (res *common.Res, err error) {
	err = c.role.DeleteRole(ctx, &model.DeleteRoleInput{DeleteRoleReq: req})
	return nil, err
}

func (c *cRoleController) UpdateRole(ctx context.Context, req *role.UpdateRoleReq) (res *common.Res, err error) {
	err = c.role.UpdateRole(ctx, &model.UpdateRoleInput{UpdateRoleReq: req})
	return nil, err
}

func (c *cRoleController) GetRoleForm(ctx context.Context, req *role.FormReq) (res *role.FormRes, err error) {
	sysRole, err := c.role.GetRoleInfo(ctx, req.RoleId)
	if err != nil {
		return nil, err
	}
	res = &role.FormRes{
		SysRole: sysRole,
	}
	return
}

func (c *cRoleController) GetMenuIds(ctx context.Context, req *role.MenuIdsReq) (res role.MenuIdsRes, err error) {
	res, err = c.role.ListMenuIdsByRoleId(ctx, req.RoleId)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cRoleController) AssignMenusToRole(ctx context.Context, req *role.MenusReq) (res *common.Res, err error) {
	r := g.RequestFromCtx(ctx)
	err = gconv.Scan(r.GetBodyString(), &req.Body)
	if err != nil {
		return nil, err
	}
	err = c.role.AssignMenusToRole(ctx, &model.AssignMenusToRoleInput{MenusReq: req})
	return nil, err
}

func (c *cRoleController) UpdateStatus(ctx context.Context, req *role.UpdateStatusReq) (res *common.Res, err error) {
	return nil, err
}
