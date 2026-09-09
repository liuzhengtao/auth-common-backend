package bizctx

import (
	"context"

	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

const (
	sessionKeyUser   = "SessionKeyUser"   // 用户信息存放在Session中的Key
	sessionKeyNotice = "SessionKeyNotice" // 存放在Session中的提示信息，往往使用后则删除
)

type sSession struct {
}

func New() *sSession {
	return &sSession{}
}

func (s *sSession) SetUser(ctx context.Context, user *entity.SysUser) error {
	return service.BizCtx().Get(ctx).Session.Set(sessionKeyUser, user)
}

func (s *sSession) GetUser(ctx context.Context) *entity.SysUser {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		v, _ := customCtx.Session.Get(sessionKeyUser)
		if !v.IsNil() {
			var user *entity.SysUser
			_ = v.Struct(&user)
			return user
		}
	}
	return nil
}

func (s *sSession) RemoveUser(ctx context.Context) error {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyUser)
	}
	return nil
}

// SetNotice 设置Notice
func (s *sSession) SetNotice(ctx context.Context, message *model.SessionNotice) error {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Set(sessionKeyNotice, message)
	}
	return nil
}

// GetNotice 获取Notice
func (s *sSession) GetNotice(ctx context.Context) (*model.SessionNotice, error) {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		var message *model.SessionNotice
		v, err := customCtx.Session.Get(sessionKeyNotice)
		if err != nil {
			return nil, err
		}
		if err = v.Scan(&message); err != nil {
			return nil, err
		}
		return message, nil
	}
	return nil, nil
}

// RemoveNotice 删除Notice
func (s *sSession) RemoveNotice(ctx context.Context) error {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyNotice)
	}
	return nil
}
func (s *sSession) IsLogin(ctx context.Context) (*entity.SysUser, bool) {
	user := s.GetUser(ctx)
	if user.Id == 0 {
		return nil, false
	} else {
		return user, true
	}
}
