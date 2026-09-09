// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	Id         int         `json:"id"         ` //
	Username   string      `json:"username"   ` // 用户名
	Nickname   string      `json:"nickname"   ` // 昵称
	Gender     int         `json:"gender"     ` // 性别((1:男;2:女))
	Password   string      `json:"password"   ` // 密码
	DeptId     int         `json:"deptId"     ` // 部门ID
	Avatar     string      `json:"avatar"     ` // 用户头像
	Mobile     string      `json:"mobile"     ` // 联系方式
	Status     int         `json:"status"     ` // 用户状态((1:正常;0:禁用))
	Email      string      `json:"email"      ` // 用户邮箱
	Deleted    int         `json:"deleted"    ` // 逻辑删除标识(0:未删除;1:已删除)
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
}
