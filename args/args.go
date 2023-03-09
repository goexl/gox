package args

import (
	"strings"

	"github.com/goexl/gox"
)

// Args 参数列表
type Args struct {
	args []any
	_    gox.CannotCopy
}

func newArgs(args []any) *Args {
	return &Args{
		args: args,
	}
}

func (a *Args) String() (args []string) {
	args = make([]string, 0, len(a.args))
	for _, arg := range a.args {
		args = append(args, gox.ToString(arg))
	}

	return
}

func (a *Args) Cli() string {
	return strings.Join(a.String(), space)
}
