package file

var _ walkOption = (*optionPattern)(nil)

type optionPattern struct {
	pattern string
}

// Pattern 配置文件匹配模式
func Pattern(pattern string) *optionPattern {
	return &optionPattern{
		pattern: pattern,
	}
}

func (m *optionPattern) applyWalk(options *walkOptions) {
	options.pattern = m.pattern
}
