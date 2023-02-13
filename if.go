package gox

var (
	_ = If[int]
	_ = Ifx[int]
)

// If 模块条件表达式，主要用法是减少大括号
func If[T any](test bool, result T) (t T) {
	if test {
		t = result
	}

	return
}

// Ifx 模拟三元表达式，主要用法是减少大括号
func Ifx[T any](test bool, first T, second T) (t T) {
	return Ternary[T](test, first, second)
}

// Ternary 模拟三元表达式，主要用法是减少大括号
func Ternary[T any](test bool, first T, second T) (t T) {
	if test {
		t = first
	} else {
		t = second
	}

	return
}
