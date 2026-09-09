package controller

import (
	"context"

	"github.com/liuzhengtao/auth-common-backend/api/common"
	"github.com/liuzhengtao/auth-common-backend/api/v1/menus"
	"github.com/liuzhengtao/auth-common-backend/internal/consts"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

type cMenusController struct {
	menu service.ISysMenusService
}

func NewMenusController() *cMenusController {
	return &cMenusController{
		menu: service.SysMenusService(),
	}
}

func (c *cMenusController) ListMenus(ctx context.Context, req *menus.ListMenusReq) (resp menus.ListMenusRes, err error) {
	resp, err = c.menu.ListMenus(ctx, req.Keyword, req.Status)
	if err != nil {
		return nil, err
	}
	return
}
func (c *cMenusController) GetMenuForm(ctx context.Context, req *menus.GetMenuFormReq) (resp menus.GetMenuFormRes, err error) {
	resp, err = c.menu.GetMenuForm(ctx, req.Id)
	return
}
func (c *cMenusController) UpdateMenu(ctx context.Context, req *menus.UpdateMenusReq) (resp *common.Res, err error) {
	if req.Type == consts.MENU.Key() {
		req.SysMenu.Type = consts.MENU.Value()
	} else if req.Type == consts.BUTTON.Key() {
		req.SysMenu.Type = consts.BUTTON.Value()
	} else if req.Type == consts.CATALOG.Key() {
		req.SysMenu.Type = consts.CATALOG.Value()
	} else if req.Type == consts.EXTLINK.Key() {
		req.SysMenu.Type = consts.EXTLINK.Value()
	}
	err = c.menu.SaveMenus(ctx, req.SysMenu)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cMenusController) SaveMenu(ctx context.Context, req *menus.SaveMenusReq) (resp *common.Res, err error) {
	if req.Type == consts.MENU.Key() {
		req.SysMenu.Type = consts.MENU.Value()
	} else if req.Type == consts.BUTTON.Key() {
		req.SysMenu.Type = consts.BUTTON.Value()
	} else if req.Type == consts.CATALOG.Key() {
		req.SysMenu.Type = consts.CATALOG.Value()
	} else if req.Type == consts.EXTLINK.Key() {
		req.SysMenu.Type = consts.EXTLINK.Value()
	}
	err = c.menu.SaveMenus(ctx, req.SysMenu)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cMenusController) DeleteMenu(ctx context.Context, req *menus.DeleteMenuReq) (resp *common.Res, err error) {
	err = c.menu.DeleteMenu(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return
}
func (c *cMenusController) ListRoutes(ctx context.Context, param *menus.ListRoutesReq) (resp *menus.ListRoutesRes, err error) {
	routeList, err := c.menu.ListRoutes(ctx)
	if err != nil {
		return nil, err
	}
	return &menus.ListRoutesRes{
		RouteList: routeList,
	}, nil
}

func (c *cMenusController) ListOptions(ctx context.Context, param *menus.ListOptionsReq) (resp menus.ListOptionsRes, err error) {
	return c.menu.OptionList(ctx)
}
