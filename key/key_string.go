package key

import (
	`github.com/goexl/gox`
)

var _ gox.Key = (*StringKey)(nil)

// StringKey 字符串键
type StringKey struct {
	key string
}

// String 字符串键
func String(k string) StringKey {
	return StringKey{
		key: k,
	}
}

func (sk *StringKey) Key() string {
	return sk.key
}
