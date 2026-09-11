package sysMenus

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/liuzhengtao/auth-common-backend/api/v1/dept"
	"github.com/liuzhengtao/auth-common-backend/api/v1/menus"
	"github.com/liuzhengtao/auth-common-backend/internal/consts"
	"github.com/liuzhengtao/auth-common-backend/internal/dao"
	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

type sSysMenusService struct{}

func New() *sSysMenusService {
	return &sSysMenusService{}
}

func init() {
	service.RegisterSysMenusService(New())
}

func (s *sSysMenusService) ListMenus(ctx context.Context, keyword, status string) (menuList []menus.MenuVO, err error) {
	orm := dao.SysMenu.Ctx(ctx)
	if !gstr.Equal(keyword, "") {
		orm = orm.Where("name like ?", "%"+keyword+"%")
	}
	if !gstr.Equal(status, "") {
		orm = orm.Where("visible", status)
	}
	var menuEntity []entity.SysMenu
	err = orm.OrderAsc("sort").Scan(&menuEntity)
	if err != nil {
		return nil, err
	}
	menuIds := gset.NewIntSet()
	parentIds := gset.NewIntSet()
	for _, menu := range menuEntity {
		menuIds.Add(gconv.Int(menu.Id))
		parentIds.Add(gconv.Int(menu.ParentId))
	}
	rootIds := make([]int64, 0)
	for _, id := range parentIds.Slice() {
		if !menuIds.Contains(id) {
			rootIds = append(rootIds, gconv.Int64(id))
		}
	}
	for _, rootId := range rootIds {
		buildMenuTree, err := s.buildMenuVOs(rootId, menuEntity)
		if err != nil {
			return nil, err
		}
		menuList = append(menuList, buildMenuTree...)
	}
	return
}

func (s *sSysMenusService) GetMenuForm(ctx context.Context, id int64) (out *model.MenuFormOutput, err error) {
	var menuForm *entity.SysMenu
	err = dao.SysMenu.Ctx(ctx).Where("id", id).Scan(&menuForm)
	if err != nil {
		return nil, err
	}
	out = &model.MenuFormOutput{
		SysMenu: menuForm,
	}
	if menuForm.Type == consts.MENU.Value() {
		out.Type = consts.MENU.Key()
	} else if menuForm.Type == consts.BUTTON.Value() {
		out.Type = consts.BUTTON.Key()
	} else if menuForm.Type == consts.CATALOG.Value() {
		out.Type = consts.CATALOG.Key()
	} else if menuForm.Type == consts.EXTLINK.Value() {
		out.Type = consts.EXTLINK.Key()
	}
	return
}

func (s *sSysMenusService) SaveMenus(ctx context.Context, menuForm entity.SysMenu) (err error) {

	if menuForm.Type == consts.CATALOG.Value() {
		//如果是外链
		if menuForm.ParentId == 0 && !strings.HasPrefix(menuForm.Path, "/") {
			// 一级目录需以 / 开头
			menuForm.Path = "/" + menuForm.Path
		}
		menuForm.Component = "Layout"
	}
	treePath, err := s.generateMenuTreePath(ctx, menuForm.ParentId)
	if err != nil {
		return err
	}
	menuForm.TreePath = treePath
	_, err = dao.SysMenu.Ctx(ctx).Data(menuForm).Save()
	return
}
func (s *sSysMenusService) generateMenuTreePath(ctx context.Context, parentId int64) (string, error) {
	if consts.ROOT_NODE_ID == parentId {
		return gconv.String(parentId), nil
	} else {
		var parent *entity.SysMenu
		err := dao.SysMenu.Ctx(ctx).Where("id", parentId).Scan(&parent)
		if err != nil {
			return "", err
		}
		if parent != nil {
			return parent.TreePath + "," + gconv.String(parent.Id), nil
		}
		return "", nil
	}
}
func (s *sSysMenusService) ListRolePerms(ctx context.Context, roles []string) (perms []string, err error) {
	return dao.SysMenu.ListRolePerms(ctx, roles)
}

func (s *sSysMenusService) DeleteMenu(ctx context.Context, id int64) error {
	orm := dao.SysMenu.Ctx(ctx).Where("id", id).WhereOr("CONCAT (',',tree_path,',') LIKE CONCAT('%,',?,',%')", id)
	array, err := orm.Fields("id").Array()
	if err != nil {
		return err
	}
	result, err := orm.Delete()
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected > 0 {
		var ids []int64
		for _, v := range array {
			ids = append(ids, gconv.Int64(v))
		}
		return service.SysRoleMenu().RemoveRoleMenusByMenuId(ctx, ids)
	}
	return nil
}

func (s *sSysMenusService) ListRoutes(ctx context.Context) (routeVo []*menus.RouteVO, err error) {
	menuList, err := dao.SysMenu.ListRoutes(ctx)
	if err != nil {
		return nil, err
	}
	return s.buildRoutes(consts.ROOT_NODE_ID, menuList)
}

func (s *sSysMenusService) OptionList(ctx context.Context) (options []dept.Option, err error) {
	var menuOptions []entity.SysMenu
	err = dao.SysMenu.Ctx(ctx).OrderAsc("sort").Scan(&menuOptions)
	if err != nil {
		return nil, err
	}
	return s.buildMenuOptions(ctx, consts.ROOT_NODE_ID, menuOptions)
}

func (s *sSysMenusService) buildMenuOptions(ctx context.Context, parentId int64, menuList []entity.SysMenu) (options []dept.Option, err error) {
	for _, menu := range menuList {
		if menu.ParentId == parentId {
			option := dept.Option{
				Value: menu.Id,
				Label: menu.Name,
			}
			children, err := s.buildMenuOptions(ctx, menu.Id, menuList)
			if err != nil {
				return nil, err
			}
			option.Children = children
			options = append(options, option)
		}
	}
	return
}
func (s *sSysMenusService) toRouteVo(menu *model.RouteBO) (vo *menus.RouteVO) {
	routeName := gstr.UcWords(gstr.CaseCamel(menu.Path))
	return &menus.RouteVO{
		Path:      menu.Path,
		Redirect:  menu.Redirect,
		Component: menu.Component,
		Name:      routeName,
		Meta: menus.Meta{
			Title:      menu.Name,
			Icon:       menu.Icon,
			Roles:      menu.Roles,
			Hidden:     menu.Visible == 0,
			KeepAlive:  menu.Type == consts.MENU.Value() && menu.KeepAlive == 1,
			AlwaysShow: menu.AlwaysShow == 1,
		},
	}
}

func (s *sSysMenusService) buildRoutes(parentId int64, menuList []*model.RouteBO) (routeVo []*menus.RouteVO, err error) {
	routeVo = make([]*menus.RouteVO, 0)
	for _, menu := range menuList {
		if menu.ParentId == parentId {
			vo := s.toRouteVo(menu)
			children, err := s.buildRoutes(menu.ID, menuList)
			if err != nil {
				return nil, err
			}
			vo.Children = children
			routeVo = append(routeVo, vo)
		}
	}
	return
}

func (s *sSysMenusService) buildMenuVOs(id int64, menuEntity []entity.SysMenu) (menusVo []menus.MenuVO, err error) {
	for _, menu := range menuEntity {
		if menu.ParentId == id {
			menuVo := menus.MenuVO{
				SysMenu: menu,
			}
			children, err := s.buildMenuVOs(menu.Id, menuEntity)
			if err != nil {
				return nil, err
			}
			menuVo.Children = children
			menusVo = append(menusVo, menuVo)
		}
	}
	return
}
