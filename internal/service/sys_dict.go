// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/liuzhengtao/auth-common-backend/api/v1/dept"
	"github.com/liuzhengtao/auth-common-backend/api/v1/dict"
	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
)

type (
	ISysDictService interface {
		ListPage(ctx context.Context, query dict.PageQuery) (output model.ListDictPageOutput, err error)
		ListDictOptions(ctx context.Context, typeCode string) (dictList []dept.Option, err error)
		UpdateTypeCode(ctx context.Context, oldCode, newCode string) error
		FindOne(ctx context.Context, id int) (info *entity.SysDict, err error)
		Remove(ctx context.Context, typeCodes []string) error
		RemoveIds(ctx context.Context, Ids []string) error
		Add(ctx context.Context, form dict.PageVo) error
		Update(ctx context.Context, form dict.PageVo) error
		Delete(ctx context.Context, ids string) error
	}
)

var (
	localSysDictService ISysDictService
)

func SysDictService() ISysDictService {
	if localSysDictService == nil {
		panic("implement not found for interface ISysDictService, forgot register?")
	}
	return localSysDictService
}

func RegisterSysDictService(i ISysDictService) {
	localSysDictService = i
}
