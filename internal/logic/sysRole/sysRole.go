package sysRole

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/liuzhengtao/auth-common-backend/api/common"
	"github.com/liuzhengtao/auth-common-backend/api/v1/dept"
	"github.com/liuzhengtao/auth-common-backend/api/v1/role"
	"github.com/liuzhengtao/auth-common-backend/internal/consts"
	"github.com/liuzhengtao/auth-common-backend/internal/dao"
	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/model/do"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

type sSysRoleService struct{}

func New() *sSysRoleService {
	return &sSysRoleService{}
}

func init() {
	service.RegisterSysRoleService(New())
}

func (s *sSysRoleService) ListPage(ctx context.Context, keywords string, pageNum int, pageSize int) (out *model.ListOutput, err error) {
	out = &model.ListOutput{}
	req := role.PageQueryReq{
		Keyword: keywords,
	}
	req.BasePageQuery = &common.BasePageQuery{
		PageNum:  pageNum,
		PageSize: pageSize,
	}
	err = dao.SysRole.ListPage(ctx, req, out)
	if err != nil {
		return nil, gerror.New(consts.DATABASE_ERROR)
	}
	return
}

func (s *sSysRoleService) GetMaximumDataScope(ctx context.Context, roles []string) (dataScope int, err error) {
	orm := dao.SysRole.Ctx(ctx)
	if len(roles) > 0 {
		orm = orm.WhereIn("code", roles)
	} else {
		orm = orm.Where("id", "-1")
	}
	min, err := orm.Min("data_scope")
	if err != nil {
		return 0, err
	}
	return gconv.Int(min), nil
}

func (s *sSysRoleService) RoleOptionsList(ctx context.Context) (list []dept.Option, err error) {
	var roles []entity.SysRole
	err = dao.SysRole.Ctx(ctx).Fields("id,name").WhereNot("code", consts.ROOT_ROLE_CODE).OrderAsc("sort").Scan(&roles)
	if err != nil {
		return nil, err
	}
	for _, roleItem := range roles {
		list = append(list, dept.Option{
			Value: roleItem.Id,
			Label: roleItem.Name,
		})
	}
	return
}

func (s *sSysRoleService) CreateRole(ctx context.Context, input *model.CreateRoleInput) (lastInsertId int64, err error) {
	var sysRole *do.SysRole
	err = gconv.Struct(input, &sysRole)
	if err != nil {
		return 0, err
	}
	return dao.SysRole.Ctx(ctx).Data(sysRole).InsertAndGetId()
}

func (s *sSysRoleService) UpdateRole(ctx context.Context, input *model.UpdateRoleInput) (err error) {
	var sysRole *do.SysRole
	err = gconv.Struct(input, &sysRole)
	if err != nil {
		return err
	}
	_, err = dao.SysRole.Ctx(ctx).Data(sysRole).Where(dao.SysRole.Columns().Id, input.Id).OmitEmpty().Update()
	return
}

func (s *sSysRoleService) DeleteRole(ctx context.Context, in *model.DeleteRoleInput) (err error) {
	ids := gstr.SplitAndTrim(in.Ids, ",")
	for _, id := range ids {
		sysRole, err := s.GetRoleInfo(ctx, gconv.Int64(id))
		if err != nil {
			return err
		} else if sysRole == nil || sysRole.Deleted == 1 {
			return gerror.New("角色不存在")
		}
		//判断是否该角色关联了用户
		users, err := service.SysUserRoleService().HasAssignedUsers(ctx, gconv.Int64(id))
		if err != nil {
			return err
		}
		if users {
			return gerror.Newf("角色【%s】已分配用户，请先解除关联", sysRole.Name)
		}
	}
	result, err := dao.SysRole.Ctx(ctx).WhereIn(dao.SysRole.Columns().Id, ids).Delete()
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected > 0 {
		return service.SysUserRoleService().DeleteByRoleIds(ctx, ids)
	}
	return
}

func (s *sSysRoleService) ListMenuIdsByRoleId(ctx context.Context, roleId int64) (menuIds []int64, err error) {
	array, err := dao.SysRoleMenu.Ctx(ctx).As("rm").InnerJoin("sys_menu m", "rm.menu_id = m.id").Fields("rm.menu_id").Where("rm.role_id", roleId).Array()
	menuIds = make([]int64, 0)
	for _, value := range array {
		menuIds = append(menuIds, value.Int64())
	}
	if err != nil {
		return nil, err
	}
	return
}

func (s *sSysRoleService) GetRoleInfo(ctx context.Context, id int64) (sysRole *entity.SysRole, err error) {
	err = dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().Id, id).Scan(&sysRole)
	if err != nil {
		return nil, err
	}
	return
}

func (s *sSysRoleService) AssignMenusToRole(ctx context.Context, in *model.AssignMenusToRoleInput) (err error) {
	sysRole, err := service.SysRoleService().GetRoleInfo(ctx, in.RoleId)
	if err != nil {
		return err
	}
	if sysRole == nil {
		return gerror.New("角色不存在")
	}
	//删除角色菜单关联
	err = service.SysRoleMenu().RemoveRoleMenus(ctx, in.RoleId)
	if err != nil {
		return err
	}
	//新增角色菜单关联
	if len(in.Body) > 0 {
		roleMenus := make([]do.SysRoleMenu, 0)
		for _, menuId := range in.Body {
			roleMenus = append(roleMenus, do.SysRoleMenu{
				RoleId: in.RoleId,
				MenuId: menuId,
			})
		}
		err = service.SysRoleMenu().SaveBatch(ctx, roleMenus)
		if err != nil {
			return err
		}
	}
	return
}
