// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure for table sys_menu.
type SysMenu struct {
	Id         int64       `json:"id"         ` //
	ParentId   int64       `json:"parentId"   ` // 父菜单ID
	TreePath   string      `json:"treePath"   ` // 父节点ID路径
	Name       string      `json:"name"       ` // 菜单名称
	Type       int         `json:"type"       ` // 菜单类型(1:菜单 2:目录 3:外链 4:按钮)
	Path       string      `json:"path"       ` // 路由路径(浏览器地址栏路径)
	Component  string      `json:"component"  ` // 组件路径(vue页面完整路径，省略.vue后缀)
	Perm       string      `json:"perm"       ` // 权限标识
	Visible    int         `json:"visible"    ` // 显示状态(1-显示;0-隐藏)
	Sort       int         `json:"sort"       ` // 排序
	Icon       string      `json:"icon"       ` // 菜单图标
	Redirect   string      `json:"redirect"   ` // 跳转路径
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
	AlwaysShow int         `json:"alwaysShow" ` // 【目录】只有一个子路由是否始终显示(1:是 0:否)
	KeepAlive  int         `json:"keepAlive"  ` // 【菜单】是否开启页面缓存(1:是 0:否)
}
