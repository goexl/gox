package gox

// Pointer 指针
type Pointer[T any] *T

// NewPointer 快速创建指针
func NewPointer[T any](from T) Pointer[T] {
	return &from
}
