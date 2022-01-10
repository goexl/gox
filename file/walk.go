package file

import (
	`os`
	`path/filepath`
)

// All 取得目录下所有文件，包含子目录
// 默认文件匹配所有文件
func All(dir string, opts ...walkOption) (paths []string, err error) {
	_options := defaultWalkOptions()
	for _, opt := range opts {
		opt.applyWalk(_options)
	}

	paths = make([]string, 0)
	err = filepath.Walk(dir, func(path string, info os.FileInfo, walkErr error) (err error) {
		if nil != walkErr {
			err = walkErr
		}
		if nil != err {
			return
		}
		if info.IsDir() {
			return
		}

		if matched, matchErr := filepath.Match(_options.pattern, filepath.Base(path)); matchErr != nil {
			err = matchErr
		} else if matched {
			if nil == _options.matchable || nil != _options.matchable && _options.matchable(path) {
				paths = append(paths, path)
			}
		}

		return
	})

	return
}
