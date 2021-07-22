package gox

type mapstructOption interface {
	apply(options *mapstructOptions)
}
