package gox

import (
	`fmt`
	`io`
	`io/ioutil`
	`net/http`
	`os`
	`path/filepath`
	`regexp`
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

// GetContentType 获得文件的ContentType
func GetContentType(filepath string) (contentType string, err error) {
	var file *os.File
	if file, err = os.Open(filepath); nil != err {
		return
	}

	return GetFileContentType(file)
}

// GetFileContentType 获得文件的ContentType
func GetFileContentType(file *os.File) (contentType string, err error) {
	buffer := make([]byte, 512)
	if _, err = file.Read(buffer); nil != err {
		return
	}
	contentType = http.DetectContentType(buffer)

	return
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

func CopyAny(src, dst string) (copyFiles []string, err error) {
	return cpAny(src, dst)
}

func MoveAny(src, dst string) (err error) {
	_, err = cpAny(src, dst)
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

// Delete 删除的文件
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
				// 如果是自己则不显示
				return nil
			}
			infos = append(infos, path)
			return nil
		},
	)

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
		if fi, err := os.Create(file); nil != err {
			return err
		} else {
			fi.Close()
		}
		return nil
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
	var (
		path     string
		filename string
		ext      string
		tmp      string
	)

	if si, err := os.Stat(name); nil == err {
		i := 1
		path, filename = filepath.Split(name)
		ext = filepath.Ext(filename)
		tmp = filename[:(len(filename) - len(ext))]

		for {
			if si.IsDir() {
				if _, err := os.Stat(name + "(" + strconv.Itoa(i) + ")"); nil == err {
					i++
				} else {
					break
				}
			} else {
				if _, err := os.Stat(filepath.Join(path, tmp+"("+strconv.Itoa(i)+")"+ext)); nil == err {
					i++
				} else {
					break
				}
			}
		}
		if si.IsDir() {
			return name + "(" + strconv.Itoa(i) + ")"
		} else {
			return filepath.Join(path, tmp+"("+strconv.Itoa(i)+")"+ext)
		}
	}

	return name
}

func cpFile(src, dst string) (copyFile string, err error) {
	var (
		in  = new(os.File)
		out = new(os.File)
		si  os.FileInfo
	)

	if in, err = os.Open(src); nil != err {
		return
	}
	defer in.Close()

	dst = renameExist(dst)
	if out, err = os.Create(dst); nil != err {
		return
	}
	defer func() {
		if e := out.Close(); nil != err {
			err = e
		}
	}()

	if _, err = io.Copy(out, in); nil != err {
		return
	}

	if err = out.Sync(); nil != err {
		return
	}

	if si, err = os.Stat(src); nil != err {
		return
	}

	if err = os.Chmod(dst, si.Mode()); nil != err {
		return
	}
	copyFile = src

	return
}

func cpDir(src, dst string) (copyFiles []string, err error) {
	var (
		si      os.FileInfo
		entries []os.FileInfo
	)

	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	if si, err = os.Stat(src); nil != err {
		return
	}
	if !si.IsDir() {
		err = fmt.Errorf("source is not a directory")
		return
	}

	dst = renameExist(dst)
	if err = os.MkdirAll(dst, si.Mode()); nil != err {
		return
	}

	if entries, err = ioutil.ReadDir(src); nil != err {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		if entry.IsDir() {
			var tmpfiles []string
			if tmpfiles, err = cpDir(srcPath, dstPath); nil != err {
				return
			}
			copyFiles = append(copyFiles, tmpfiles...)

		} else {
			var tmpfile string
			if tmpfile, err = cpFile(srcPath, dstPath); nil != err {
				return
			}
			copyFiles = append(copyFiles, tmpfile)
		}
	}
	return
}

func cpAny(src, dst string) (copyFiles []string, err error) {
	var (
		si os.FileInfo
		di os.FileInfo
	)

	if si, err = os.Stat(src); nil != err {
		return
	}
	if si.IsDir() {
		if di, err = os.Stat(dst); nil == err {
			if os.SameFile(si, di) {
				err = fmt.Errorf("directory is itself: %s", dst)
				return
			}
			dst += "/" + filepath.Base(src)
			dst = renameExist(dst)

			var tmpFiles []string
			if tmpFiles, err = cpDir(src, dst); nil != err {
				return
			}
			copyFiles = append(copyFiles, tmpFiles...)

			return
		}

		var tmpFiles []string
		if tmpFiles, err = cpDir(src, dst); nil != err {
			return
		}
		copyFiles = append(copyFiles, tmpFiles...)
	}
	if di, err = os.Stat(dst); nil == err {
		if di.IsDir() {
			var tmpFile string
			if tmpFile, err = cpFile(src, dst+"/"+filepath.Base(src)); nil != err {
				return
			}
			copyFiles = append(copyFiles, tmpFile)

			return
		}
		if os.SameFile(si, di) {
			return
		}
		var tmpFile string
		if tmpFile, err = cpFile(src, dst+"/"+filepath.Base(src)); nil != err {
			return
		}
		copyFiles = append(copyFiles, tmpFile)

		return
	}

	var tmpFile string
	if tmpFile, err = cpFile(src, dst+"/"+filepath.Base(src)); nil != err {
		return
	}
	copyFiles = append(copyFiles, tmpFile)

	return
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

// 有效文件名（Windows标准）
func ValidFilename(filename string) bool {
	fileRegexStr := `^[^\\\./:\*\?\"<>\|]{1}[^\\/:\*\?\"<>\|]{0,254}$`
	filenamRegex := regexp.MustCompile(fileRegexStr)
	return filenamRegex.MatchString(filename)
}

// 有效文件夹名
func ValidFilepath(filepath string) bool {
	pathRegexStr := `^[^\\\/\?\*\&quot;\'\&gt;\&lt;\:\|]*$`
	pathRegex := regexp.MustCompile(pathRegexStr)
	return pathRegex.MatchString(filepath)
}

// dir以/开头
func GetDirFatherDeep(dir string) int {
	return strings.Count(dir, "/")
}

func GetDirSonDeep(dir string) int {
	isDir, err := IsDir(dir)
	if !isDir || nil != err {
		return 0
	}

	d := 1
	getDirSonDeep(dir, &d)

	return d
}

// 遍历的文件夹
func getDirSonDeep(dir string, deep *int) {
	fileInfo, err := ioutil.ReadDir(dir)
	if nil != err {
		return
	}

	// 是否加一标记 每次进入才加 不用每次加一
	isAdd := false
	// 遍历这个文件夹
	for _, fi := range fileInfo {
		// 判断是不是目录
		if fi.IsDir() {
			if !isAdd {
				*deep = *deep + 1
				isAdd = true
			}
			getDirSonDeep(dir+`/`+fi.Name(), deep)
		}
	}
}
