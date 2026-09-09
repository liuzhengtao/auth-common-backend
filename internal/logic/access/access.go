package access

import (
	"context"

	casbin "github.com/dobyte/gf-casbin"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/liuzhengtao/auth-common-backend/internal/config"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
)

type sAccess struct {
	Enforce *casbin.Enforcer
}

func New() *sAccess {
	accessS := sAccess{}
	ctx := gctx.New()
	cfgVal, _ := g.Cfg().Get(ctx, "casbin")
	var options *casbin.Options
	db := g.DB(config.Get().DbGroup)
	confErr := gconv.Scan(cfgVal, &options)
	if confErr != nil {
		panic(confErr)
	}
	options.DB = db
	enforce, err := casbin.NewEnforcer(options)
	if err != nil {
		panic(err)
	}
	accessS.Enforce = enforce
	return &accessS
}
func (s *sAccess) CheckAccess(ctx context.Context, user *entity.SysUser, path string) (isPass bool, err error) {
	return
}
