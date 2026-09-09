package sysDictType

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/liuzhengtao/auth-common-backend/api/v1/dict"
	"github.com/liuzhengtao/auth-common-backend/internal/dao"
	"github.com/liuzhengtao/auth-common-backend/internal/model"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

type sSysDictTypeService struct {
	logger *glog.Logger
}

func New() *sSysDictTypeService {
	return &sSysDictTypeService{
		logger: g.Log().Line(true),
	}
}

func init() {
	service.RegisterSysDictTypeService(New())
}

func (s *sSysDictTypeService) TypeListPage(ctx context.Context, query dict.TypePageQuery) (output model.ListDictTypePageOutput, err error) {
	columns := dao.SysDictType.Columns()
	orm := dao.SysDictType.Ctx(ctx)
	if !gstr.Equal(query.Keyword, "") {
		orm = orm.WhereLike("name", "%"+query.Keyword+"%").WhereOr("code like ? ", "%"+query.Keyword+"%")
	}
	result, totalCount, err := orm.Page(query.PageNum, query.PageSize).Fields(
		columns.Id,
		columns.Name,
		columns.Code,
		columns.Status,
		columns.Remark,
	).AllAndCount(false)
	output = model.ListDictTypePageOutput{
		Total: gconv.Int64(totalCount),
	}
	err = result.Structs(&output.List)
	return
}
func (s *sSysDictTypeService) FindOne(ctx context.Context, id int) (info *entity.SysDictType, err error) {
	err = dao.SysDictType.Ctx(ctx).Where(dao.SysDictType.Columns().Id, id).Scan(&info)
	return
}

func (s *sSysDictTypeService) Add(ctx context.Context, form dict.TypePageVo) error {
	_, err := dao.SysDictType.Ctx(ctx).Data(form).Insert()
	return err
}

func (s *sSysDictTypeService) Update(ctx context.Context, form dict.TypePageVo) error {
	one, err := s.FindOne(ctx, form.Id)
	if err != nil {
		return err
	}
	if one == nil {
		return gerror.New("字典类型不存在")
	}
	result, err := dao.SysDictType.Ctx(ctx).Data(form).Save()
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected > 0 {
		oldCode := one.Code
		newCode := form.Code
		if !gstr.Equal(oldCode, newCode) {
			return service.SysDictService().UpdateTypeCode(ctx, oldCode, newCode)
		}
	}
	return nil
}

func (s *sSysDictTypeService) RemoveById(ctx context.Context, Ids []string) error {
	_, err := dao.SysDictType.Ctx(ctx).WhereIn(dao.SysDictType.Columns().Id, Ids).Delete()
	return err
}

func (s *sSysDictTypeService) Delete(ctx context.Context, ids string) error {
	var dictTypeCodes []string
	result, err := dao.SysDictType.Ctx(ctx).WhereIn("id", gstr.Split(ids, ",")).Fields("code").Array()
	if err != nil {
		return err
	}
	for _, value := range result {
		dictTypeCodes = append(dictTypeCodes, value.String())
	}
	if len(dictTypeCodes) > 0 {
		err = service.SysDictService().Remove(ctx, dictTypeCodes)
		if err != nil {
			return err
		}
	}
	return s.RemoveById(ctx, gstr.Split(ids, ","))
}
