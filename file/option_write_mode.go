package file

var _ copyOption = (*optionWriteMode)(nil)

type optionWriteMode struct {
	mode writeMode
}

// WriteMode 配置复制模式
func WriteMode(mode writeMode) *optionWriteMode {
	return &optionWriteMode{
		mode: mode,
	}
}

func (m *optionWriteMode) apply(options *options) {
	// options.fileMode = m.fileMode
}

func (m *optionWriteMode) applyCopy(options *copyOptions) {
	options.mode = m.mode
}
