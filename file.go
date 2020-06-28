package gox

import (
	`fmt`
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
