package file

import (
	`regexp`
)

// ValidFilename 有效文件名（Windows标准）
func ValidFilename(filename string) bool {
	fileRegexStr := `^[^\\\./:\*\?\"<>\|]{1}[^\\/:\*\?\"<>\|]{0,254}$`
	filenameRegex := regexp.MustCompile(fileRegexStr)

	return filenameRegex.MatchString(filename)
}

// ValidFilepath 有效文件夹名
func ValidFilepath(filepath string) bool {
	pathRegexStr := `^[^\\\/\?\*\&quot;\'\&gt;\&lt;\:\|]*$`
	pathRegex := regexp.MustCompile(pathRegexStr)

	return pathRegex.MatchString(filepath)
}
