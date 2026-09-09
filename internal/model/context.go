package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
)

type Context struct {
	Session *ghttp.Session  // 当前Session管理对象
	User    *entity.SysUser // 上下文用户信息
	Data    g.Map           // 自定KV变量，业务模块根据需要设置，不固定
}

// SessionNotice 存放在Session中的提示信息，往往使用后则删除
type SessionNotice struct {
	Type    string // 消息类型
	Content string // 消息内容
}
