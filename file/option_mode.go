package file

import (
	`os`
)

var _ option = (*optionMode)(nil)

type optionMode struct {
	mode os.FileMode
}

// Mode 配置权限
func Mode(mode os.FileMode) *optionMode {
	return &optionMode{
		mode: mode,
	}
}

func (m *optionMode) apply(options *options) {
	options.fileMode = m.mode
}
