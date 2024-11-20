package gox

// Reset 字段重置回调，用于更新字段
type Reset[T any] func(*T)
