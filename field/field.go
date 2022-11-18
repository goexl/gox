package field

import (
	"github.com/goexl/gox"
)

var (
	_                = New[int]
	_ gox.Field[any] = (*field[int])(nil)
)

type field[T any] struct {
	key   string
	value *T
}

func New[T any](key string, value *T) *field[T] {
	return &field[T]{
		key:   key,
		value: value,
	}
}

func (f *field[T]) Key() string {
	return f.key
}

func (f *field[T]) Value() any {
	return f.value
}
