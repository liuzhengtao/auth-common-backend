package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/liuzhengtao/auth-common-backend/install"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server with auth-common plugin",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			webServer := g.Server()
			if err = install.Install(webServer); err != nil {
				return err
			}
			webServer.Run()
			return nil
		},
	}
)
