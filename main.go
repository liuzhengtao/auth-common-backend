package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/liuzhengtao/auth-common-backend/common"
)

func main() {
	common.Cmd.Run(gctx.GetInitCtx())
}
