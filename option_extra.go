package gox

var _ pageOption = (*optionExtra)(nil)

type optionExtra struct {
	extra map[string]interface{}
}

// Extra 配置额外信息
func Extra(extra map[string]interface{}) *optionExtra {
	return &optionExtra{
		extra: extra,
	}
}

func (e *optionExtra) applyPage(options *pageOptions) {
	options.extra = e.extra
}
