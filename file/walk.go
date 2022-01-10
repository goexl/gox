package file

import (
	`os`
	`path/filepath`
)

// Files 循环遍历目录
func Files(dir string, opts ...walkOption) (paths []string, err error) {
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
