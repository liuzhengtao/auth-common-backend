package install

import "github.com/gogf/gf/v2/os/glog"

// Option Install 可选配置。
type Option func(*options)

type options struct {
	prefix         string
	dbGroup        string
	skipSchemaInit bool
	logger         glog.ILogger
	distributed    *bool
	redisGroup     string
}

func defaultOptions() *options {
	return &options{}
}

// WithPrefix 覆盖路由前缀，默认 /api/v1（也可通过 config authCommon.routePrefix 配置）。
func WithPrefix(prefix string) Option {
	return func(o *options) {
		o.prefix = prefix
	}
}

// WithDBGroup 覆盖数据库配置组名，默认 default（也可通过 config authCommon.dbGroup 配置）。
func WithDBGroup(group string) Option {
	return func(o *options) {
		o.dbGroup = group
	}
}

// WithSkipSchemaInit 跳过检表/建表/种子初始化（宿主自行管理库表时使用）。
func WithSkipSchemaInit(skip bool) Option {
	return func(o *options) {
		o.skipSchemaInit = skip
	}
}

// WithLogger 注入本库使用的 Logger；未传入时懒加载 glog.New().Line(true)，不使用 g.Log() 单例。
func WithLogger(l glog.ILogger) Option {
	return func(o *options) {
		o.logger = l
	}
}

// WithDistributed 开启分布式模式（Redis 共享验证码 + Token 黑名单）；也可通过 config authCommon.distributed 配置。
func WithDistributed(enabled bool) Option {
	return func(o *options) {
		o.distributed = &enabled
	}
}

// WithRedisGroup 覆盖 Redis 配置组名，默认 default（也可通过 config authCommon.redisGroup 配置）。
func WithRedisGroup(group string) Option {
	return func(o *options) {
		o.redisGroup = group
	}
}
