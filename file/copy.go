package file

import (
	`io`
	`io/ioutil`
	`os`
	`path/filepath`

	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

// Copy 文件复制
// 如果文件有冲突，默认使用覆盖模式
func Copy(from string, to string, opts ...copyOption) (err error) {
	_options := defaultCopyOptions()
	for _, opt := range opts {
		opt.applyCopy(_options)
	}

	if !Exist(from) { // 判断源文件是否存在
		err = gox.NewFieldsError(`源文件不存在`, field.String(`filepath`, from))
	} else if Exist(to) { // 判断目的文件是否存在
		if WriteModeError == _options.mode {
			err = gox.NewFieldsError(`目的文件已存在`, field.String(`filepath`, to))
		} else if WriteModeSkip == _options.mode {
			return
		}
	}
	if nil != err {
		return
	}

	if _dir, dirErr := IsDir(from); nil != dirErr {
		err = dirErr
	} else if _dir {
		err = copyDir(from, to, _options)
	} else {
		err = copyFile(from, to, _options)
	}

	return
}

func copyDir(from string, to string, options *copyOptions) (err error) {
	var fis []os.FileInfo
	if fis, err = ioutil.ReadDir(from); nil != err {
		return
	}

	// 如果目的目录不存在，则创建目录
	if !Exist(to) {
		if fi, fiErr := os.Stat(from); nil == fiErr {
			err = os.MkdirAll(to, fi.Mode())
		} else {
			err = fiErr
		}
	}
	if nil != err {
		return
	}

	// 遍历文件列表，按分类分别复制文件、目录、链接
	for _, fi := range fis {
		source := filepath.Join(from, fi.Name())
		dest := filepath.Join(to, fi.Name())

		var info os.FileInfo
		if info, err = os.Stat(source); nil != err {
			return
		}

		switch info.Mode() & os.ModeType {
		case os.ModeDir:
			err = copyDir(source, dest, options)
		case os.ModeSymlink:
			err = copySymlink(source, dest)
		default:
			err = copyFile(source, dest, options)
		}
	}

	return
}

func copyFile(from string, to string, _ *copyOptions) (err error) {
	var toFile *os.File
	if toFile, err = os.Create(to); nil != err {
		return
	}
	defer func() {
		_ = toFile.Close()
	}()

	var fromFile *os.File
	if fromFile, err = os.Open(from); nil != err {
		return
	}
	defer func() {
		_ = fromFile.Close()
	}()

	// 复制文件内容
	if _, err = io.Copy(toFile, fromFile); nil != err {
		return
	}

	// 调整文件权限和源文件相同
	var fi os.FileInfo
	if fi, err = fromFile.Stat(); nil != err {
		return
	}

	isSymlink := fi.Mode()&os.ModeSymlink != 0
	if !isSymlink { // 不处理链接文件
		err = os.Chmod(to, fi.Mode())
	}

	return
}

func copySymlink(from string, to string) (err error) {
	var link string
	if link, err = os.Readlink(from); nil != err {
		return
	}
	err = os.Symlink(link, to)

	return
}
