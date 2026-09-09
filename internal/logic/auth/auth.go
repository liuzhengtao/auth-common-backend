package auth

import (
	"context"
	"time"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"

	lib "gitee.com/zhengtao313/lib/utility"

	"github.com/liuzhengtao/auth-common-backend/internal/config"
	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

type sAuth struct {
	jwtMiddleWare *jwt.GfJWTMiddleware
	crypto        *lib.Crypt
	logger        *glog.Logger
}

func init() {
	service.RegisterAuth(New())
}
func New() *sAuth {
	cfg := config.Get()
	rAuth := sAuth{}
	jwtMid := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "auth-common",
		Key:             []byte(cfg.JwtSecret),
		Timeout:         time.Minute * 120,
		MaxRefresh:      time.Hour * 24,
		IdentityKey:     "info",
		TokenLookup:     "header:Authorization,query:token,cookie:jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Unauthorized:    rAuth.Unauthorized,
		PayloadFunc:     rAuth.PayLoadFunc,
		IdentityHandler: rAuth.IdentityHandler,
	})
	crypto := &lib.Crypt{
		AesKey: cfg.JwtAesKey,
	}
	rAuth.jwtMiddleWare = jwtMid
	rAuth.crypto = crypto
	rAuth.logger = g.Log().Line(true)
	return &rAuth
}

func (s *sAuth) MiddlewareFunc() ghttp.HandlerFunc {
	return s.jwtMiddleWare.MiddlewareFunc()
}

func (s *sAuth) GetIdentity(ctx context.Context) (info *model.UserAuthInfo) {
	identityVar := s.jwtMiddleWare.GetIdentity(ctx)
	err := gconv.Scan(identityVar, &info)
	if err != nil {
		s.logger.Error(ctx, "GetIdentity转化为UserAuthInfo失败", err)
		return nil
	}
	return
}

func (s *sAuth) PayLoadFunc(data any) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]any)
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

func (s *sAuth) IdentityHandler(ctx context.Context) any {
	claims := jwt.ExtractClaims(ctx)
	decrypt, err := s.crypto.AesDecrpt(ctx, gconv.String(claims[s.jwtMiddleWare.IdentityKey]))
	if err != nil {
		s.logger.Error(ctx, "aes解密失败", err)
		return nil
	}
	return decrypt
}

// Unauthorized 登录验证失败，处理不进行授权的逻辑
func (s *sAuth) Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	curRouter := r.Router.Uri
	if curRouter == "" {
		r.Middleware.Next()
	} else {
		r.Response.WriteJsonExit(g.Map{
			"code": code,
			"msg":  message,
		})
	}
}

func (s *sAuth) GetPayload(ctx context.Context) string {
	return s.jwtMiddleWare.GetPayload(ctx)
}

func (s *sAuth) GeneratorToken(ctx context.Context, data model.UserAuthInfo) (token string, expire time.Time, err error) {
	genData := gmap.NewStrAnyMap()
	aesStr, err := s.crypto.AesEncrypt(ctx, gjson.MustEncodeString(data))
	if err != nil {
		return "", time.Time{}, err
	}
	genData.Set("info", aesStr)
	genData.Set("nickname", data.Nickname)
	return s.jwtMiddleWare.TokenGenerator(genData.Map())
}

func (s *sAuth) RefreshHandler(ctx context.Context) (token string, expire time.Time) {
	token, expire = s.jwtMiddleWare.RefreshHandler(ctx)
	return
}
func (s *sAuth) LogoutHandler(ctx context.Context) {
	s.jwtMiddleWare.LogoutHandler(ctx)
	return
}

func (s *sAuth) GetUser(ctx context.Context) (*entity.SysUser, error) {
	authUserInfo := s.GetIdentity(ctx)
	if authUserInfo == nil {
		return nil, gerror.New("获取token信息失败")
	}
	user, err := service.SysUserService().GetUser(ctx, int64(authUserInfo.UserId))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *sAuth) GetRoles(ctx context.Context) (roles []string, err error) {
	authUserInfo := s.GetIdentity(ctx)
	if authUserInfo == nil {
		return nil, gerror.New("获取token信息失败")
	}
	return authUserInfo.Roles, nil
}
