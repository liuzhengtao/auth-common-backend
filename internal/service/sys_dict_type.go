// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/liuzhengtao/auth-common-backend/api/v1/dict"
	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
)

type (
	ISysDictTypeService interface {
		TypeListPage(ctx context.Context, query dict.TypePageQuery) (output model.ListDictTypePageOutput, err error)
		FindOne(ctx context.Context, id int) (info *entity.SysDictType, err error)
		Add(ctx context.Context, form dict.TypePageVo) error
		Update(ctx context.Context, form dict.TypePageVo) error
		RemoveById(ctx context.Context, Ids []string) error
		Delete(ctx context.Context, ids string) error
	}
)

var (
	localSysDictTypeService ISysDictTypeService
)

func SysDictTypeService() ISysDictTypeService {
	if localSysDictTypeService == nil {
		panic("implement not found for interface ISysDictTypeService, forgot register?")
	}
	return localSysDictTypeService
}

func RegisterSysDictTypeService(i ISysDictTypeService) {
	localSysDictTypeService = i
}
