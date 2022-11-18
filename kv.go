package gox

// Kv 键值对
type Kv[K any, V any] interface {
	Key[K]
	Value[V]
}
