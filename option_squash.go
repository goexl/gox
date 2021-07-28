package gox

var _ mapstructOption = (*optionSquash)(nil)

type optionSquash struct {
	squash bool
}

// Squash 配置是否嵌套
func Squash(squash bool) *optionSquash {
	return &optionSquash{
		squash: squash,
	}
}

// EnableSquash 开启嵌套
func EnableSquash() *optionSquash {
	return Squash(true)
}

// DisableSquash 关闭嵌套
func DisableSquash() *optionSquash {
	return Squash(false)
}

func (s *optionSquash) applyMapstruct(options *mapstructOptions) {
	options.squash = s.squash
}
