package exec

import (
	`io`
)

var _ option = (*optionStderr)(nil)

type optionStderr struct {
	stderr io.Writer
}

// Stderr 配置标准错误流
func Stderr(stderr io.Writer) *optionStderr {
	return &optionStderr{
		stderr: stderr,
	}
}

func (s *optionStderr) apply(options *options) {
	options.stderr = s.stderr
}
