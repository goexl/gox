package builder

import (
	"unicode"

	"github.com/goexl/gox/internal/text/internal/core"
	"github.com/goexl/gox/internal/text/internal/param"
)

type Split struct {
	params *param.Split
}

func NewSplit(from string) *Split {
	return &Split{
		params: param.NewSplit(from),
	}
}

func (s *Split) Letter() (split *Split) {
	s.params.Callback = unicode.IsLetter
	split = s

	return
}

func (s *Split) Build() *core.Split {
	return core.NewSplit(s.params)
}
