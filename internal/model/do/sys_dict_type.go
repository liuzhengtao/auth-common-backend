// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictType is the golang structure of table sys_dict_type for DAO operations like Where/Data.
type SysDictType struct {
	g.Meta     `orm:"table:sys_dict_type, do:true"`
	Id         interface{} // 主键
	Name       interface{} // 类型名称
	Code       interface{} // 类型编码
	Status     interface{} // 状态(0:正常;1:禁用)
	Remark     interface{} // 备注
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
}
