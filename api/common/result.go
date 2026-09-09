package common

import (
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/liuzhengtao/auth-common-backend/internal/consts"
)

type Result struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResultSuccess(data interface{}) *Result {
	return &Result{
		Code: consts.SUCCESS,
		Msg:  consts.GetResultMessage(consts.SUCCESS),
		Data: data,
	}
}

func ResultFailed(code string) *Result {
	msg := consts.GetResultMessage(code)
	rcode := code
	if gstr.Equal(msg, "") {
		rcode = "-1"
		msg = code
	}
	return &Result{
		Code: rcode,
		Msg:  msg,
	}
}

func ResultSystemFailed() *Result {
	return &Result{
		Code: consts.SYSTEM_EXECUTION_ERROR,
		Msg:  consts.GetResultMessage(consts.SYSTEM_EXECUTION_ERROR),
	}
}
func ResultJudge(status bool) *Result {
	if status {
		return ResultSuccess(nil)
	} else {
		return ResultSystemFailed()
	}
}
