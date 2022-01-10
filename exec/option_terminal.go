package exec

import (
	`os`
)

var _ option = (*optionTerminal)(nil)

type optionTerminal struct{}

// Terminal 配置终端显示
func Terminal() *optionTerminal {
	return &optionTerminal{}
}

func (t *optionTerminal) apply(options *options) {
	options.stdout = os.Stdout
	options.stderr = os.Stderr
}
