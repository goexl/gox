package builder

import (
	"github.com/goexl/gox/internal/text/internal/core"
	"github.com/goexl/gox/internal/text/internal/kernel"
	"github.com/goexl/gox/internal/text/internal/param"
)

type Switch struct {
	params *param.Switch
}

func NewSwitch(from string) *Switch {
	return &Switch{
		params: param.NewSwitch(from),
	}
}

func (s *Switch) Lowercase() (swh *Switch) {
	s.params.Type = kernel.TypeLowercase
	swh = s

	return
}

func (s *Switch) Camel() (swh *Switch) {
	s.params.Type = kernel.TypeCamel
	swh = s

	return
}

func (s *Switch) Underscore() (swh *Switch) {
	s.params.Type = kernel.TypeUnderscore
	swh = s

	return
}

func (s *Switch) Strike() (swh *Switch) {
	s.params.Type = kernel.TypeStrike
	swh = s

	return
}

func (s *Switch) Uppercase() (swh *Switch) {
	s.params.Type = kernel.TypeUppercase
	swh = s

	return
}

func (s *Switch) Head() (swh *Switch) {
	s.params.Position = kernel.PositionHead
	swh = s

	return
}

func (s *Switch) Tail() (swh *Switch) {
	s.params.Position = kernel.PositionTail
	swh = s

	return
}

func (s *Switch) All() (swh *Switch) {
	s.params.Position = kernel.PositionAll
	swh = s

	return
}

func (s *Switch) None() (swh *Switch) {
	s.params.Position = kernel.PositionNone
	swh = s

	return
}

func (s *Switch) Caption() (swh *Switch) {
	s.params.Type = kernel.TypeUppercase
	s.params.Position = kernel.PositionHead
	swh = s

	return
}

func (s *Switch) Build() *core.Switch {
	return core.NewSwitch(s.params)
}
