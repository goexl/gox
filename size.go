package gox

const (
	// FilesizeByte 一个字
	FilesizeByte Datasize = 1
	// FilesizeB 字节
	FilesizeB = 8 * FilesizeByte
	// FilesizeKB 千字节
	FilesizeKB = 1024 * FilesizeB
	// FilesizeMB 兆字节
	FilesizeMB = 1024 * FilesizeKB
	// FilesizeGB 千兆字节
	FilesizeGB = 1024 * FilesizeMB
	// FilesizeTB 太字节
	FilesizeTB = 1024 * FilesizeGB
)

// Datasize 文件大小
type Bytesize int64

func (f Datasize) String() {
	unit := `b`
	num := 0.0
	switch {
	case f > FilesizeTB:
		unit = `TB`
		num = f / FilesizeTB
	}
}
