package gox

import (
	`fmt`
	`io`
	`io/ioutil`
	`os`
	`path/filepath`
	`strconv`
	`strings`
	`syscall`
)

type FileType int

var (
	FileTypeAny  FileType = 0
	FileTypeDir  FileType = 1
	FileTypeFile FileType = 2
)

// GetFileName 获得文件名
func GetFileName(filePath string) string {
	return filePath[0 : len(filePath)-len(filepath.Ext(filePath))]
}

// GetFileNameWithExt 获得带扩展名的文件名
func GetFileNameWithExt(filePath string, ext string) (path string) {
	if "" == strings.TrimSpace(ext) {
		path = filePath
	} else {
		path = fmt.Sprintf("%s.%s", GetFileName(filePath), ext)
	}

	return
}

// CopyFile 复制文件
func CopyFile(src, dst string) (bytes int64, err error) {
	cpFile(src, dst)

	return
}

func CopyAny(src, dst string) error {
	return cpAny(src, dst)
}

func MoveAny(src, dst string) (err error) {
	err = cpAny(src, dst)
	if nil != err {
		return
	}
	_, err = Delete(src)
	if nil != err {
		Delete(dst)
		return
	}
	return
}

func Rename(src, dst string) error {
	return renameAny(src, dst)
}

// 返回删除的文件
func Delete(src string) (deleteFiles []string, err error) {
	var isDir bool
	if isDir, err = IsDir(src); nil != err {
		return
	}
	if isDir {
		if deleteFiles, err = GetAllFilesBy(src, FileTypeFile); nil != err {
			return
		}
	} else {
		deleteFiles = append(deleteFiles, src)
	}

	err = os.RemoveAll(src)

	return
}

func ListDir(src string) (infos []string, err error) {
	var isDir bool
	if isDir, err = IsDir(src); !isDir {
		return
	}

	err = filepath.Walk(src,
		func(path string, f os.FileInfo, err error) error {
			if nil == f {
				return err
			}
			if path == src {
				//如果是自己则不显示
				return nil
			}
			infos = append(infos, path)
			return nil
		})
	return
}

func CreateDir(dirName string) error {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dirName, 0755)
	}

	return err
}

func CreateFile(file string) error {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		_, err = os.Create(file)
	}
	return err
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) (bool, error) {
	stat, err := os.Stat(path)
	if nil != err {
		return false, err
	}
	if !stat.IsDir() {
		return false, nil
	}

	return true, nil
}

func renameExist(name string) string {
	if _, err := os.Stat(name); nil == err {
		i := 1
		for {
			if _, err := os.Stat(name + "(" + strconv.Itoa(i) + ")"); nil == err {
				i++
			} else {
				break
			}
		}
		return name + "(" + strconv.Itoa(i) + ")"
	}

	return name
}

func cpFile(src, dst string) error {
	in, err := os.Open(src)
	if nil != err {
		return err
	}
	var out *os.File
	defer in.Close()
	dst = renameExist(dst)
	out, err = os.Create(dst)
	if nil != err {
		return err
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()
	_, err = io.Copy(out, in)
	if nil != err {
		return err
	}
	err = out.Sync()
	if nil != err {
		return err
	}
	si, err := os.Stat(src)
	if nil != err {
		return err
	}

	return os.Chmod(dst, si.Mode())
}

func cpDir(src, dst string) error {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)
	si, err := os.Stat(src)
	if nil != err {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}
	dst = renameExist(dst)
	err = os.MkdirAll(dst, si.Mode())
	if nil != err {
		return err
	}
	entries, err := ioutil.ReadDir(src)
	if nil != err {
		return err
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		if entry.IsDir() {
			err = cpDir(srcPath, dstPath)
			if nil != err {
				return err
			}
		} else {
			err = cpFile(srcPath, dstPath)
			if nil != err {
				return err
			}
		}
	}

	return err
}

func cpAny(src, dst string) error {
	srcinfo, err := os.Stat(src)
	if nil != err {
		return err
	}
	if srcinfo.IsDir() {
		dstinfo, err := os.Stat(dst)
		if nil == err {
			if os.SameFile(srcinfo, dstinfo) {
				return fmt.Errorf("directory is itself: %s", dst)
			}
			dst += "/" + filepath.Base(src)
			dst = renameExist(dst)
			return cpDir(src, dst)
		}
		return cpDir(src, dst)
	}
	dstinfo, err := os.Stat(dst)
	if nil == err {
		if dstinfo.IsDir() {
			return cpFile(src, dst+"/"+filepath.Base(src))
		}
		if os.SameFile(srcinfo, dstinfo) {
			return nil
		}
		return cpFile(src, dst)
	}

	return cpFile(src, dst)
}

func renameFile(src, dst string) error {
	si, err := os.Stat(src)
	if nil != err {
		return err
	}
	if si.IsDir() {
		return fmt.Errorf("source is not a file")
	}
	dst = renameExist(dst)
	os.Rename(src, dst)

	return nil
}

func renameDir(src, dst string) error {
	si, err := os.Stat(src)
	if nil != err {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}
	dst = renameExist(dst)
	syscall.Rename(src, dst)

	return nil
}

func renameAny(src, dst string) error {
	si, err := os.Stat(src)
	if nil != err {
		return err
	}
	if si.IsDir() {
		return renameDir(src, dst)
	} else {
		return renameFile(src, dst)
	}
}

func GetAllFilesBy(pathName string, ty FileType) (allFiles []string, err error) {
	fileInfos, err := ioutil.ReadDir(pathName)
	if nil != err {
		return
	}
	for _, fi := range fileInfos {
		if fi.IsDir() {
			fullDir := filepath.Join(pathName, fi.Name())
			if ty == FileTypeDir || ty == FileTypeAny {
				allFiles = append(allFiles, fullDir)
			}

			fs := []string{}
			fs, err = GetAllFilesBy(fullDir, ty)
			if nil != err {
				return
			}
			allFiles = append(allFiles, fs...)
		} else {
			fullName := filepath.Join(pathName, fi.Name())
			if ty == FileTypeFile || ty == FileTypeAny {
				allFiles = append(allFiles, fullName)
			}
		}
	}
	return
}

// err为空存在,不会空则
func DirNotExistCratae(path string) (err error) {
	if _, err = os.Stat(path); nil != err {
		if !os.IsNotExist(err) {
			return
		}
		if err = os.MkdirAll(path, os.ModePerm); nil != err {
			return
		}
	}
	return
}



