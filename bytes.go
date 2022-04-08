package gox

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/goexl/exc"
)

const (
	bytesSepratorSpace = ``

	BytesB  Bytes = 1
	BytesKB Bytes = 1024 * BytesB
	BytesMB Bytes = 1024 * BytesKB
	BytesGB Bytes = 1024 * BytesMB
	BytesTB Bytes=1024*BytesGB
	BytesPB Bytes=1024*BytesTB
	BytesEB Bytes=1024*BytesPB
)

// Bytes 字节大小
type Bytes int64

// ParseBytes 解析字节大小
func ParseBytes(str string) (bytes Bytes, err error) {
	sizes := strings.Split(str, bytesSepratorSpace)
	if 0 == len(sizes) {
		err = fmt.Errorf(exceptionBytesFormat, str)
	}
	if nil != err {
		return
	}

	// 逐步解析各个单位
	for _,size:=range sizes{
		u
	}

	return
}
