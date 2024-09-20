package param

import (
	"github.com/goexl/gox/internal/text/internal/kernel"
)

type Switch struct {
	*Upper

	From string
	Type kernel.Type
}

func NewSwitch(from string) *Switch {
	return &Switch{
		Upper: NewUpper(),

		From: from,
		Type: kernel.TypeUnderscore,
	}
}
