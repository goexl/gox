package gox

import (
	"fmt"

	"github.com/goexl/gox/internal/text"
)

// Column 列
type Column string

func (c Column) String() string {
	return fmt.Sprintf("`%s`", string(c))
}

func (c Column) Negative() string {
	return fmt.Sprintf("-`%s`", c)
}

func (c Column) Negation(value any) string {
	return fmt.Sprintf("-`%s` - `%s`", c, text.ToString(value))
}

func (c Column) Asc() string {
	return fmt.Sprintf("`%s` ASC", string(c))
}

func (c Column) Desc() string {
	return fmt.Sprintf("`%s` DESC", string(c))
}
