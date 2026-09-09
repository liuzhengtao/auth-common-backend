package user

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/liuzhengtao/auth-common-backend/api/common"
)

type Form struct {
	ID       int64  `p:"id" dc:"用户ID" json:"id,omitempty"`
	UserName string `p:"username" v:"required#用户名不能为空" json:"username,omitempty" `
	NickName string `p:"nickname" v:"required#昵称不能为空" json:"nickname,omitempty"`
	Mobile   string `p:"mobile" v:"required|phone#手机号格式不正确" json:"mobile,omitempty"`
	Gender   int    `p:"gender" dc:"性别" json:"gender,omitempty"`
	Avatar   string `p:"avatar" dc:"用户头像" json:"avatar,omitempty"`
	Email    string `p:"email" dc:"邮箱" json:"email,omitempty"`
	Status   int    `p:"status" v:"in:0,1" dc:"用户状态(1:正常;0:禁用)" json:"status,omitempty"`
	DeptId   int    `p:"deptId" dc:"部门ID" json:"deptId,omitempty"`
	RoleIds  []int  `p:"roleIds" dc:"角色ID集合" json:"roleIds,omitempty"`
}

type PageQueryReq struct {
	g.Meta    `path:"/page" method:"get" tags:"用户接口" dc:"用户分页查询"`
	Keyword   string `p:"keywords" dc:"关键字(用户名/昵称/手机号)"`
	Status    int    `p:"status" d:"4"  dc:"用户状态"`
	DeptId    int64  `p:"deptId" dc:"部门ID"`
	StartTime string `p:"startTime" dc:"开始时间"`
	EndTime   string `p:"endTime" dc:"结束时间"`
	common.BasePageQuery
}

type UserPageVO struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	Mobile      string `json:"mobile"`
	GenderLabel string `json:"genderLabel"`
	Avatar      string `json:"avatar"`
	Email       string `json:"email"`
	Status      int    `json:"status"`
	DeptName    string `json:"deptName"`
	RoleNames   string `json:"roleNames"`
	CreateTime  string `json:"createTime"`
}

type ListPagedUsersRes struct {
	List  []*UserPageVO `json:"list"`
	Total int           `json:"total"`
}

type SaveUserReq struct {
	g.Meta `path:"/" method:"post" tags:"用户接口" dc:"用户新增"`
	Form
	Password string `p:"password"  dc:"密码"`
}

type GetUserFormReq struct {
	g.Meta `path:"/{userId}/form" method:"get" tags:"用户接口" dc:"用户表单数据"`
	UserId int64 `p:"userId" in:"path" v:"required#用户ID不能为空" dc:"用户ID"`
}

type GetUserFormRes struct {
	*Form
}

type NoResultRes struct {
}
type UpdateUserReq struct {
	g.Meta `path:"/{userId}" method:"put" tags:"用户接口" dc:"用户更新"`
	UserId int64 `p:"userId" in:"path" v:"required#用户ID不能为空" dc:"用户ID"`
	Form
}

type DeleteUserReq struct {
	g.Meta `path:"/{ids}" method:"delete" tags:"用户接口" dc:"删除用户"`
	Ids    string `p:"ids" in:"path" v:"required#要删除的ids不能为空" dc:"用户ID，多个以英文逗号(,)分割"`
}

type UpdatePasswdReq struct {
	g.Meta   `path:"/{userId}/password" method:"patch" tags:"用户接口" dc:"修改密码"`
	UserId   int64  `p:"userId" in:"path" v:"required#用户ID不能为空" dc:"用户ID"`
	Password string `p:"password" v:"required#密码不能为空" dc:"新密码"`
}

type UpdateUserStatusReq struct {
	g.Meta `path:"/{userId}/status" method:"patch" tags:"用户接口" dc:"修改用户状态"`
	UserId int64 `p:"userId" in:"path" v:"required#用户ID不能为空" dc:"用户ID"`
	Status int   `p:"status" v:"in:0,1" dc:"用户状态(1:正常;0:禁用)"`
}

type GetCurrentUserInfoReq struct {
	g.Meta `path:"/me" method:"get" tags:"用户接口" dc:"获取当前登录用户信息"`
}

type GetCurrentUserInfoRes struct {
	*UserInfoVO
}

type UserInfoVO struct {
	UserId   int64    `json:"userId" dc:"用户ID"`
	Username string   `json:"username" dc:"用户名"`
	Nickname string   `json:"nickname" dc:"昵称"`
	Avatar   string   `json:"avatar" dc:"头像"`
	Roles    []string `json:"roles" dc:"用户角色编码集合"`
	Perms    []string `json:"perms" dc:"用户权限标识集合"`
}
