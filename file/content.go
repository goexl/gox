package file

import (
	`net/http`
	`os`
)

// ContentType 获得文件的类型
func ContentType(path string) (contentType string, err error) {
	var file *os.File
	if file, err = os.Open(path); nil != err {
		return
	}
	contentType, err = ContentTypeForFile(file)

	return
}

// ContentTypeForFile 获得文件的类型
func ContentTypeForFile(file *os.File) (contentType string, err error) {
	buffer := make([]byte, 512)
	if _, err = file.Read(buffer); nil != err {
		return
	}
	contentType = http.DetectContentType(buffer)

	return
}
