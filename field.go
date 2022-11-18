package gox

import (
	"fmt"
	"strings"
)

type (
	// Field 字段，一个键值对，要求键必须是字符串
	Field[T any] Kv[string, T]

	// Fields 字段列表
	Fields[T any] []Field[T]
)

// Fields 支持自身连写
func (f Fields[T]) Fields() Fields[T] {
	return f
}

// Connects 连接其它字段
func (f Fields[T]) Connects(fields ...Field[T]) (new Fields[T]) {
	// 默认创建16个元素，然后再做精简
	new = make([]Field[T], 0, 16)
	for _, _f := range f {
		new = append(new, _f)
	}
	for _, field := range fields {
		new = append(new, field)
	}
	fields = fields[:0]

	return
}

// Connect 连接其它字段
func (f Fields[T]) Connect(field Field[T]) (new Fields[T]) {
	new = make([]Field[T], 0, len(f)+1)
	for _, _f := range f {
		new = append(new, _f)
	}
	new = append(new, field)

	return
}

func (f Fields[T]) String() string {
	kvs := make([]string, 0, len(f))
	for _, field := range f {
		kvs = append(kvs, fmt.Sprintf(`%s = %v`, field.Key(), field.Value()))
	}

	return fmt.Sprintf(`[%s]`, strings.Join(kvs, `,`))
}
