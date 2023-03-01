package gox

import (
	"fmt"
	"strings"
)

// Fields 字段列表
type Fields[T any] []Field[T]

func (f Fields[T]) Add(fields ...Field[T]) Fields[T] {
	return append(f, fields...)
}

func (f Fields[T]) String() string {
	kvs := make([]string, 0, len(f))
	for _, field := range f {
		kvs = append(kvs, fmt.Sprintf("%s = %v", field.Key(), field.Value()))
	}

	return fmt.Sprintf("[%s]", strings.Join(kvs, ","))
}
