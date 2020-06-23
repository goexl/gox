package gox

import (
	"fmt"
	"strings"
)

// Paging 分页对象
type Paging struct {
	// 当前页
	Page int `default:"1" json:"page" validate:"min=1"`
	// 每页个数
	PerPage int `default:"20" json:"perPage" validate:"min=1"`
	// 查询关键字
	Keyword string `json:"keyword"`
	// 排序顺序
	SortOrder string `default:"DESC" json:"sortOrder" validate:"oneof=asc ASC ascending ASCENDING desc DESC descending DESCENDING"`
}

// Sorter 排序
type Sorter interface {
	SortFieldName() string
}

// 排序字符串
func (p *Paging) OrderBy(sorter Sorter) string {
	sortOrder := "ASC"
	if strings.HasPrefix(strings.ToLower(p.SortOrder), "desc") {
		sortOrder = "DESC"
	}

	return fmt.Sprintf("%s %s", sorter.SortFieldName(), sortOrder)
}

// MySQL 获得MySQL需要的分页参数
func (p *Paging) MySQL() (start int, offset int) {
	return p.PerPage, (p.Page - 1) * p.PerPage
}

// Start 获得开始下标
func (p *Paging) Start() int {
	return (p.Page - 1) * p.PerPage
}

// Limit 获得限制个数
func (p *Paging) Limit() int {
	return p.PerPage
}

type pageData struct {
	CurrentPage int         `json:"currentPage"`
	HasNext     bool        `json:"hasNext"`
	HasPrev     bool        `json:"hasPrev"`
	TotalNum    int64       `json:"totalNum"`
	TotalPage   int64       `json:"totalPage"`
	Items       interface{} `json:"items"`
}

// NewPage 生成新的分页数据对象
func NewPage(items interface{}, totalNum int64, perPage int, page int) *pageData {
	totalPage := totalNum / int64(perPage)
	if (totalNum % int64(perPage)) > 0 {
		totalPage += 1
	}

	hasPrev := false
	if page > 1 {
		hasPrev = true
	}

	hasNext := false
	if int64(page) < totalPage {
		hasNext = true
	}

	return &pageData{
		CurrentPage: page,
		HasNext:     hasNext,
		HasPrev:     hasPrev,
		TotalNum:    totalNum,
		TotalPage:   totalPage,
		Items:       items,
	}
}
