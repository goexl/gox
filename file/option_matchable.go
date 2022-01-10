package file

var _ walkOption = (*optionMatchable)(nil)

type optionMatchable struct {
	matchable matchable
}

// Matchable 配置文件是否可被浏览
func Matchable(matchable matchable) *optionMatchable {
	return &optionMatchable{
		matchable: matchable,
	}
}

func (m *optionMatchable) applyWalk(options *walkOptions) {
	options.matchable = m.matchable
}
