package install

// Option Install 可选配置。
type Option func(*options)

type options struct {
	prefix         string
	dbGroup        string
	skipSchemaInit bool
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
