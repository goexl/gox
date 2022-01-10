package exec

var _ option = (*optionArgs)(nil)

type optionArgs struct {
	args []string
}

// Args 配置命令运行的参数
func Args(args []string) *optionArgs {
	return &optionArgs{
		args: args,
	}
}

func (a *optionArgs) apply(options *options) {
	options.args = a.args
}
