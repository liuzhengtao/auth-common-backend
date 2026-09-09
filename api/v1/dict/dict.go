package dict

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/liuzhengtao/auth-common-backend/api/common"
	"github.com/liuzhengtao/auth-common-backend/api/v1/dept"
)

type ListDictOptionsReq struct {
	g.Meta   `path:"/types/{typeCode}/options" method:"get" tags:"字典接口" dc:"用户表单数据"`
	TypeCode string `p:"typeCode" in:"path" v:"required#字典类型编码不能为空" dc:"字典类型编码"`
}

type PageQuery struct {
	common.BasePageQuery
	Keyword  string `p:"keywords"  dc:"查询关键字"`
	TypeCode string `p:"typeCode"  dc:"字典类型编码"`
}

type TypePageQuery struct {
	common.BasePageQuery
	Keyword string `p:"keywords"  dc:"查询关键字"`
}

type ListReq struct {
	g.Meta `path:"/page" method:"get" tags:"字典接口" dc:"字典分页列表"`
	PageQuery
}

type ListTypeReq struct {
	g.Meta `path:"/types/page" method:"get" tags:"字典接口" dc:"字典类型分页列表"`
	TypePageQuery
}

type GetTypeFormReq struct {
	g.Meta `path:"/types/{id}/form" method:"get" tags:"字典接口" dc:"字典类型表单数据"`
	Id     int `p:"id" in:"path" v:"required#字典类型ID不能为空"`
}
type GetDictFormReq struct {
	g.Meta `path:"/{id}/form" method:"get" tags:"字典接口" dc:"字典数据表单数据"`
	Id     int `p:"id" in:"path" v:"required#字典ID不能为空"`
}

type GetTypeFormRes TypePageVo
type GetDictFormRes PageVo
type AddTypeFormReq struct {
	g.Meta `path:"/types" method:"post" tags:"字典接口" dc:"新增字典类型"`
	TypePageVo
}

type UpdateTypeFormReq struct {
	g.Meta `path:"/types/{typeId}" method:"put" tags:"字典接口" dc:"修改字典类型"`
	TypeId int `p:"typeId" in:"path" v:"required#字典类型ID不能为空"`
	TypePageVo
}
type DeleteTypeFormReq struct {
	g.Meta `path:"/types/{ids}" method:"delete" tags:"字典接口" dc:"删除字典类型"`
	Ids    string `p:"ids" in:"path" v:"required#字典类型ID字符串不能为空" dc:"字典类型ID，多个以英文逗号(,)分割"`
}

type AddFormReq struct {
	g.Meta `path:"/" method:"post" tags:"字典接口" dc:"新增字典"`
	PageVo
}

type UpdateFormReq struct {
	g.Meta `path:"/{id}" method:"put" tags:"字典接口" dc:"修改字典"`
	Id     int `p:"id" in:"path" v:"required#字典类型ID不能为空"`
	PageVo
}
type DeleteFormReq struct {
	g.Meta `path:"/{ids}" method:"delete" tags:"字典接口" dc:"删除字典"`
	Ids    string `p:"ids" in:"path" v:"required#字典类型ID字符串不能为空" dc:"字典类型ID，多个以英文逗号(,)分割"`
}

type PageVo struct {
	Id        int      `json:"id" dc:"字典ID"`
	Name      string   `json:"name" dc:"字典名称"`
	Code      string   `json:"typeCode" dc:"字典编码"`
	Value     string   `json:"value" dc:"字典值"`
	Status    int      `json:"status" dc:"状态(1:启用;0:禁用)"`
	DictItems []PageVo `json:"dictItems" dc:"字典项列表"`
}

type TypePageVo struct {
	Id     int    `json:"id" dc:"字典类型ID"`
	Name   string `json:"name" dc:"类型名称"`
	Code   string `json:"code" dc:"类型编码"`
	Status int    `json:"status" dc:"状态(1:启用;0:禁用)"`
	Remark string `json:"remark" dc:"备注"`
}

type ListRes struct {
	List  []PageVo `json:"list" dc:"字典列表"`
	Total int64    `json:"total" dc:"总记录数"`
}

type ListTypeRes struct {
	List  []TypePageVo `json:"list" dc:"字典类型列表"`
	Total int64        `json:"total" dc:"总记录数"`
}

type ListDictOptionsRes []dept.Option
