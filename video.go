package gox

const (
	VideoTypeFlv  VideoType = "flv"
	VideoTypeRtmp VideoType = "rtmp"
	VideoTypeHls  VideoType = "m3u8"
	VideoTypeUdp  VideoType = "flv"
)

// VideoType 视频格式类型
type VideoType string
