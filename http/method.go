package http

import (
	"strings"
)

// Method 方法
type Method string

func (m Method) String() string {
	return string(m)
}

func (m Method) Uppercase() string {
	return strings.ToUpper(m.String())
}

func (m Method) Lowercase() string {
	return strings.ToLower(m.String())
}
