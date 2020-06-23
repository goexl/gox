package gox

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

// Zip 压缩文件
func Zip(destFileName string, zipFileNames ...string) (err error) {
	var destFile *os.File

	if destFile, err = os.Create(destFileName); nil != err {
		return
	}
	defer destFile.Close()

	fileWriter := zip.NewWriter(destFile)
	defer fileWriter.Close()

	for _, zipFileName := range zipFileNames {
		if err = compressZip(zipFileName, "", fileWriter); err != nil {
			return
		}
	}

	return
}

func compressZip(zipFileName string, prefix string, zw *zip.Writer) (err error) {
	var (
		zipFile *os.File

		fileInfo   os.FileInfo
		fileInfos  []os.FileInfo
		fileHeader *zip.FileHeader

		writer io.Writer
	)

	if zipFile, err = os.Create(zipFileName); nil != err {
		return
	}
	defer zipFile.Close()

	if fileInfo, err = zipFile.Stat(); err != nil {
		return
	}

	if fileInfo.IsDir() {
		prefix = prefix + "/" + fileInfo.Name()
		if fileInfos, err = zipFile.Readdir(-1); err != nil {
			return
		}

		for _, fi := range fileInfos {
			var f *os.File
			if f, err = os.Open(zipFile.Name() + "/" + fi.Name()); err != nil {
				return
			}
			if err = compressZip(f, prefix, zw); err != nil {
				return
			}
		}
	} else {
		if fileHeader, err = zip.FileInfoHeader(fileInfo); nil != err {
			return
		}
		fileHeader.Name = prefix + "/" + fileHeader.Name

		if writer, err = zw.CreateHeader(fileHeader); err != nil {
			return
		}
		if _, err = io.Copy(writer, zipFile); err != nil {
			return
		}
	}

	return
}

// 解压
func UnZip(zipFileName, destFileName string) (err error) {
	var (
		zipReader *zip.ReadCloser
	)

	if zipReader, err = zip.OpenReader(zipFileName); err != nil {
		return
	}
	defer zipReader.Close()

	for _, zipFile := range zipReader.File {
		var reader io.ReadCloser
		if reader, err = zipFile.Open(); err != nil {
			return
		}
		defer reader.Close()

		fileName := destFileName + zipFile.Name
		if err = os.MkdirAll(getDir(fileName), 0755); err != nil {
			return
		}

		var file *os.File
		if file, err = os.Create(fileName); err != nil {
			return
		}
		defer file.Close()

		if _, err = io.Copy(file, reader); err != nil {
			return
		}
		file.Close()
		reader.Close()
	}

	return
}

func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("开始下标有错误")
	}

	if end < start || end > length {
		panic("结束下载有错误")
	}

	return string(rs[start:end])
}
