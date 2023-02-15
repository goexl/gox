package gox

// Field 字段，一个键值对，要求键必须是字符串
type Field[T any] Kv[string, T]
