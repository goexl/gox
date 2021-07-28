package gox

type mapstructOption interface {
	applyMapstruct(options *mapstructOptions)
}
