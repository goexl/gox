package gox

// Resettable 可被重新设置的
type Resettable[T any] interface {
	Reset(Reset[T])
}
