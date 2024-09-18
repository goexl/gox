package builder

import (
	"strings"

	"github.com/goexl/gox/internal/text/internal/core"
	"github.com/goexl/gox/internal/text/internal/internal/token"
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

func (s *Split) Named() (split *Split) {
	s.params.Callback = token.Named
	split = s

	return
}

func (s *Split) Separator(separator string) (split *Split) {
	s.params.Callback = func(from string) []string {
		return strings.Split(from, separator)
	}
	split = s

	return
}

func (s *Split) Build() *core.Split {
	return core.NewSplit(s.params)
}
