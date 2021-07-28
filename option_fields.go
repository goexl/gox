package gox

var _ pageOption = (*optionFields)(nil)

type optionFields struct {
	fields []Field
}

// Fields 配置额外信息
func Fields(fields ...Field) *optionFields {
	return &optionFields{
		fields: fields,
	}
}

func (f *optionFields) applyPage(options *pageOptions) {
	options.fields = f.fields
}
