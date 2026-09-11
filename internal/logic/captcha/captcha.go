package captcha

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/mojocn/base64Captcha"

	"github.com/liuzhengtao/auth-common-backend/api/v1/auth"
	"github.com/liuzhengtao/auth-common-backend/internal/applog"
	"github.com/liuzhengtao/auth-common-backend/internal/consts"
	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

type sCaptchaService struct {
	captcha *base64Captcha.Captcha
	cache   *gcache.Cache
}

func init() {
	service.RegisterCaptchaService(NewCaptchaService())
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
	captchaStore := base64Captcha.DefaultMemStore
	driver := newDriver()
	captcha := base64Captcha.NewCaptcha(driver, captchaStore)
	//cache := lib.RegisterCache(gctx.GetInitCtx(), "redis")
	return &sCaptchaService{captcha: captcha}
}

func (s *sCaptchaService) NewAndStore(ctx context.Context, captchaStoreKey string) error {
	request := g.RequestFromCtx(ctx)
	_, content, answer := s.captcha.Driver.GenerateIdQuestionAnswer()
	item, err := s.captcha.Driver.DrawCaptcha(content)
	if err != nil {
		return err
	}
	err = s.captcha.Store.Set(captchaStoreKey, answer)
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
	_, content, answer := s.captcha.Driver.GenerateIdQuestionAnswer()
	item, err := s.captcha.Driver.DrawCaptcha(content)
	if err != nil {
		applog.Get().Error(ctx, "DrawCaptcha生成错误", err)
		return nil, gerror.New(consts.SYSTEM_EXECUTION_ERROR)
	}
	captchaStoreKey := guid.S()
	err = s.captcha.Store.Set(captchaStoreKey, answer)
	if err != nil {
		applog.Get().Error(ctx, "保存captchaStore报错", err)
		return nil, gerror.New(consts.SYSTEM_EXECUTION_ERROR)
	}
	return &model.CaptchaResult{
		CaptchaRes: &auth.CaptchaRes{CaptchaId: captchaStoreKey, CaptchaBase64: item.EncodeB64string()},
	}, nil
}
func (s *sCaptchaService) Store(ctx context.Context, captchaStoreKey, captchaStoreVal string) error {
	return s.captcha.Store.Set(captchaStoreKey, captchaStoreVal)
}

func (s *sCaptchaService) VerifyAndClear(r *ghttp.Request, captchaStoreKey string, value string) bool {
	return s.captcha.Verify(captchaStoreKey, value, true)
}
