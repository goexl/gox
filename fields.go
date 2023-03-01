package gox

import (
	"fmt"
	"strings"
)

// Fields 字段列表
type Fields[T any] []Field[T]

func (f Fields[T]) Add(fields ...Field[T]) (new Fields[T]) {
	new = make([]Field[T], 0, len(f)+len(fields))
	new = append(new, f...)
	new = append(new, fields...)
	fields = fields[:0]

	return
}

func (f Fields[T]) String() string {
	kvs := make([]string, 0, len(f))
	for _, field := range f {
		kvs = append(kvs, fmt.Sprintf("%s = %v", field.Key(), field.Value()))
	}

	return fmt.Sprintf("[%s]", strings.Join(kvs, ","))
}
