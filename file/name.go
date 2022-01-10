package file

import (
	`fmt`
	`path/filepath`
	`strings`
)

// Name 获得文件名称
func Name(path string, opts ...nameOption) (filename string) {
	_options := defaultNameOptions()
	for _, opt := range opts {
		opt.applyName(_options)
	}

	switch _options._type {
	case TypeDir:
		filename = dir(path)
	case TypeFile:
		fallthrough
	default:
		filename = name(path, _options.ext)
	}

	return
}

// NewName 新文件名，在避免文件名冲突的情况下
func NewName(path string) (filename string) {
	for {
		index := 1
		filename = filepath.Join(filepath.Dir(path), name(path, fmt.Sprintf(`%d.%s`, index, filepath.Ext(path))))
		if !Exist(filename) {
			break
		}
	}

	return
}

func name(path string, ext string) (filename string) {
	base := filepath.Base(path)
	_ext := filepath.Ext(path)
	_name := base[:len(base)-len(_ext)]

	if `` == strings.TrimSpace(ext) {
		filename = _name
	} else {
		filename = fmt.Sprintf(`%s.%s`, _name, ext)
	}

	return
}

func dir(path string) string {
	return filepath.Join(filepath.Dir(path), name(filepath.Base(path), ``))
}
