package auth

import "github.com/gogf/gf/v2/frame/g"

type CaptchaReq struct {
	g.Meta `path:"/captcha" method:"get" tags:"认证中心" dc:"获取验证码"`
}

type LoginReq struct {
	g.Meta      `path:"/login" method:"post" tags:"认证中心" dc:"登录"`
	Username    string `p:"username" v:"required#用户名不能为空" json:"username,omitempty"`
	Password    string `p:"password" v:"required#密码不能为空" json:"password,omitempty"`
	CaptchaCode string `p:"captchaCode" v:"required#验证码code不能为空" json:"captchaCode,omitempty"`
	CaptchaId   string `p:"captchaId" v:"required#验证码id不能为空" json:"captchaId,omitempty"`
}

type LogoutReq struct {
	g.Meta `path:"/auth/logout" method:"delete" tags:"认证中心" dc:"登出"`
}
type CaptchaRes struct {
	CaptchaId     string `json:"captchaId"`
	CaptchaBase64 string `json:"captchaBase64"`
}

type LoginRes struct {
	AccessToken  string `json:"access_token,omitempty" dc:"访问token"`
	TokenType    string `json:"token_type,omitempty" dc:"token 类型 例如 Bearer"`
	RefreshToken string `json:"refresh_token,omitempty" dc:"刷新token"`
	Expires      int64  `json:"expires,omitempty" dc:"token 过期时间(单位：毫秒)"`
}
