package install

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/liuzhengtao/auth-common-backend/internal/applog"
	"github.com/liuzhengtao/auth-common-backend/internal/config"
	"github.com/liuzhengtao/auth-common-backend/internal/controller"
	"github.com/liuzhengtao/auth-common-backend/internal/dao"
	_ "github.com/liuzhengtao/auth-common-backend/internal/logic"
	"github.com/liuzhengtao/auth-common-backend/internal/logic/captcha"
	"github.com/liuzhengtao/auth-common-backend/middleware"
)

type sInstall struct {
	version      string
	name         string
	author       string
	description  string
	pluginServer *ghttp.Server
}

func NewInstall() *sInstall {
	return &sInstall{
		version:     "1.0",
		name:        "github.com/liuzhengtao/auth-common-backend",
		author:      "boby",
		description: "auth-common GoFrame plugin",
	}
}

func (s *sInstall) Name() string        { return s.name }
func (s *sInstall) Version() string     { return s.version }
func (s *sInstall) Description() string { return s.description }
func (s *sInstall) Author() string      { return s.author }
func (s *sInstall) Remove() error       { return nil }

// Install 兼容旧用法：等价于包级 Install(server)。
func (s *sInstall) Install(server *ghttp.Server) error {
	s.pluginServer = server
	return Install(server)
}

// Install 将鉴权与系统管理 API 挂载到宿主 Server。
// 默认检查宿主数据库表，缺失则建表并初始化种子数据。
func Install(server *ghttp.Server, opts ...Option) error {
	if server == nil {
		return gerror.New("auth-common: server is nil")
	}
	o := applyOptions(opts...)

	ctx := gctx.New()
	cfg := config.Load(ctx)
	if o.prefix != "" {
		config.SetRoutePrefix(o.prefix)
	}
	if o.dbGroup != "" {
		config.SetDBGroup(o.dbGroup)
	}
	if o.redisGroup != "" {
		config.SetRedisGroup(o.redisGroup)
	}
	if o.distributed != nil {
		config.SetDistributed(*o.distributed)
	}
	cfg = config.Get()
	dao.Init(cfg.DbGroup)

	if cfg.Distributed {
		if err := ensureRedis(ctx, cfg.RedisGroup); err != nil {
			return err
		}
		captcha.UseRedisStore(cfg.RedisGroup)
	}

	if !o.skipSchemaInit {
		if err := ensureSchema(ctx); err != nil {
			return err
		}
	}

	registerRoutes(server, cfg.RoutePrefix)
	return nil
}

// ensureRedis 校验宿主 Redis 配置组可用（需宿主 blank-import redis 驱动）。
func ensureRedis(ctx context.Context, group string) error {
	if group == "" {
		group = "default"
	}
	if _, err := g.Redis(group).Do(ctx, "PING"); err != nil {
		return gerror.Wrapf(err, "auth-common: redis group %q unavailable (blank-import github.com/gogf/gf/contrib/nosql/redis/v2 and configure redis.%s)", group, group)
	}
	return nil
}

func applyOptions(opts ...Option) *options {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}
	if o.logger != nil {
		applog.Set(o.logger)
	}
	return o
}

func registerRoutes(server *ghttp.Server, prefix string) {
	if prefix == "" {
		prefix = "/api/v1"
	}
	server.Group(prefix, func(group *ghttp.RouterGroup) {
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.ErrorHandler)
			group.Group("/auth", func(group *ghttp.RouterGroup) {
				group.Bind(controller.NewAuthController())
			})
			group.Middleware(middleware.Auth)
			group.DELETE("/auth/logout", controller.NewAuthController().Logout)
			group.Group("/users", func(group *ghttp.RouterGroup) {
				group.Bind(controller.NewSysUserController())
			})
			group.Group("/menus", func(group *ghttp.RouterGroup) {
				group.Bind(controller.NewMenusController())
			})
			group.Group("/dept", func(group *ghttp.RouterGroup) {
				group.Bind(controller.NewSysDeptController())
			})
			group.Group("/dict", func(group *ghttp.RouterGroup) {
				group.Bind(controller.NewSysDictController())
			})
			group.Group("/roles", func(group *ghttp.RouterGroup) {
				group.Bind(controller.NewRoleController())
			})
		})
	})
}

// MustInstall 同 Install，失败时 panic。
func MustInstall(server *ghttp.Server, opts ...Option) {
	if err := Install(server, opts...); err != nil {
		panic(err)
	}
}

// EnsureSchema 仅执行检表/建表/种子，不注册路由（便于单独初始化）。
func EnsureSchema(ctx context.Context, opts ...Option) error {
	o := applyOptions(opts...)
	cfg := config.Load(ctx)
	if o.dbGroup != "" {
		config.SetDBGroup(o.dbGroup)
	}
	cfg = config.Get()
	dao.Init(cfg.DbGroup)
	return ensureSchema(ctx)
}
