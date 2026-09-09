package utility

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/liuzhengtao/auth-common-backend/internal/config"
)

func EncryptData(data any) any {
	needEncryptData := gconv.String(data) + config.Get().EncryptSalt
	return gmd5.MustEncryptString(needEncryptData)
}
