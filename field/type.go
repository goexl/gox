package field

import (
	"github.com/goexl/gox"
)

var (
	_                = Type[int]
	_ gox.Field[int] = (*_type[int])(nil)
)

type _type[T any] struct {
	key   string
	value T
}

func Type[T any](key string, value T) *_type[T] {
	return &_type[T]{
		key:   key,
		value: value,
	}
}

func (t *_type[T]) Key() string {
	return t.key
}

func (t *_type[T]) Value() T {
	return t.value
}
