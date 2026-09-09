// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserDao is the data access object for table sys_user.
type SysUserDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysUserColumns // columns contains all the column names of Table for convenient usage.
}

// SysUserColumns defines and stores column names for table sys_user.
type SysUserColumns struct {
	Id         string //
	Username   string // 用户名
	Nickname   string // 昵称
	Gender     string // 性别((1:男;2:女))
	Password   string // 密码
	DeptId     string // 部门ID
	Avatar     string // 用户头像
	Mobile     string // 联系方式
	Status     string // 用户状态((1:正常;0:禁用))
	Email      string // 用户邮箱
	Deleted    string // 逻辑删除标识(0:未删除;1:已删除)
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// sysUserColumns holds the columns for table sys_user.
var sysUserColumns = SysUserColumns{
	Id:         "id",
	Username:   "username",
	Nickname:   "nickname",
	Gender:     "gender",
	Password:   "password",
	DeptId:     "dept_id",
	Avatar:     "avatar",
	Mobile:     "mobile",
	Status:     "status",
	Email:      "email",
	Deleted:    "deleted",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewSysUserDao creates and returns a new DAO object for table data access.
func NewSysUserDao(group string) *SysUserDao {
	if group == "" {
		group = "default"
	}
	return &SysUserDao{
		group:   group,
		table:   "sys_user",
		columns: sysUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysUserDao) Columns() SysUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
