package captcha

import (
	"context"
	"sync"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/mojocn/base64Captcha"

	"github.com/liuzhengtao/auth-common-backend/api/v1/auth"
	"github.com/liuzhengtao/auth-common-backend/internal/applog"
	"github.com/liuzhengtao/auth-common-backend/internal/consts"
	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/redisstore"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

type sCaptchaService struct {
	mu      sync.RWMutex
	captcha *base64Captcha.Captcha
}

var localCaptcha *sCaptchaService

func init() {
	localCaptcha = NewCaptchaService()
	service.RegisterCaptchaService(localCaptcha)
}

func newDriver() *base64Captcha.DriverString {
	driver := &base64Captcha.DriverString{
		Height:          44,
		Width:           126,
		NoiseCount:      5,
		ShowLineOptions: base64Captcha.OptionShowSineLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowHollowLine,
		Length:          4,
		Source:          "1234567890",
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	return driver.ConvertFonts()
}

func NewCaptchaService() *sCaptchaService {
	store := base64Captcha.DefaultMemStore
	driver := newDriver()
	return &sCaptchaService{captcha: base64Captcha.NewCaptcha(driver, store)}
}

// UseRedisStore 切换为 Redis 验证码存储（分布式模式由 Install 调用）。
func UseRedisStore(group string) {
	if localCaptcha == nil {
		return
	}
	store := redisstore.NewCaptchaStore(group, 5*time.Minute)
	localCaptcha.mu.Lock()
	defer localCaptcha.mu.Unlock()
	localCaptcha.captcha = base64Captcha.NewCaptcha(newDriver(), store)
}

func (s *sCaptchaService) getCaptcha() *base64Captcha.Captcha {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.captcha
}

func (s *sCaptchaService) NewAndStore(ctx context.Context, captchaStoreKey string) error {
	request := g.RequestFromCtx(ctx)
	c := s.getCaptcha()
	_, content, answer := c.Driver.GenerateIdQuestionAnswer()
	item, err := c.Driver.DrawCaptcha(content)
	if err != nil {
		return err
	}
	err = c.Store.Set(captchaStoreKey, answer)
	if err != nil {
		return err
	}
	_, err = item.WriteTo(request.Response.Writer)
	if err != nil {
		return err
	}
	return nil
}

func (s *sCaptchaService) GetCaptcha(ctx context.Context) (out *model.CaptchaResult, err error) {
	c := s.getCaptcha()
	_, content, answer := c.Driver.GenerateIdQuestionAnswer()
	item, err := c.Driver.DrawCaptcha(content)
	if err != nil {
		applog.Get().Error(ctx, "DrawCaptcha生成错误", err)
		return nil, gerror.New(consts.SYSTEM_EXECUTION_ERROR)
	}
	captchaStoreKey := guid.S()
	err = c.Store.Set(captchaStoreKey, answer)
	if err != nil {
		applog.Get().Error(ctx, "保存captchaStore报错", err)
		return nil, gerror.New(consts.SYSTEM_EXECUTION_ERROR)
	}
	return &model.CaptchaResult{
		CaptchaRes: &auth.CaptchaRes{CaptchaId: captchaStoreKey, CaptchaBase64: item.EncodeB64string()},
	}, nil
}

func (s *sCaptchaService) Store(ctx context.Context, captchaStoreKey, captchaStoreVal string) error {
	return s.getCaptcha().Store.Set(captchaStoreKey, captchaStoreVal)
}

func (s *sCaptchaService) VerifyAndClear(r *ghttp.Request, captchaStoreKey string, value string) bool {
	return s.getCaptcha().Verify(captchaStoreKey, value, true)
}
