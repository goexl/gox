package gox

// Value 值
type Value[T any] interface {
	// Value 返回具体的值
	Value() T
}
