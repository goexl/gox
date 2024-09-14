package param

import (
	"github.com/goexl/gox/internal/text/internal/kernel"
)

type Switch struct {
	From     string
	Position kernel.Position
	Type     kernel.Type
}

func NewSwitch(from string) *Switch {
	return &Switch{
		From:     from,
		Position: kernel.PositionHead,
		Type:     kernel.TypeUnderscore,
	}
}
