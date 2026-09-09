package install

import (
	"context"
	_ "embed"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/liuzhengtao/auth-common-backend/internal/config"
	"github.com/liuzhengtao/auth-common-backend/internal/consts"
	"github.com/liuzhengtao/auth-common-backend/internal/dao"
	"github.com/liuzhengtao/auth-common-backend/internal/model/do"
	"github.com/liuzhengtao/auth-common-backend/utility"
)

//go:embed sql/schema.sql
var schemaSQL []byte

var requiredTables = []string{
	"sys_dept",
	"sys_user",
	"sys_role",
	"sys_menu",
	"sys_user_role",
	"sys_role_menu",
	"sys_dict_type",
	"sys_dict",
}

// ensureSchema 检查宿主库表，缺失则建表；无 admin 时初始化种子数据。
func ensureSchema(ctx context.Context) error {
	if ctx == nil {
		ctx = gctx.New()
	}
	db := g.DB(config.Get().DbGroup)
	if db == nil {
		return gerror.Newf("auth-common: database not configured (g.DB(%s) is nil)", config.Get().DbGroup)
	}
	if _, err := db.GetValue(ctx, "SELECT 1"); err != nil {
		return gerror.Wrap(err, "auth-common: cannot connect to database")
	}

	existing, err := tableSet(ctx)
	if err != nil {
		return err
	}

	var missing []string
	for _, name := range requiredTables {
		if !existing[name] {
			missing = append(missing, name)
		}
	}
	if len(missing) > 0 {
		g.Log().Infof(ctx, "auth-common: creating missing tables: %v", missing)
		if err = execSchemaSQL(ctx); err != nil {
			return err
		}
	}

	return seedIfNeeded(ctx)
}

func tableSet(ctx context.Context) (map[string]bool, error) {
	set := make(map[string]bool, len(requiredTables))
	tables, err := g.DB(config.Get().DbGroup).Tables(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "auth-common: list tables failed")
	}
	for _, t := range tables {
		set[strings.ToLower(t)] = true
	}
	return set, nil
}

func execSchemaSQL(ctx context.Context) error {
	stmts := splitSQLStatements(string(schemaSQL))
	for _, stmt := range stmts {
		if _, err := g.DB(config.Get().DbGroup).Exec(ctx, stmt); err != nil {
			return gerror.Wrapf(err, "auth-common: exec schema failed: %s", truncate(stmt, 80))
		}
	}
	return nil
}

func splitSQLStatements(sql string) []string {
	parts := strings.Split(sql, ";")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		lines := strings.Split(p, "\n")
		var kept []string
		for _, line := range lines {
			trim := strings.TrimSpace(line)
			if trim == "" || strings.HasPrefix(trim, "--") {
				continue
			}
			kept = append(kept, line)
		}
		s := strings.TrimSpace(strings.Join(kept, "\n"))
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

func seedIfNeeded(ctx context.Context) error {
	count, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Username, "admin").Count()
	if err != nil {
		return gerror.Wrap(err, "auth-common: check admin user failed")
	}
	if count > 0 {
		g.Log().Debug(ctx, "auth-common: seed skipped, admin already exists")
		return nil
	}

	g.Log().Info(ctx, "auth-common: initializing default seed data")
	passwd := gconv.String(utility.EncryptData(consts.DEFAULT_PASSWORD))

	return g.DB(config.Get().DbGroup).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		return seedData(ctx, tx, passwd)
	})
}

func seedData(ctx context.Context, tx gdb.TX, passwd string) error {
	if _, err := tx.Model("sys_dept").Ctx(ctx).Data(do.SysDept{
		Id:       1,
		Name:     "总公司",
		ParentId: 0,
		TreePath: "0",
		Sort:     1,
		Status:   1,
		Deleted:  0,
	}).Insert(); err != nil {
		return gerror.Wrap(err, "seed sys_dept")
	}

	if _, err := tx.Model("sys_role").Ctx(ctx).Data([]do.SysRole{
		{Id: 1, Name: "超级管理员", Code: consts.ROOT_ROLE_CODE, Sort: 1, Status: 1, DataScope: 0, Deleted: 0},
		{Id: 2, Name: "系统管理员", Code: "ADMIN", Sort: 2, Status: 1, DataScope: 0, Deleted: 0},
	}).Insert(); err != nil {
		return gerror.Wrap(err, "seed sys_role")
	}

	if _, err := tx.Model("sys_user").Ctx(ctx).Data([]do.SysUser{
		{Id: 1, Username: "root", Nickname: "超级管理员", Gender: 1, Password: passwd, DeptId: 1, Status: 1, Deleted: 0},
		{Id: 2, Username: "admin", Nickname: "管理员", Gender: 1, Password: passwd, DeptId: 1, Status: 1, Deleted: 0},
	}).Insert(); err != nil {
		return gerror.Wrap(err, "seed sys_user")
	}

	if _, err := tx.Model("sys_user_role").Ctx(ctx).Data([]do.SysUserRole{
		{UserId: 1, RoleId: 1},
		{UserId: 2, RoleId: 2},
	}).Insert(); err != nil {
		return gerror.Wrap(err, "seed sys_user_role")
	}

	menus := []do.SysMenu{
		{Id: 1, ParentId: 0, TreePath: "0", Name: "系统管理", Type: 2, Path: "/system", Component: "Layout", Visible: 1, Sort: 1, Icon: "system", AlwaysShow: 1, KeepAlive: 0},
		{Id: 2, ParentId: 1, TreePath: "0,1", Name: "用户管理", Type: 1, Path: "user", Component: "system/user/index", Visible: 1, Sort: 1, Icon: "user", KeepAlive: 1},
		{Id: 3, ParentId: 1, TreePath: "0,1", Name: "角色管理", Type: 1, Path: "role", Component: "system/role/index", Visible: 1, Sort: 2, Icon: "role", KeepAlive: 1},
		{Id: 4, ParentId: 1, TreePath: "0,1", Name: "菜单管理", Type: 1, Path: "menu", Component: "system/menu/index", Visible: 1, Sort: 3, Icon: "menu", KeepAlive: 1},
		{Id: 5, ParentId: 1, TreePath: "0,1", Name: "部门管理", Type: 1, Path: "dept", Component: "system/dept/index", Visible: 1, Sort: 4, Icon: "tree", KeepAlive: 1},
		{Id: 6, ParentId: 1, TreePath: "0,1", Name: "字典管理", Type: 1, Path: "dict", Component: "system/dict/index", Visible: 1, Sort: 5, Icon: "dict", KeepAlive: 1},
		{Id: 7, ParentId: 2, TreePath: "0,1,2", Name: "新增用户", Type: 4, Perm: "sys:user:add", Visible: 1, Sort: 1},
		{Id: 8, ParentId: 2, TreePath: "0,1,2", Name: "编辑用户", Type: 4, Perm: "sys:user:edit", Visible: 1, Sort: 2},
		{Id: 9, ParentId: 2, TreePath: "0,1,2", Name: "删除用户", Type: 4, Perm: "sys:user:delete", Visible: 1, Sort: 3},
		{Id: 10, ParentId: 3, TreePath: "0,1,3", Name: "编辑角色", Type: 4, Perm: "sys:role:edit", Visible: 1, Sort: 1},
	}
	if _, err := tx.Model("sys_menu").Ctx(ctx).Data(menus).Insert(); err != nil {
		return gerror.Wrap(err, "seed sys_menu")
	}

	roleMenus := make([]do.SysRoleMenu, 0, len(menus)*2)
	for _, m := range menus {
		roleMenus = append(roleMenus,
			do.SysRoleMenu{RoleId: 1, MenuId: m.Id},
			do.SysRoleMenu{RoleId: 2, MenuId: m.Id},
		)
	}
	if _, err := tx.Model("sys_role_menu").Ctx(ctx).Data(roleMenus).Insert(); err != nil {
		return gerror.Wrap(err, "seed sys_role_menu")
	}

	if _, err := tx.Model("sys_dict_type").Ctx(ctx).Data([]do.SysDictType{
		{Id: 1, Name: "性别", Code: "gender", Status: 0},
		{Id: 2, Name: "状态", Code: "status", Status: 0},
	}).Insert(); err != nil {
		return gerror.Wrap(err, "seed sys_dict_type")
	}

	if _, err := tx.Model("sys_dict").Ctx(ctx).Data([]do.SysDict{
		{Id: 1, TypeCode: "gender", Name: "男", Value: "1", Sort: 1, Status: 1, Defaulted: 1},
		{Id: 2, TypeCode: "gender", Name: "女", Value: "2", Sort: 2, Status: 1, Defaulted: 0},
		{Id: 3, TypeCode: "status", Name: "正常", Value: "1", Sort: 1, Status: 1, Defaulted: 1},
		{Id: 4, TypeCode: "status", Name: "禁用", Value: "0", Sort: 2, Status: 1, Defaulted: 0},
	}).Insert(); err != nil {
		return gerror.Wrap(err, "seed sys_dict")
	}

	return nil
}
