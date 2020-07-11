package gox

import (
	`fmt`
	`io`
	`os`
	`path/filepath`
)

// GetFileName 获得文件名
func GetFileName(filePath string) string {
	return filePath[0 : len(filePath)-len(filepath.Ext(filePath))]
}

// GetFileNameWithExt 获得带扩展名的文件名
func GetFileNameWithExt(filePath string, ext string) (path string) {
	if "" == filepath.Ext(filePath) {
		path = filePath
	} else {
		path = fmt.Sprintf("%s.%s", GetFileName(filePath), ext)
	}

	return
}

// CopyFile 复制文件
func CopyFile(src, dest string) (bytes int64, err error) {
	var (
		srcFile  *os.File
		destFile *os.File
	)

	if _, err = os.Stat(src); err != nil {
		return
	}

	if srcFile, err = os.Open(src); err != nil {
		return
	}
	defer srcFile.Close()

	if destFile, err = os.Create(dest); err != nil {
		return
	}
	defer destFile.Close()

	// 复制文件
	bytes, err = io.Copy(destFile, srcFile)

	return
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	dir := false

	stat, err := os.Stat(path)
	if err != nil {
		dir = false
	}
	dir = stat.IsDir()

	return dir
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
