package auth

import (
	"context"

	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

func GetUser(ctx context.Context) (*entity.SysUser, error) {
	user, err := service.Auth().GetUser(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetIdentity(ctx context.Context) *model.UserAuthInfo {
	identity := service.Auth().GetIdentity(ctx)
	if identity == nil {
		return nil
	}
	return identity
}

func GetRoles(ctx context.Context) (roles []string, err error) {
	roles, err = service.Auth().GetRoles(ctx)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
