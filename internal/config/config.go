package config

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/liuzhengtao/auth-common-backend/internal/consts"
)

// Config auth-common 插件配置，可通过宿主 config.yaml 的 authCommon 节点覆盖。
type Config struct {
	JwtSecret   string `json:"jwtSecret"`
	JwtAesKey   string `json:"jwtAesKey"`
	EncryptSalt string `json:"encryptSalt"`
	RoutePrefix string `json:"routePrefix"`
}

var (
	mu     sync.RWMutex
	loaded bool
	cfg    Config
)

func defaultConfig() Config {
	return Config{
		JwtSecret:   consts.JwtSecretKey,
		JwtAesKey:   consts.JwtAesKey,
		EncryptSalt: consts.EncryptSaltKey,
		RoutePrefix: "/api/v1",
	}
}

// Load 从宿主配置加载（幂等）。优先使用 authCommon 节点，缺省回落 consts。
func Load(ctx context.Context) Config {
	mu.Lock()
	defer mu.Unlock()
	if loaded {
		return cfg
	}
	cfg = defaultConfig()
	if ctx == nil {
		ctx = context.Background()
	}
	v, err := g.Cfg().Get(ctx, "authCommon")
	if err == nil && !v.IsNil() && !v.IsEmpty() {
		_ = v.Scan(&cfg)
		if cfg.JwtSecret == "" {
			cfg.JwtSecret = consts.JwtSecretKey
		}
		if cfg.JwtAesKey == "" {
			cfg.JwtAesKey = consts.JwtAesKey
		}
		if cfg.EncryptSalt == "" {
			cfg.EncryptSalt = consts.EncryptSaltKey
		}
		if cfg.RoutePrefix == "" {
			cfg.RoutePrefix = "/api/v1"
		}
	}
	loaded = true
	return cfg
}

// Get 返回已加载配置；未加载时用默认值并尝试从配置文件读取。
func Get() Config {
	mu.RLock()
	if loaded {
		c := cfg
		mu.RUnlock()
		return c
	}
	mu.RUnlock()
	return Load(context.Background())
}

// SetRoutePrefix 由 Install Option 覆盖路由前缀。
func SetRoutePrefix(prefix string) {
	if prefix == "" {
		return
	}
	mu.Lock()
	defer mu.Unlock()
	if !loaded {
		cfg = defaultConfig()
		loaded = true
	}
	cfg.RoutePrefix = prefix
}
