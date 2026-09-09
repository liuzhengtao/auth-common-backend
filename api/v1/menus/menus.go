package menus

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/liuzhengtao/auth-common-backend/api/v1/dept"
	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
)

type ListRoutesReq struct {
	g.Meta `path:"/routes" method:"get" tags:"菜单" dc:"获取菜单路由"`
}

type ListMenusReq struct {
	g.Meta  `path:"/" method:"get" tags:"菜单" dc:"获取菜单列表"`
	Keyword string `p:"keywords" dc:"菜单名称"`
	Status  string `p:"status" dc:"状态(1->显示；0->隐藏)"`
}

type SaveMenusReq struct {
	g.Meta `path:"/" method:"post" tags:"菜单" dc:"新增菜单"`
	entity.SysMenu
	Type string `p:"type" v:"required#菜单类型不能为空"`
}

type UpdateMenusReq struct {
	g.Meta `path:"/{id}" method:"put" tags:"菜单" dc:"修改菜单"`
	entity.SysMenu
	Type string `p:"type" v:"required#菜单类型不能为空"`
	Id   int64  `p:"id" in:"path" v:"required#菜单ID不能为空"`
}

type DeleteMenuReq struct {
	g.Meta `path:"/{id}" method:"delete" tags:"菜单" dc:"删除菜单"`
	Id     int64 `p:"id" in:"path" v:"required#菜单ID不能为空"`
}

type GetMenuFormReq struct {
	g.Meta `path:"/{id}/form" method:"get" tags:"菜单" dc:"菜单表单数据"`
	Id     int64 `p:"id" in:"path" v:"required#菜单ID不能为空"`
}
type GetMenuFormRes *model.MenuFormOutput
type MenuVO struct {
	entity.SysMenu
	Children []MenuVO `json:"children"`
}

type ListMenusRes []MenuVO
type ListRoutesRes struct {
	RouteList []*RouteVO `json:"routes"`
}

type ListOptionsReq struct {
	g.Meta `path:"/options" method:"get" tags:"菜单" dc:"菜单下拉列表"`
}

type ListOptionsRes []dept.Option
type RouteVO struct {
	Path      string     `json:"path"`
	Component string     `json:"component"`
	Redirect  string     `json:"redirect"`
	Name      string     `json:"name"`
	Meta      Meta       `json:"meta"`
	Children  []*RouteVO `json:"children"`
}

type Meta struct {
	Title      string   `json:"title"`
	Icon       string   `json:"icon"`
	Hidden     bool     `json:"hidden"`
	Roles      []string `json:"roles"`
	KeepAlive  bool     `json:"keepAlive"`
	AlwaysShow bool     `json:"alwaysShow"`
}
