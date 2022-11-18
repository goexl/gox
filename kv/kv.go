package kv

import (
	"github.com/goexl/gox"
)

var (
	_                     = New[int, int]
	_ gox.Kv[string, int] = (*kv[string, int])(nil)
)

type kv[K any, V any] struct {
	key   K
	value V
}

func New[K any, V any](key K, value V) *kv[K, V] {
	return &kv[K, V]{
		key:   key,
		value: value,
	}
}

func (f *kv[K, V]) Key() K {
	return f.key
}

func (f *kv[K, V]) Value() V {
	return f.value
}
