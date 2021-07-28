package gox

type pageOption interface {
	applyPage(options *pageOptions)
}
