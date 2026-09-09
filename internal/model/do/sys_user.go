// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure of table sys_user for DAO operations like Where/Data.
type SysUser struct {
	g.Meta     `orm:"table:sys_user, do:true"`
	Id         interface{} //
	Username   interface{} // 用户名
	Nickname   interface{} // 昵称
	Gender     interface{} // 性别((1:男;2:女))
	Password   interface{} // 密码
	DeptId     interface{} // 部门ID
	Avatar     interface{} // 用户头像
	Mobile     interface{} // 联系方式
	Status     interface{} // 用户状态((1:正常;0:禁用))
	Email      interface{} // 用户邮箱
	Deleted    interface{} // 逻辑删除标识(0:未删除;1:已删除)
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
}
