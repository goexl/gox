package http

import (
	"strings"
)

var _ = ParseMethod

// Method 方法
type Method string

// ParseMethod 解析方法
func ParseMethod(method string) (meth Method) {
	switch strings.ToLower(method) {
	case "get":
		meth = MethodGet
	case "post":
		meth = MethodPost
	case "put":
		meth = MethodPut
	case "patch":
		meth = MethodPatch
	case "delete":
		meth = MethodDelete
	}

	return
}

func (m Method) String() string {
	return string(m)
}

func (m Method) Uppercase() string {
	return strings.ToUpper(m.String())
}

func (m Method) Lowercase() string {
	return strings.ToLower(m.String())
}
