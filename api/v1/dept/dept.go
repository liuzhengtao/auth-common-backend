package dept

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
)

type DeptOptionsReq struct {
	g.Meta `path:"/options" method:"get" tags:"部门" dc:"获取部门下拉选项"`
}

type DeptOptionsRes []Option

type Option struct {
	Label    string   `json:"label"`
	Value    any      `json:"value"`
	Children []Option `json:"children"`
}

type ListDepartmentsReq struct {
	g.Meta  `path:"/" method:"get" tags:"部门" dc:"获取部门列表"`
	Keyword string `p:"keywords" summary:"关键字(部门名称)"`
	Status  int    `p:"status" default:"4" summary:"状态(1->正常;0->禁用)"`
}
type DeptVo struct {
	entity.SysDept
	Children []DeptVo `json:"children"`
}
type DepartmentForm struct {
	Id       int    `p:"id" json:"id" summary:"部门ID"`
	Name     string `p:"name" json:"name" v:"required#请输入部门名称" dc:"部门名称"`
	Sort     int    `p:"sort" json:"sort" dc:"排序"`
	Status   int    `p:"status" json:"status" dc:"状态(1->正常;0->禁用)"`
	ParentId int    `p:"parentId" json:"parentId" v:"required#请选择上级部门" dc:"父部门ID"`
}

type AddReq struct {
	g.Meta `path:"/" method:"post" tags:"部门" dc:"新增部门"`
	DepartmentForm
}

type UpdateReq struct {
	g.Meta `path:"/{deptId}" method:"put" tags:"部门" dc:"修改部门"`
	DeptId int `p:"deptId" in:"path" dc:"部门Id"`
	DepartmentForm
}
type DeleteReq struct {
	g.Meta `path:"/{ids}" method:"delete" tags:"部门" dc:"删除部门"`
	Ids    string `p:"ids" in:"path"  dc:"部门ID，多个以英文逗号(,)分割"`
}

type GetFormReq struct {
	g.Meta `path:"/{deptId}/form" method:"get" tags:"部门" dc:"获取部门表单数据"`
	DeptId int `p:"deptId" in:"path"`
}

type GetFormRes DepartmentForm

type ListDepartmentsRes []DeptVo
