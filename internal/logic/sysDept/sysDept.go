package sysDept

import (
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/liuzhengtao/auth-common-backend/api/v1/dept"
	"github.com/liuzhengtao/auth-common-backend/internal/consts"
	"github.com/liuzhengtao/auth-common-backend/internal/dao"
	"github.com/liuzhengtao/auth-common-backend/internal/model/do"
	"github.com/liuzhengtao/auth-common-backend/internal/model/entity"
	"github.com/liuzhengtao/auth-common-backend/internal/service"
)

type sSysDeptService struct{}

func New() *sSysDeptService {
	return &sSysDeptService{}
}

func init() {
	service.RegisterSysDeptService(New())
}

func recurDeptTreeOptions(parentId int, deptList []*entity.SysDept) (list []dept.Option, err error) {
	list = make([]dept.Option, 0)
	for _, sysDept := range deptList {
		if gconv.Int(sysDept.ParentId) == parentId {
			option := dept.Option{
				Value: gconv.Int(sysDept.Id),
				Label: sysDept.Name,
			}
			children, err := recurDeptTreeOptions(gconv.Int(sysDept.Id), deptList)
			if err != nil {
				return nil, err
			}
			option.Children = children
			list = append(list, option)
		}
	}
	return
}

func (s *sSysDeptService) DeptOptionsList(ctx context.Context) (list []dept.Option, err error) {
	deptList, err := dao.SysDept.DeptOptionsList(ctx)
	if err != nil {
		return nil, err
	}
	deptIdsSet := gset.NewIntSet()
	parentIdsSet := gset.NewIntSet()
	for _, deptItem := range deptList {
		deptIdsSet.Add(gconv.Int(deptItem.Id))
		parentIdsSet.Add(gconv.Int(deptItem.ParentId))
	}
	rootIds := parentIdsSet.Diff(deptIdsSet).Slice()
	list = make([]dept.Option, 0)
	for _, rootId := range rootIds {
		rootIdList, err := recurDeptTreeOptions(rootId, deptList)
		if err != nil {
			return nil, err
		}
		list = append(list, rootIdList...)
	}
	return
}

func (s *sSysDeptService) ListDepartments(ctx context.Context, keywords string, status int) (list []dept.DeptVo, err error) {
	orm := dao.SysDept.Ctx(ctx).OrderAsc(dao.SysDept.Columns().Sort)
	if !gstr.Equal(keywords, "") {
		orm = orm.Where(dao.SysDept.Columns().Name+" like ?", "%"+keywords+"%")
	}
	if status != 4 {
		orm = orm.Where(dao.SysDept.Columns().Status, status)
	}
	var depts []entity.SysDept
	err = orm.Scan(&depts)
	if err != nil {
		return nil, err
	}
	deptIds := gset.NewIntSet()
	parentIds := gset.NewIntSet()
	for _, dt := range depts {
		deptIds.Add(gconv.Int(dt.Id))
		parentIds.Add(gconv.Int(dt.ParentId))
	}
	list = make([]dept.DeptVo, 0)
	for _, rootId := range parentIds.Diff(deptIds).Slice() {
		deptTrees, err := s.recurDeptTree(ctx, rootId, depts)
		if err != nil {
			return nil, err
		}
		list = append(list, deptTrees...)
	}
	return
}

func (s *sSysDeptService) recurDeptTree(ctx context.Context, parentId int, deptList []entity.SysDept) (list []dept.DeptVo, err error) {
	list = make([]dept.DeptVo, 0)
	for _, dt := range deptList {
		if gconv.Int(dt.ParentId) == parentId {
			dv := dept.DeptVo{
				SysDept: dt,
			}
			children, err := s.recurDeptTree(ctx, gconv.Int(dt.Id), deptList)
			if err != nil {
				return nil, err
			}
			dv.Children = children
			list = append(list, dv)
		}
	}
	return
}

func (s *sSysDeptService) Add(ctx context.Context, formData dept.DepartmentForm) error {
	count, err := dao.SysDept.Ctx(ctx).Where(dao.SysDept.Columns().Name, formData.Name).Where(dao.SysDept.Columns().Status, 1).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("该部门已存在")
	}
	dp := do.SysDept{
		Name:     formData.Name,
		ParentId: formData.ParentId,
		Sort:     formData.Sort,
		Status:   formData.Status,
	}
	//生成部门路径(tree_path)，格式：父节点tree_path + , + 父节点ID，用于删除部门时级联删除子部门
	treePath, err := s.generateDeptTreePath(ctx, formData.ParentId)
	if err != nil {
		return err
	}
	dp.TreePath = treePath
	_, err = dao.SysDept.Ctx(ctx).Insert(dp)
	if err != nil {
		return err
	}
	return nil
}

func (s *sSysDeptService) Update(ctx context.Context, formData dept.DepartmentForm) error {
	count, err := dao.SysDept.Ctx(ctx).Where(dao.SysDept.Columns().Name, formData.Name).WhereNot(dao.SysDept.Columns().Id, formData.Id).Where(dao.SysDept.Columns().Status, 1).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("该部门已存在")
	}
	dp := do.SysDept{
		Name:     formData.Name,
		ParentId: formData.ParentId,
		Sort:     formData.Sort,
		Status:   formData.Status,
	}
	dp.Id = formData.Id
	//生成部门路径(tree_path)，格式：父节点tree_path + , + 父节点ID，用于删除部门时级联删除子部门
	treePath, err := s.generateDeptTreePath(ctx, formData.ParentId)
	if err != nil {
		return err
	}
	dp.TreePath = treePath
	_, err = dao.SysDept.Ctx(ctx).Data(dp).Save()
	if err != nil {
		return err
	}
	return nil
}

func (s *sSysDeptService) Delete(ctx context.Context, ids string) error {
	idsArr := gstr.Split(ids, ",")
	for _, id := range idsArr {
		_, err := dao.SysDept.Ctx(ctx).Where(dao.SysDept.Columns().Id, id).WhereOr("CONCAT(',',tree_path,',') like CONCAT('%,',?,',%')", id).Delete()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *sSysDeptService) FindOne(ctx context.Context, id int) (dept *entity.SysDept, err error) {
	err = dao.SysDept.Ctx(ctx).Where(dao.SysDept.Columns().Id, id).Scan(&dept)
	return
}
func (s *sSysDeptService) GetForm(ctx context.Context, id int) (dept *dept.DepartmentForm, err error) {
	columns := dao.SysDept.Columns()
	err = dao.SysDept.Ctx(ctx).Where(columns.Id, id).Fields(
		columns.Id,
		columns.Name,
		columns.ParentId,
		columns.Sort,
		columns.Status,
	).Scan(&dept)
	return
}

func (s *sSysDeptService) generateDeptTreePath(ctx context.Context, parentId int) (treePath string, err error) {
	if parentId == consts.ROOT_NODE_ID {
		return gconv.String(consts.ROOT_NODE_ID), nil
	} else {
		parentInfo, err := s.FindOne(ctx, parentId)
		if err != nil {
			return "", err
		}
		if parentInfo != nil {
			return parentInfo.TreePath + "," + gconv.String(parentInfo.Id), nil
		}
	}
	return "", gerror.New("部门路径生成失败")
}
