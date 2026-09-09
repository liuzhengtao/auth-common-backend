package sysUser

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"

	"github.com/liuzhengtao/auth-common-backend/api/common"
	"github.com/liuzhengtao/auth-common-backend/api/v1/user"
	"github.com/liuzhengtao/auth-common-backend/internal/consts"
	"github.com/liuzhengtao/auth-common-backend/internal/dao"
	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/model/do"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
	"github.com/liuzhengtao/auth-common-backend/utility"
)

type sSysUserService struct {
	logger *glog.Logger
}

func New() *sSysUserService {
	return &sSysUserService{
		logger: g.Log().Line(true),
	}
}

func init() {
	service.RegisterSysUserService(New())
}

func (s *sSysUserService) ListPageUsers(ctx context.Context, queryParams *model.UserPageQueryInput) (rows []*user.UserPageVO, total int, err error) {
	pageNum := queryParams.PageNum
	pageSize := queryParams.PageSize
	commonPage := common.NewPageWithCurrentSize(pageNum, pageSize)
	page := &model.UserPage{Page: commonPage}
	err = dao.SysUser.ListPageUsers(ctx, page, queryParams)
	if err != nil {
		s.logger.Error(ctx, "ListPageUsers查询数据失败", err)
		return nil, 0, gerror.New(consts.SYSTEM_EXECUTION_ERROR)
	}
	total = page.Total
	rows = make([]*user.UserPageVO, 0)
	for _, record := range page.Records {
		m := record.(map[string]interface{})
		genderLabel := "男"
		if gconv.Int(m["gender"]) == 2 {
			genderLabel = "女"
		} else if gconv.Int(m["gender"]) == 0 {
			genderLabel = "未知"
		}
		rows = append(rows, &user.UserPageVO{
			Id:          gconv.Int(m["id"]),
			Username:    gconv.String(m["username"]),
			Nickname:    gconv.String(m["nickname"]),
			Mobile:      gconv.String(m["mobile"]),
			GenderLabel: genderLabel,
			Avatar:      gconv.String(m["avatar"]),
			Status:      gconv.Int(m["status"]),
			DeptName:    gconv.String(m["dept_name"]),
			RoleNames:   gconv.String(m["role_names"]),
			CreateTime:  gconv.String(m["create_time"]),
		})
	}
	return
}

func (s *sSysUserService) CreateUser(ctx context.Context, in *model.CreateUserInput) (lastInsertId int64, err error) {
	var sysUser *do.SysUser
	err = gconv.Struct(in, &sysUser)
	if err != nil {
		s.logger.Error(ctx, "Struct参数转换失败")
		return 0, gerror.New(consts.PARAM_TRANSFORM_ERROR)
	}
	if gstr.Equal(in.Password, "") {
		sysUser.Password = consts.DEFAULT_PASSWORD
	}
	sysUser.Password = utility.EncryptData(sysUser.Password)
	lastInsertId, err = dao.SysUser.Ctx(ctx).Data(sysUser).InsertAndGetId()
	if err != nil {
		s.logger.Error(ctx, "InsertAndGetId插入数据失败", err)
		return 0, gerror.New(consts.DATABASE_ERROR)
	}
	return
}

func (s *sSysUserService) GetUserFormData(ctx context.Context, userId int64) (form *user.Form, err error) {
	err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, userId).Where(dao.SysUser.Columns().Deleted, 0).Fields("id,username,nickname,mobile,gender,avatar,email,status,dept_id as deptId").Scan(&form)
	if err != nil {
		s.logger.Error(ctx, "Scan查询数据失败", err)
		return nil, gerror.New(consts.PARAM_TRANSFORM_ERROR)
	}
	values, err := dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.Columns().UserId, userId).Fields("role_id as roleIds").Array()
	if err != nil {
		return nil, err
	}
	form.RoleIds = make([]int, 0)
	for _, value := range values {
		form.RoleIds = append(form.RoleIds, value.Int())
	}
	return
}

func (s *sSysUserService) GetUser(ctx context.Context, userId int64) (user *entity.SysUser, err error) {
	err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, userId).Scan(&user)
	if err != nil {
		s.logger.Error(ctx, "数据库查询失败", err)
		return nil, gerror.New(consts.DATABASE_ERROR)
	}
	return
}

func (s *sSysUserService) UpdateUser(ctx context.Context, in *model.UpdateUserInput) (err error) {
	var sysUser *do.SysUser
	err = gconv.Struct(in, &sysUser)
	if err != nil {
		s.logger.Error(ctx, "Struct参数转换失败")
		return gerror.New(consts.PARAM_TRANSFORM_ERROR)
	}
	_, err = dao.SysUser.Ctx(ctx).Data(sysUser).Where(dao.SysUser.Columns().Id, in.UserId).OmitEmpty().Update()
	if err != nil {
		s.logger.Error(ctx, "Update更新数据失败", err)
		return gerror.New(consts.DATABASE_ERROR)
	}
	// 保存用户角色
	err = service.SysUserRoleService().SaveUserRoles(ctx, gconv.Int(in.UserId), in.RoleIds)
	if err != nil {
		s.logger.Error(ctx, "SaveUserRoles更新数据失败", err)
		return gerror.New(consts.DATABASE_ERROR)
	}
	return
}

func (s *sSysUserService) DeleteUsers(ctx context.Context, in *model.DeleteUserInput) (err error) {
	ids := gstr.SplitAndTrim(in.Ids, ",")
	result, err := dao.SysUser.Ctx(ctx).WhereIn(dao.SysUser.Columns().Id, ids).Data(do.SysUser{Deleted: 1}).Update()
	if err != nil {
		s.logger.Error(ctx, "DeleteUsers更新数据失败", err)
		return gerror.New(consts.DATABASE_ERROR)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected > 0 {
		return service.SysUserRoleService().DeleteByUserIds(ctx, ids)
	}
	return
}

func (s *sSysUserService) UpdatePasswd(ctx context.Context, in *model.UpdatePasswdInput) (err error) {
	//密码加密以后修改数据
	passwd := utility.EncryptData(in.Password)
	_, err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, in.UserId).Data(do.SysUser{Password: passwd}).Update()
	if err != nil {
		s.logger.Error(ctx, "UpdatePasswd更新数据失败", err)
		return gerror.New(consts.DATABASE_ERROR)
	}
	return
}

func (s *sSysUserService) UpdateStatus(ctx context.Context, in *model.UpdateStatusInput) (err error) {
	_, err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, in.UserId).Data(do.SysUser{Status: in.Status}).Update()
	if err != nil {
		s.logger.Error(ctx, "UpdateStatus更新数据失败", err)
		return gerror.New(consts.DATABASE_ERROR)
	}
	return
}

func (s *sSysUserService) GetCurrentUserInfo(ctx context.Context) (out *model.GetCurrentInfoOutput, err error) {
	authUserInfo := service.Auth().GetIdentity(ctx)
	sysUser, err := s.GetUser(ctx, gconv.Int64(authUserInfo.UserId))
	if err != nil {
		return nil, err
	}
	out = &model.GetCurrentInfoOutput{
		UserInfoVO: &user.UserInfoVO{
			UserId:   gconv.Int64(authUserInfo.UserId),
			Nickname: sysUser.Nickname,
			Avatar:   sysUser.Avatar,
			Username: sysUser.Username,
		},
	}
	out.Roles = authUserInfo.Roles
	out.Perms = authUserInfo.Perms
	return
}

func (s *sSysUserService) GetUserToToken(ctx context.Context, username, password string) (user *entity.SysUser, err error) {
	passwd := utility.EncryptData(password)
	err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Username, username).Scan(&user)
	if err != nil {
		s.logger.Error(ctx, "用户名查询失败", username, err)
		return nil, gerror.New(consts.DATABASE_ERROR)
	}
	if user == nil {
		s.logger.Error(ctx, "用户名不存在", username)
		return nil, gerror.New(consts.USER_NOT_EXIST)
	}
	if user.Status == 0 {
		s.logger.Error(ctx, "用户已被禁用", username)
		return nil, gerror.New(consts.USER_ACCOUNT_LOCKED)
	}
	if !gstr.Equal(user.Password, gconv.String(passwd)) {
		s.logger.Error(ctx, "密码错误", username)
		return nil, gerror.New(consts.USER_PASSWORD_ERROR)
	}
	return
}

func (s *sSysUserService) GetUserAuthInfo(ctx context.Context, username string) (*model.UserAuthInfo, error) {
	info, err := dao.SysUser.GetUserAuthInfo(ctx, username)
	if err != nil {
		return nil, err
	}
	if info != nil {
		roles := info.Roles
		if !gutil.IsEmpty(roles) {
			perms, err := service.SysMenusService().ListRolePerms(ctx, roles)
			if err != nil {
				return nil, err
			}
			info.Perms = perms
		}
		dataScope, err := service.SysRoleService().GetMaximumDataScope(ctx, roles)
		if err != nil {
			return nil, err
		}
		info.DataScope = dataScope
	}
	return info, nil
}
