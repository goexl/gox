package gox

var _ mapstructOption = (*mapstructOptionTag)(nil)

type mapstructOptionTag struct {
	tag string
}

// Tag 配置标签名称
func Tag(tag string) *mapstructOptionTag {
	return &mapstructOptionTag{
		tag: tag,
	}
}

func (s *mapstructOptionTag) apply(options *mapstructOptions) {
	options.tag = s.tag
}
