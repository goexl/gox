package gox

var _ mapstructOption = (*mapstructOptionSquash)(nil)

type mapstructOptionSquash struct {
	squash bool
}

// Squash 配置是否嵌套
func Squash(squash bool) *mapstructOptionSquash {
	return &mapstructOptionSquash{
		squash: squash,
	}
}

// EnableSquash 开启嵌套
func EnableSquash() *mapstructOptionSquash {
	return Squash(true)
}

// DisableSquash 关闭嵌套
func DisableSquash() *mapstructOptionSquash {
	return Squash(false)
}

func (s *mapstructOptionSquash) apply(options *mapstructOptions) {
	options.squash = s.squash
}
