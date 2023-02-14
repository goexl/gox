package check

const (
	hasTypeAny checkerType = iota
	hasTypeAll
)

type checkerType uint8
