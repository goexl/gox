package gox

import (
	"fmt"
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
	return fmt.Sprintf("-`%s` - `%s`", c, ToString(value))
}
