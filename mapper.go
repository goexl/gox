package gox

type Mapper[T any] interface {
	Map(T) string
}
