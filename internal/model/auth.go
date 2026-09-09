package model

import "github.com/liuzhengtao/auth-common-backend/api/v1/auth"

type LoginInput struct {
	Username string `p:"username" v:"required#用户名不能为空" json:"username,omitempty"`
	Password string `p:"password" v:"required#密码不能为空" json:"password,omitempty"`
}
type LoginOutput struct {
	LoginRes *auth.LoginRes
}
