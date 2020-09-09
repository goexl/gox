package gox

import (
	`errors`
	`io`
	`os`
	`path/filepath`
)

func cpFile(src string, dst string) (copyFile string, err error) {
	var (
		srcFile *os.File
		dstFile *os.File
		info    os.FileInfo
	)

	if srcFile, err = os.Open(src); nil != err {
		return
	}
	defer func() {
		err = srcFile.Close()
	}()

	ext := filepath.Ext(dst)
	if 0 == len(ext) {
		_, srcName := filepath.Split(src)
		dst = filepath.Join(dst, srcName)
	}

	if err = DirNotExistCreate(filepath.Dir(dst)); nil != err {
		return
	}

	dst = renameExist(dst)
	if dstFile, err = os.Create(dst); nil != err {
		return
	}
	defer func() {
		err = dstFile.Close()
	}()

	if _, err = io.Copy(dstFile, srcFile); nil != err {
		return
	}

	if info, err = os.Stat(src); nil == err {
		if err = os.Chmod(dst, info.Mode()); nil != err {
			return
		}
	}
	copyFile = src

	return
}

func cpDir(src string, dst string) (copyFiles []string, err error) {
	var (
		srcStat os.FileInfo
		srcDir  *os.File
		obs     []os.FileInfo
	)

	if srcStat, err = os.Stat(src); nil != err {
		return
	}

	if err = os.MkdirAll(dst, srcStat.Mode()); nil != err {
		return
	}

	if srcDir, err = os.Open(src); nil != err {
		return
	}
	defer func() {
		srcDir.Close()
	}()

	if obs, err = srcDir.Readdir(-1); nil != err {
		return
	}

	var errs []error

	for _, obj := range obs {
		srcPath := src + "/" + obj.Name()
		dstPath := renameExist(dst + "/" + obj.Name())

		if obj.IsDir() {
			var tmpFiles []string
			if tmpFiles, err = cpDir(srcPath, dstPath); nil != err {
				errs = append(errs, err)
				continue
			}
			copyFiles = append(copyFiles, tmpFiles...)

		} else {
			var tmpFile string
			if tmpFile, err = cpFile(srcPath, dstPath); nil != err {
				errs = append(errs, err)
				continue
			}
			copyFiles = append(copyFiles, tmpFile)

		}
	}

	var errString string
	for _, err := range errs {
		errString += err.Error() + "\n"
	}

	if errString != "" {
		err = errors.New(errString)
	}

	return
}

func cpAny(src, dst string) (copyFiles []string, err error) {
	var si os.FileInfo

	if si, err = os.Stat(src); nil != err {
		return
	}

	if si.IsDir() {
		tmpDst := renameExist(dst + "/" + si.Name())
		copyFiles, err = cpDir(src, tmpDst)
	} else {
		var tmpFile string
		if tmpFile, err = cpFile(src, dst); nil != err {
			return
		}

		copyFiles = append(copyFiles, tmpFile)
	}

	return
}
