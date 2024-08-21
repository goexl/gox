package gox

import (
	"github.com/goexl/gox/internal/callback"
)

var (
	_ = If[int]
	_ = Iff[int]
	_ = Ift[int]
	_ = Ifx[int]
)

// If 模块条件表达式，主要用法是减少大括号
func If[T any](condition bool, result T) (t T) {
	if condition {
		t = result
	}

	return
}

// Iff 模块条件表达式，主要用法是减少大括号
func Iff[T any](condition bool, result callback.Callback[T]) (t T) {
	if condition {
		t = result()
	}

	return
}

// Ift 模拟三元表达式，主要用法是减少大括号
func Ift[T any](condition bool, first T, second T) (t T) {
	return Ternary(condition, first, second)
}

// Ternary 模拟三元表达式，主要用法是减少大括号
func Ternary[T any](condition bool, first T, second T) (t T) {
	if condition {
		t = first
	} else {
		t = second
	}

	return
}

// Ifx 模拟三元表达式，主要用法是减少大括号
func Ifx[T any](condition bool, first callback.Callback[T], second callback.Callback[T]) (t T) {
	return TernaryFunc(condition, first, second)
}

// TernaryFunc 模拟三元表达式，主要用法是减少大括号
func TernaryFunc[T any](condition bool, first callback.Callback[T], second callback.Callback[T]) (t T) {
	if condition {
		t = first()
	} else {
		t = second()
	}

	return
}
