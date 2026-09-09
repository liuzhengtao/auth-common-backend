// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDept is the golang structure of table sys_dept for DAO operations like Where/Data.
type SysDept struct {
	g.Meta     `orm:"table:sys_dept, do:true"`
	Id         interface{} // 主键
	Name       interface{} // 部门名称
	ParentId   interface{} // 父节点id
	TreePath   interface{} // 父节点id路径
	Sort       interface{} // 显示顺序
	Status     interface{} // 状态(1:正常;0:禁用)
	Deleted    interface{} // 逻辑删除标识(1:已删除;0:未删除)
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	CreateBy   interface{} // 创建人ID
	UpdateBy   interface{} // 修改人ID
}
