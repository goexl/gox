package gox

var _ pageOption = (*optionPage)(nil)

type optionPage struct {
	page int32
}

// Page 配置页数
func Page(page int32) *optionPage {
	return &optionPage{
		page: page,
	}
}

func (p *optionPage) applyPage(options *pageOptions) {
	options.page = p.page
}
