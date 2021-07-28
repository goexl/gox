package gox

var _ mapstructOption = (*optionTag)(nil)

type optionTag struct {
	tag string
}

// Tag 配置标签名称
func Tag(tag string) *optionTag {
	return &optionTag{
		tag: tag,
	}
}

func (t *optionTag) applyMapstruct(options *mapstructOptions) {
	options.tag = t.tag
}
