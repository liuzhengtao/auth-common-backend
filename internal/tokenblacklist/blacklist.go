package tokenblacklist

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v4"

	"github.com/liuzhengtao/auth-common-backend/internal/config"
)

const keyPrefix = "auth:token:bl:"

// defaultTTL 与 gf-jwt Timeout（120 分钟）对齐，解析 exp 失败时兜底。
const defaultTTL = 120 * time.Minute

func redisGroup() string {
	group := config.Get().RedisGroup
	if group == "" {
		return "default"
	}
	return group
}

func hashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

func key(token string) string {
	return keyPrefix + hashToken(token)
}

// RemainingTTL 解析 JWT exp 得到剩余有效期；失败则返回 defaultTTL。
func RemainingTTL(token string) time.Duration {
	token = strings.TrimSpace(token)
	if token == "" {
		return defaultTTL
	}
	parser := jwt.NewParser(jwt.WithoutClaimsValidation())
	claims := jwt.MapClaims{}
	_, _, err := parser.ParseUnverified(token, claims)
	if err != nil {
		return defaultTTL
	}
	exp, ok := claims["exp"]
	if !ok {
		return defaultTTL
	}
	var expUnix int64
	switch v := exp.(type) {
	case float64:
		expUnix = int64(v)
	case int64:
		expUnix = v
	default:
		return defaultTTL
	}
	remain := time.Until(time.Unix(expUnix, 0))
	if remain <= 0 {
		return time.Second
	}
	return remain
}

// Add 将 token 写入黑名单，TTL 为剩余有效期。
func Add(ctx context.Context, token string) error {
	token = strings.TrimSpace(token)
	if token == "" {
		return nil
	}
	ttl := RemainingTTL(token)
	secs := int64(ttl.Seconds())
	if secs < 1 {
		secs = 1
	}
	err := g.Redis(redisGroup()).SetEX(ctx, key(token), "1", secs)
	return err
}

// IsBlocked 判断 token 是否在黑名单中。
func IsBlocked(ctx context.Context, token string) bool {
	token = strings.TrimSpace(token)
	if token == "" {
		return false
	}
	v, err := g.Redis(redisGroup()).Get(ctx, key(token))
	if err != nil || v.IsNil() || v.IsEmpty() {
		return false
	}
	return true
}

// ExtractToken 按与 gf-jwt 一致的查找顺序提取原始 token：header Authorization / query token / cookie jwt。
func ExtractToken(r *ghttp.Request) string {
	if r == nil {
		return ""
	}
	auth := r.Header.Get("Authorization")
	if auth != "" {
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) == 2 && strings.EqualFold(parts[0], "Bearer") {
			return strings.TrimSpace(parts[1])
		}
		return strings.TrimSpace(auth)
	}
	if t := strings.TrimSpace(r.Get("token").String()); t != "" {
		return t
	}
	if t := strings.TrimSpace(r.Cookie.Get("jwt").String()); t != "" {
		return t
	}
	return ""
}
