// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure of table sys_role for DAO operations like Where/Data.
type SysRole struct {
	g.Meta     `orm:"table:sys_role, do:true"`
	Id         interface{} //
	Name       interface{} // 角色名称
	Code       interface{} // 角色编码
	Sort       interface{} // 显示顺序
	Status     interface{} // 角色状态(1-正常；0-停用)
	DataScope  interface{} // 数据权限(0-所有数据；1-部门及子部门数据；2-本部门数据；3-本人数据)
	Deleted    interface{} // 逻辑删除标识(0-未删除；1-已删除)
	CreateTime *gtime.Time // 更新时间
	UpdateTime *gtime.Time // 创建时间
}
