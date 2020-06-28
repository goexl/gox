package gox

import (
	"archive/zip"
	`bytes`
	"io"
	`io/ioutil`
	"os"
	`path/filepath`
	"strings"

	`golang.org/x/text/encoding/simplifiedchinese`
	`golang.org/x/text/transform`
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
			if err = compressZip(filepath.Join(zipFile.Name(), fi.Name()), prefix, zw); err != nil {
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
		zipReader       *zip.ReadCloser
		realZipFileName string
	)

	if zipReader, err = zip.OpenReader(zipFileName); err != nil {
		return
	}
	defer zipReader.Close()

	// 如果解压后不是放在当前目录就按照保存目录去创建目录
	if "" != destFileName {
		if err = os.MkdirAll(destFileName, 0755); nil != err {
			return
		}
	}

	for _, zipFile := range zipReader.File {
		if zipFile.Flags == 0 {
			// 如果标致位是0，则是默认的本地编码，默认为gbk
			i := bytes.NewReader([]byte(zipFile.Name))
			decoder := transform.NewReader(i, simplifiedchinese.GB18030.NewDecoder())
			content, _ := ioutil.ReadAll(decoder)
			realZipFileName = string(content)
		} else {
			// 如果标志为是 1 << 11也就是2048，则是utf-8编码
			realZipFileName = zipFile.Name
		}

		var reader io.ReadCloser
		if reader, err = zipFile.Open(); err != nil {
			return
		}
		defer reader.Close()

		path := filepath.Join(destFileName, realZipFileName)
		// 如果是目录，就创建目录
		if zipFile.FileInfo().IsDir() {
			if err = os.MkdirAll(path, zipFile.Mode()); nil != err {
				return
			}
			// 因为是目录，跳过当前循环，因为后面都是文件的处理
			continue
		}

		var file *os.File
		if file, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, zipFile.Mode()); err != nil {
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
	return subString(path, 0, strings.LastIndex(path, string(filepath.Separator)))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("开始下标有错误")
	}

	if end < start || end > length {
		panic("结束下标有错误")
	}

	return string(rs[start:end])
}
