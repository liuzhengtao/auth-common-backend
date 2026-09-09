// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictType is the golang structure for table sys_dict_type.
type SysDictType struct {
	Id         int64       `json:"id"         ` // 主键
	Name       string      `json:"name"       ` // 类型名称
	Code       string      `json:"code"       ` // 类型编码
	Status     int         `json:"status"     ` // 状态(0:正常;1:禁用)
	Remark     string      `json:"remark"     ` // 备注
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
}
