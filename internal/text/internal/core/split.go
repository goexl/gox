package core

import (
	"github.com/goexl/gox/internal/text/internal/param"
)

type Split struct {
	params *param.Split
}

func NewSplit(params *param.Split) *Split {
	return &Split{
		params: params,
	}
}

func (s *Split) Apply() []string {
	return s.params.Callback(s.params.From)
}
