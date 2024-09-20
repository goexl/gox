package builder

import (
	"github.com/goexl/gox/internal/text/internal/kernel"
	"github.com/goexl/gox/internal/text/internal/param"
)

type Upper[F any] struct {
	from   *F
	params *param.Upper
	self   *param.Upper
}

func NewUpper[F any](from *F, params *param.Upper) *Upper[F] {
	return &Upper[F]{
		from:   from,
		params: params,
		self:   param.NewUpper(),
	}
}

func (u *Upper[F]) Head() (upper *Upper[F]) {
	u.self.Position = kernel.PositionHead
	upper = u

	return
}

func (u *Upper[F]) Tail() (upper *Upper[F]) {
	u.self.Position = kernel.PositionTail
	upper = u

	return
}

func (u *Upper[F]) All() (upper *Upper[F]) {
	u.self.Position = kernel.PositionAll
	upper = u

	return
}

func (u *Upper[F]) None() (upper *Upper[F]) {
	u.self.Position = kernel.PositionNone
	upper = u

	return
}

func (u *Upper[F]) Build() (f *F) {
	*u.params = *u.self
	f = u.from

	return
}
