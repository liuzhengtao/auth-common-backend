// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDict is the golang structure for table sys_dict.
type SysDict struct {
	Id         int64       `json:"id"         ` // 主键
	TypeCode   string      `json:"typeCode"   ` // 字典类型编码
	Name       string      `json:"name"       ` // 字典项名称
	Value      string      `json:"value"      ` // 字典项值
	Sort       int         `json:"sort"       ` // 排序
	Status     int         `json:"status"     ` // 状态(1:正常;0:禁用)
	Defaulted  int         `json:"defaulted"  ` // 是否默认(1:是;0:否)
	Remark     string      `json:"remark"     ` // 备注
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
}
