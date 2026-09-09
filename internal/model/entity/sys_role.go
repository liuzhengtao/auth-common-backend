// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	Id         int64       `json:"id"         ` //
	Name       string      `json:"name"       ` // 角色名称
	Code       string      `json:"code"       ` // 角色编码
	Sort       int         `json:"sort"       ` // 显示顺序
	Status     int         `json:"status"     ` // 角色状态(1-正常；0-停用)
	DataScope  int         `json:"dataScope"  ` // 数据权限(0-所有数据；1-部门及子部门数据；2-本部门数据；3-本人数据)
	Deleted    int         `json:"deleted"    ` // 逻辑删除标识(0-未删除；1-已删除)
	CreateTime *gtime.Time `json:"createTime" ` // 更新时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 创建时间
}
