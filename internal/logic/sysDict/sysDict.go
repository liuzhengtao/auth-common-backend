package sysDict

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/liuzhengtao/auth-common-backend/api/v1/dept"
	"github.com/liuzhengtao/auth-common-backend/api/v1/dict"
	"github.com/liuzhengtao/auth-common-backend/internal/dao"
	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

type sSysDictService struct{}

func New() *sSysDictService {
	return &sSysDictService{}
}

func init() {
	service.RegisterSysDictService(New())
}
func (s *sSysDictService) ListPage(ctx context.Context, query dict.PageQuery) (output model.ListDictPageOutput, err error) {
	columns := dao.SysDict.Columns()
	orm := dao.SysDict.Ctx(ctx)
	if !gstr.Equal(query.Keyword, "") {
		orm = orm.WhereLike("name", "%"+query.Keyword+"%")
	}
	if !gstr.Equal(query.TypeCode, "") {
		orm = orm.Where("type_code", query.TypeCode)
	}
	result, totalCount, err := orm.Page(query.PageNum, query.PageSize).Order("sort asc").Fields(
		columns.Id,
		columns.Name,
		columns.Value,
		columns.Status,
	).AllAndCount(false)
	output = model.ListDictPageOutput{
		Total: gconv.Int64(totalCount),
	}
	err = result.Structs(&output.List)
	return
}

func (s *sSysDictService) ListDictOptions(ctx context.Context, typeCode string) (dictList []dept.Option, err error) {
	var sysDicts []entity.SysDict
	err = dao.SysDict.Ctx(ctx).Fields("name,value").Where("type_code", typeCode).Scan(&sysDicts)
	if err != nil {
		return nil, err
	}
	dictList = make([]dept.Option, 0)
	for _, dict := range sysDicts {
		dictList = append(dictList, dept.Option{
			Value: dict.Value,
			Label: dict.Name,
		})
	}
	return
}

func (s *sSysDictService) UpdateTypeCode(ctx context.Context, oldCode, newCode string) error {
	_, err := dao.SysDict.Ctx(ctx).Where("type_code", oldCode).Update(g.Map{
		"type_code": newCode,
	})
	return err
}

func (s *sSysDictService) FindOne(ctx context.Context, id int) (info *entity.SysDict, err error) {
	err = dao.SysDict.Ctx(ctx).Where(dao.SysDict.Columns().Id, id).Scan(&info)
	return
}

func (s *sSysDictService) Remove(ctx context.Context, typeCodes []string) error {
	_, err := dao.SysDict.Ctx(ctx).WhereIn(dao.SysDict.Columns().TypeCode, typeCodes).Delete()
	return err
}

func (s *sSysDictService) RemoveIds(ctx context.Context, Ids []string) error {
	_, err := dao.SysDict.Ctx(ctx).WhereIn(dao.SysDict.Columns().Id, Ids).Delete()
	return err
}

func (s *sSysDictService) Add(ctx context.Context, form dict.PageVo) error {
	_, err := dao.SysDict.Ctx(ctx).Data(form).Insert()
	return err
}
func (s *sSysDictService) Update(ctx context.Context, form dict.PageVo) error {
	one, err := s.FindOne(ctx, form.Id)
	if err != nil {
		return err
	}
	if one == nil {
		return gerror.New("字典不存在")
	}
	_, err = dao.SysDict.Ctx(ctx).Data(form).Save()
	return nil
}

func (s *sSysDictService) Delete(ctx context.Context, ids string) error {
	if len(gstr.Split(ids, ",")) == 0 {
		return gerror.New("删除数据为空")
	}
	return s.RemoveIds(ctx, gstr.Split(ids, ","))
}
