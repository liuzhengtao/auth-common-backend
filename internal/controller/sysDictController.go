package controller

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/liuzhengtao/auth-common-backend/api/common"
	"github.com/liuzhengtao/auth-common-backend/api/v1/dict"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

type cSysDictController struct {
	dict     service.ISysDictService
	dictType service.ISysDictTypeService
}

func NewSysDictController() *cSysDictController {
	return &cSysDictController{
		dict:     service.SysDictService(),
		dictType: service.SysDictTypeService(),
	}
}
func (c *cSysDictController) ListDictOptions(ctx context.Context, req *dict.ListDictOptionsReq) (resp dict.ListDictOptionsRes, err error) {
	return c.dict.ListDictOptions(ctx, req.TypeCode)
}

func (c *cSysDictController) ListPage(ctx context.Context, req *dict.ListReq) (resp *dict.ListRes, err error) {
	output, err := c.dict.ListPage(ctx, req.PageQuery)
	return (*dict.ListRes)(&output), err
}

func (c *cSysDictController) TypeListPage(ctx context.Context, req *dict.ListTypeReq) (resp *dict.ListTypeRes, err error) {
	output, err := c.dictType.TypeListPage(ctx, req.TypePageQuery)
	return (*dict.ListTypeRes)(&output), err
}

func (c *cSysDictController) AddType(ctx context.Context, req *dict.AddTypeFormReq) (resp *common.Res, err error) {
	err = c.dictType.Add(ctx, req.TypePageVo)
	return
}
func (c *cSysDictController) UpdateType(ctx context.Context, req *dict.UpdateTypeFormReq) (resp *common.Res, err error) {
	req.TypePageVo.Id = req.TypeId
	err = c.dictType.Update(ctx, req.TypePageVo)
	return
}
func (c *cSysDictController) DeleteType(ctx context.Context, req *dict.DeleteTypeFormReq) (resp *common.Res, err error) {
	err = c.dictType.Delete(ctx, req.Ids)
	return
}

func (c *cSysDictController) AddDict(ctx context.Context, req *dict.AddFormReq) (resp *common.Res, err error) {
	err = c.dict.Add(ctx, req.PageVo)
	return
}
func (c *cSysDictController) UpdateDict(ctx context.Context, req *dict.UpdateFormReq) (resp *common.Res, err error) {
	req.PageVo.Id = req.Id
	err = c.dict.Update(ctx, req.PageVo)
	return
}
func (c *cSysDictController) DeleteDict(ctx context.Context, req *dict.DeleteFormReq) (resp *common.Res, err error) {
	err = c.dict.Delete(ctx, req.Ids)
	return
}

func (c *cSysDictController) GetTypeForm(ctx context.Context, req *dict.GetTypeFormReq) (resp *dict.GetTypeFormRes, err error) {
	info, err := c.dictType.FindOne(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(info, &resp)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cSysDictController) GetDictForm(ctx context.Context, req *dict.GetDictFormReq) (resp *dict.GetDictFormRes, err error) {
	info, err := c.dict.FindOne(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(info, &resp)
	if err != nil {
		return nil, err
	}
	return
}
