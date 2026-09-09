// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDept is the golang structure for table sys_dept.
type SysDept struct {
	Id         int64       `json:"id"         ` // 主键
	Name       string      `json:"name"       ` // 部门名称
	ParentId   int64       `json:"parentId"   ` // 父节点id
	TreePath   string      `json:"treePath"   ` // 父节点id路径
	Sort       int         `json:"sort"       ` // 显示顺序
	Status     int         `json:"status"     ` // 状态(1:正常;0:禁用)
	Deleted    int         `json:"deleted"    ` // 逻辑删除标识(1:已删除;0:未删除)
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
	CreateBy   int64       `json:"createBy"   ` // 创建人ID
	UpdateBy   int64       `json:"updateBy"   ` // 修改人ID
}
