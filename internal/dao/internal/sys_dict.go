// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDictDao is the data access object for table sys_dict.
type SysDictDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysDictColumns // columns contains all the column names of Table for convenient usage.
}

// SysDictColumns defines and stores column names for table sys_dict.
type SysDictColumns struct {
	Id         string // 主键
	TypeCode   string // 字典类型编码
	Name       string // 字典项名称
	Value      string // 字典项值
	Sort       string // 排序
	Status     string // 状态(1:正常;0:禁用)
	Defaulted  string // 是否默认(1:是;0:否)
	Remark     string // 备注
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// sysDictColumns holds the columns for table sys_dict.
var sysDictColumns = SysDictColumns{
	Id:         "id",
	TypeCode:   "type_code",
	Name:       "name",
	Value:      "value",
	Sort:       "sort",
	Status:     "status",
	Defaulted:  "defaulted",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewSysDictDao creates and returns a new DAO object for table data access.
func NewSysDictDao() *SysDictDao {
	return &SysDictDao{
		group:   "default",
		table:   "sys_dict",
		columns: sysDictColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysDictDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysDictDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysDictDao) Columns() SysDictColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysDictDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysDictDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysDictDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
