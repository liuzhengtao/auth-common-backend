// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDict is the golang structure of table sys_dict for DAO operations like Where/Data.
type SysDict struct {
	g.Meta     `orm:"table:sys_dict, do:true"`
	Id         interface{} // 主键
	TypeCode   interface{} // 字典类型编码
	Name       interface{} // 字典项名称
	Value      interface{} // 字典项值
	Sort       interface{} // 排序
	Status     interface{} // 状态(1:正常;0:禁用)
	Defaulted  interface{} // 是否默认(1:是;0:否)
	Remark     interface{} // 备注
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
}
