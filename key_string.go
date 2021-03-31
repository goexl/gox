package gox

var _ Key = (*StringKey)(nil)

// StringKey 字符串键
type StringKey struct {
	key string
}

func (sk *StringKey) Key() string {
	return sk.key
}
