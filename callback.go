package gox

// Callback 通用回调
type Callback[T any] func(*T) error
