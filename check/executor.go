package check

type executor[T any] interface {
	check(check T, targets T) bool
}
