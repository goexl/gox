package param

import (
	"github.com/goexl/gox/internal/text/internal/kernel"
)

type Upper struct {
	Position kernel.Position
}

func NewUpper() *Upper {
	return &Upper{
		Position: kernel.PositionNone,
	}
}
