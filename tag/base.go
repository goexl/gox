package tag

import (
	"strings"
)

// base 基础标签
type base struct {
	options []string
}

// Contains 判断标签是否包含某个属性
func (b *base) Contains(option string) (contains bool) {
	contains = false

	for _, opt := range b.options {
		if strings.ToLower(opt) == option {
			contains = true
		}
		if contains {
			break
		}
	}

	return
}
