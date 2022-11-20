package gox

var (
	_ = UriSchemeUnknown
	_ = UriSchemeSocksV4
	_ = UriSchemeSocksV5
	_ = UriSchemeHttp
	_ = UriSchemeHttps
	_ = UriSchemeRtmp
	_ = UriSchemeMqtt
	_ = UriSchemeMqtts
	_ = UriSchemeWs
	_ = UriSchemeWss
)

const (
	// UriSchemeUnknown 未知协议
	UriSchemeUnknown UriScheme = ``
	// UriSchemeSocksV4 Socks协议
	UriSchemeSocksV4 UriScheme = "socks4"
	// UriSchemeSocksV5 Socks协议
	UriSchemeSocksV5 UriScheme = "socks4"
	// UriSchemeHttp Http协议
	UriSchemeHttp UriScheme = "http"
	// UriSchemeHttps Https协议
	UriSchemeHttps UriScheme = "https"
	// UriSchemeRtmp Rtmp协议
	UriSchemeRtmp UriScheme = "rtmp"
	// UriSchemeMqtt MQTT协议
	UriSchemeMqtt UriScheme = "mqtt"
	// UriSchemeMqtts MQTT安全协议
	UriSchemeMqtts UriScheme = "mqtts"
	// UriSchemeWs Websocket协议
	UriSchemeWs UriScheme = "ws"
	// UriSchemeWss Websocket安全协议
	UriSchemeWss UriScheme = "wss"
)

// UriScheme 协议
type UriScheme string
