package gox

import (
	`fmt`
	`strings`
)

// Paging 分页对象
type Paging struct {
	// 当前页
	Page int `default:"1" json:"page" validate:"min=1"`
	// 每页个数
	Size int `default:"20" json:"size" validate:"min=1"`
	// 查询关键字
	Keyword string `json:"keyword"`
	// 排序顺序
	Order string `default:"DESC" json:"order" validate:"oneof=asc ASC ascending ASCENDING desc DESC descending DESCENDING"`
}


// OrderBy 排序字符串
func (p *Paging) OrderBy(sorter Sorter) string {
	order := "ASC"
	if strings.HasPrefix(strings.ToLower(p.Order), "desc") {
		order = "DESC"
	}

	return fmt.Sprintf("`%s` %s", sorter.SortField(), order)
}

// MySQL 获得MySQL需要的分页参数
func (p *Paging) MySQL() (start int, offset int) {
	return p.Size, (p.Page - 1) * p.Size
}

// Start 获得开始下标
func (p *Paging) Start() int {
	return (p.Page - 1) * p.Size
}

// Limit 获得限制个数
func (p *Paging) Limit() int {
	return p.Size
}
