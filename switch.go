package gox

// Switch 一个通用的开关
type Switch uint64

func (s Switch) Enabled(checked Switch) bool {
	return 0 != s&checked
}
