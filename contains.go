package gox

var _ = Contains[int]

// Contains 判断元素是否在数组中
func Contains[T any](items *[]T, item T) bool {
	return -1 != Index(items, item)
}
