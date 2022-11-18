package gox

import (
	"reflect"
)

var _ = Set[int]

// Set 设置方法，当`first`值为空时，设置为`second`的值，否则设置`first`的值
func Set[T any](setter func(value T), first T, second T) {
	if reflect.ValueOf(first).IsZero() {
		setter(second)
	} else {
		setter(first)
	}
}
