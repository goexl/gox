package field

import (
	"github.com/goexl/gox"
)

var (
	_                = New[int]
	_ gox.Field[any] = (*_any[int])(nil)
)

type _any[T any] struct {
	key   string
	value T
}

func New[T any](key string, value T) *_any[T] {
	return &_any[T]{
		key:   key,
		value: value,
	}
}

func (a *_any[T]) Key() string {
	return a.key
}

func (a *_any[T]) Value() any {
	return a.value
}
