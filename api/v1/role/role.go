package role

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/liuzhengtao/auth-common-backend/api/common"
	"github.com/liuzhengtao/auth-common-backend/api/v1/dept"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
)

type PageQueryReq struct {
	g.Meta  `path:"/page" method:"get" tags:"角色接口" dc:"角色分页查询"`
	Keyword string `p:"keywords" dc:"关键字(用户名/昵称/手机号)"`
	*common.BasePageQuery
}
type PageQueryRes struct {
	Total uint32   `json:"total" dc:"总条数"`
	List  gdb.List `json:"list" dc:"列表"`
}
type ListRolesOptionsReq struct {
	g.Meta `path:"/options" method:"get" tags:"角色接口" dc:"角色下拉列表"`
}
type ListRolesOptionsRes []dept.Option
type Data struct {
	Id        int64  `p:"id" dc:"角色ID"`
	Name      string `p:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Code      string `p:"code" v:"required#角色编码不能为空" dc:"角色编码"`
	Sort      int32  `p:"sort" dc:"排序"`
	Status    int32  `p:"status" dc:"状态" dc:"角色状态(1-正常；0-停用)"`
	DataScope int32  `p:"dataScope" dc:"数据权限"`
}

type SaveRoleReq struct {
	g.Meta `path:"/" method:"post" tags:"角色接口" dc:"新增角色"`
	Data
}

type UpdateRoleReq struct {
	g.Meta `path:"/{id}" method:"put" tags:"角色接口" dc:"新增角色"`
	Data
}

type FormReq struct {
	g.Meta `path:"/{roleId}/form" method:"get" tags:"角色接口" dc:"新增角色"`
	RoleId int64 `p:"roleId" in:"path" v:"required#角色ID不能为空" dc:"角色ID"`
}

type FormRes struct {
	*entity.SysRole
}

type DeleteRoleReq struct {
	g.Meta `path:"/{ids}" method:"delete" tags:"角色接口" dc:"删除角色"`
	Ids    string `p:"ids" in:"path" v:"required#角色ID不能为空" dc:"删除角色，多个以英文逗号(,)分割"`
}

type MenuIdsReq struct {
	g.Meta `path:"/{roleId}/menuIds" method:"get" tags:"角色接口" dc:"获取角色的菜单ID集合"`
	RoleId int64 `p:"roleId" in:"path" v:"required#角色ID不能为空" dc:"角色ID"`
}

type MenuIdsRes []int64

type MenusReq struct {
	g.Meta `path:"/{roleId}/menus" method:"put" tags:"角色接口" dc:"分配菜单权限给角色"`
	RoleId int64   `p:"roleId" in:"path" v:"required#角色ID不能为空" dc:"角色ID"`
	Body   []int64 `p:"menuIds" dc:"菜单ID"`
}

type MenusRes struct {
	common.Result
}

type UpdateStatusReq struct {
	g.Meta `path:"/{roleId}/status" method:"put" tags:"角色接口" dc:"更新角色状态"`
	RoleId int64 `p:"roleId" in:"path" v:"required#角色ID不能为空" dc:"角色ID"`
	Status uint  `p:"status"  v:"required#角色状态不能为空"  dc:"角色状态"`
}

type RolePermsBO struct {
	RoleCode string   `json:"roleCode"`
	Perms    []string `json:"perms"`
}
