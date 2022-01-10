package file

import (
	`os`
	`syscall`

	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

// Create 创建文件或者目录
// 默认创建文件
// 默认权限是0777
func Create(path string, opts ...option) (err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	if Exist(path) {
		switch _options.writeMode {
		case WriteModeError:
			err = gox.NewFieldsError(`文件已存在`, field.String(`path`, path))
		case WriteModeSkip:
			return
		case WriteModeOverride:
			err = Delete(path)
		case WriteModeRename:
			path = NewName(path)
		}
	}
	if nil != err {
		return
	}

	// 创建文件或者目录
	switch _options._type {
	case TypeDir:
		err = os.MkdirAll(path, _options.fileMode)
	default:
		_, err = os.Create(path)
	}
	if nil != err {
		return
	}

	// 改变文件的拥有者
	if nil != _options.owner {
		err = os.Chown(path, _options.owner.uid, _options.owner.gid)
	}

	return
}

// Exist 判断文件是否存在
func Exist(filename string) (exist bool) {
	if _, err := os.Stat(filename); nil != err && os.IsNotExist(err) {
		exist = false
	} else {
		exist = true
	}

	return
}

// Rename 重命名文件或者目录
func Rename(from string, to string) error {
	return syscall.Rename(from, to)
}

// Delete 删除文件或者目录
func Delete(filename string) error {
	return os.RemoveAll(filename)
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) (dir bool, err error) {
	var stat os.FileInfo
	if stat, err = os.Stat(path); nil != err {
		return
	}

	dir = stat.IsDir()
	stat = nil

	return
}
