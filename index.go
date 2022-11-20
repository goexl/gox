package gox

var _ = Index[int]

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
