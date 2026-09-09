// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
)

type (
	IBizCtx interface {
		// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
		Init(r *ghttp.Request, customCtx *model.Context)
		// Get 获得上下文变量，如果没有设置，那么返回nil
		Get(ctx context.Context) *model.Context
		// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
		SetUser(ctx context.Context, ctxUser *entity.SysUser)
		// SetData 将上下文信息设置到上下文请求中，注意是完整覆盖
		SetData(ctx context.Context, data g.Map)
	}
	ISession interface {
		SetUser(ctx context.Context, user *entity.SysUser) error
		GetUser(ctx context.Context) *entity.SysUser
		RemoveUser(ctx context.Context) error
		// SetNotice 设置Notice
		SetNotice(ctx context.Context, message *model.SessionNotice) error
		// GetNotice 获取Notice
		GetNotice(ctx context.Context) (*model.SessionNotice, error)
		// RemoveNotice 删除Notice
		RemoveNotice(ctx context.Context) error
		IsLogin(ctx context.Context) (*entity.SysUser, bool)
	}
)

var (
	localBizCtx  IBizCtx
	localSession ISession
)

func BizCtx() IBizCtx {
	if localBizCtx == nil {
		panic("implement not found for interface IBizCtx, forgot register?")
	}
	return localBizCtx
}

func RegisterBizCtx(i IBizCtx) {
	localBizCtx = i
}

func Session() ISession {
	if localSession == nil {
		panic("implement not found for interface ISession, forgot register?")
	}
	return localSession
}

func RegisterSession(i ISession) {
	localSession = i
}
