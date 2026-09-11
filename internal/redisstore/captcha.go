package redisstore

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mojocn/base64Captcha"
)

const captchaKeyPrefix = "auth:captcha:"

// CaptchaStore 基于宿主 Redis 的 base64Captcha.Store 实现。
type CaptchaStore struct {
	group  string
	ttl    time.Duration
	ctxFun func() context.Context
}

// NewCaptchaStore 创建 Redis 验证码存储；group 对应宿主 redis 配置组，ttl 建议 5 分钟。
func NewCaptchaStore(group string, ttl time.Duration) *CaptchaStore {
	if group == "" {
		group = "default"
	}
	if ttl <= 0 {
		ttl = 5 * time.Minute
	}
	return &CaptchaStore{
		group: group,
		ttl:   ttl,
		ctxFun: func() context.Context {
			return context.Background()
		},
	}
}

var _ base64Captcha.Store = (*CaptchaStore)(nil)

func (s *CaptchaStore) key(id string) string {
	return captchaKeyPrefix + id
}

func (s *CaptchaStore) Set(id string, value string) error {
	ctx := s.ctxFun()
	return g.Redis(s.group).SetEX(ctx, s.key(id), value, int64(s.ttl.Seconds()))
}

func (s *CaptchaStore) Get(id string, clear bool) string {
	ctx := s.ctxFun()
	v, err := g.Redis(s.group).Get(ctx, s.key(id))
	if err != nil || v.IsNil() || v.IsEmpty() {
		return ""
	}
	if clear {
		_, _ = g.Redis(s.group).Del(ctx, s.key(id))
	}
	return v.String()
}

func (s *CaptchaStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v != "" && v == answer
}
