package text

import (
	"github.com/goexl/gox/internal/text/internal/builder"
)

type String struct {
	from string
}

func NewString(from string) *String {
	return &String{
		from: from,
	}
}

func (s String) Switch() *builder.Switch {
	return builder.NewSwitch(s.from)
}
