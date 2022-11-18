package gox

var _ = If[int]

// If 模拟三元表达式
func If[T any](test bool, first T, second T) (t T) {
	if test {
		t = first
	} else {
		t = second
	}

	return
}
