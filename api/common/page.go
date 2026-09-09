package common

import "github.com/liuzhengtao/auth-common-backend/internal/consts"

type Page struct {
	Records  []interface{} `json:"records"`
	Total    int           `json:"total"`
	Size     int           `json:"size"`
	Current  int           `json:"current"`
	MaxLimit int           `json:"maxLimit"`
}

func NewPageWithCurrentSize(current, size int) *Page {
	return &Page{
		Current: current,
		Size:    size,
	}
}

func NewPageTotal(current, size, total int) *Page {
	return &Page{
		Current: current,
		Size:    size,
		Total:   total,
	}
}

func (p *Page) HasNext() bool {
	return p.Current < p.Total/p.Size
}

func (p *Page) GetPages() int {
	if p.Size == 0 {
		return 0
	}
	pages := p.Total / p.Size
	if p.Total%p.Size != 0 {
		pages++
	}
	return pages
}

type PageResult struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}

type Data struct {
	List  []interface{} `json:"list"`
	Total int64         `json:"total"`
}

func Success(page Page) *PageResult {
	pageResult := &PageResult{
		Data: Data{
			List:  page.Records,
			Total: int64(page.Total),
		},
	}
	pageResult.Code = consts.SUCCESS
	pageResult.Msg = consts.GetResultMessage(consts.SUCCESS)
	return pageResult
}

type BasePageQuery struct {
	PageNum  int `p:"pageNum" v:"required#页码不能为空" dc:"页码" d:"1"`
	PageSize int `p:"pageSize" v:"required#每页记录数不能为空" dc:"每页记录数" d:"10"`
}

type Res struct{}
