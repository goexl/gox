package gox

var (
	_ = Index[int]
	_ = Contains[int]
	_ = Diff[int]
)

// Index 在列表中找出元素的下标
func Index[T any](items *[]T, item T) (index int) {
	index = -1
	for _index, _value := range *items {
		if any(item) == any(_value) {
			index = _index
		}
		if -1 != index {
			break
		}
	}

	return
}

// Contains 判断元素是否在数组中
func Contains[T any](items *[]T, item T) bool {
	return -1 != Index(items, item)
}

// Diff 差集
func Diff[T any](first *[]T, second *[]T) (diff []T) {
	set := make(map[any]bool)
	diff = make([]T, 0)

	for _, value := range *first {
		set[value] = true
	}
	for _, value := range *second {
		if !set[value] {
			diff = append(diff, value)
		}
	}

	return
}
