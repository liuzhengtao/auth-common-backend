// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysRoleDao is the data access object for table sys_role.
type SysRoleDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysRoleColumns // columns contains all the column names of Table for convenient usage.
}

// SysRoleColumns defines and stores column names for table sys_role.
type SysRoleColumns struct {
	Id         string //
	Name       string // 角色名称
	Code       string // 角色编码
	Sort       string // 显示顺序
	Status     string // 角色状态(1-正常；0-停用)
	DataScope  string // 数据权限(0-所有数据；1-部门及子部门数据；2-本部门数据；3-本人数据)
	Deleted    string // 逻辑删除标识(0-未删除；1-已删除)
	CreateTime string // 更新时间
	UpdateTime string // 创建时间
}

// sysRoleColumns holds the columns for table sys_role.
var sysRoleColumns = SysRoleColumns{
	Id:         "id",
	Name:       "name",
	Code:       "code",
	Sort:       "sort",
	Status:     "status",
	DataScope:  "data_scope",
	Deleted:    "deleted",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewSysRoleDao creates and returns a new DAO object for table data access.
func NewSysRoleDao(group string) *SysRoleDao {
	if group == "" {
		group = "default"
	}
	return &SysRoleDao{
		group:   group,
		table:   "sys_role",
		columns: sysRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysRoleDao) Columns() SysRoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
