package key

import (
	`github.com/storezhang/gox`
)

var _ gox.Key = (*StringKey)(nil)

// StringKey 字符串键
type StringKey struct {
	key string
}

// NewStringKey 字符串键
func NewStringKey(k string) StringKey {
	return StringKey{
		key: k,
	}
}

func (sk *StringKey) Key() string {
	return sk.key
}
