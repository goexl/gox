package gox

// Key 键
type Key[T any] interface {
	// Key 返回键值
	Key() T
}
