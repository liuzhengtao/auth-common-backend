package main

import (
	"context"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/liuzhengtao/auth-common-backend/internal/cmd"
)

func main() {
	cmd.Main.Run(context.Background())
}
