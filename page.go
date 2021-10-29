package gox

import (
	`encoding/json`
)

type page struct {
	// 当前页码
	Page int32 `json:"-"`
	// 是否还有下一页数据
	HasNext bool `json:"-"`
	// 是否有上一页数据
	HasPrev bool `json:"-"`
	// 总共数量
	Total int64 `json:"-"`
	// 总共页数
	Pages int64 `json:"-"`
	// 数据列表
	Items interface{} `json:"-"`
	// 额外数据
	Fields []Field `json:"-"`
}

// NewPage 生成新的分页数据对象
func NewPage(items interface{}, total int64, opts ...pageOption) *page {
	options := defaultPageOptions()
	for _, opt := range opts {
		opt.applyPage(options)
	}

	pages := total / int64(options.size)
	if (total % int64(options.size)) > 0 {
		pages += 1
	}

	hasPrev := false
	if options.page > 1 {
		hasPrev = true
	}

	hasNext := false
	if int64(options.page) < pages {
		hasNext = true
	}

	return &page{
		Page:    options.page,
		HasNext: hasNext,
		HasPrev: hasPrev,
		Total:   total,
		Pages:   pages,
		Items:   items,
		Fields:  options.fields,
	}
}

func (p page) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{}, 7)
	data[`page`] = p.Page
	data[`hasNext`] = p.HasNext
	data[`hasPrev`] = p.HasPrev
	data[`total`] = p.Total
	data[`items`] = p.Items
	data[`pages`] = p.Pages
	for _, field := range p.Fields {
		data[field.Key()] = field.Value()
	}

	return json.Marshal(data)
}
