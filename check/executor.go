package check

type executor[T any] interface {
	check(check T, item T) bool
}
