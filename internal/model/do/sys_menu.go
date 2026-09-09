// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure of table sys_menu for DAO operations like Where/Data.
type SysMenu struct {
	g.Meta     `orm:"table:sys_menu, do:true"`
	Id         interface{} //
	ParentId   interface{} // 父菜单ID
	TreePath   interface{} // 父节点ID路径
	Name       interface{} // 菜单名称
	Type       interface{} // 菜单类型(1:菜单 2:目录 3:外链 4:按钮)
	Path       interface{} // 路由路径(浏览器地址栏路径)
	Component  interface{} // 组件路径(vue页面完整路径，省略.vue后缀)
	Perm       interface{} // 权限标识
	Visible    interface{} // 显示状态(1-显示;0-隐藏)
	Sort       interface{} // 排序
	Icon       interface{} // 菜单图标
	Redirect   interface{} // 跳转路径
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	AlwaysShow interface{} // 【目录】只有一个子路由是否始终显示(1:是 0:否)
	KeepAlive  interface{} // 【菜单】是否开启页面缓存(1:是 0:否)
}
