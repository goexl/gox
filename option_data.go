package gox

var _ pageOption = (*optionData)(nil)

type optionData struct {
	fields []Field
}

// Data 配置额外信息
func Data(fields ...Field) *optionData {
	return &optionData{
		fields: fields,
	}
}

func (d *optionData) applyPage(options *pageOptions) {
	options.fields = d.fields
}
